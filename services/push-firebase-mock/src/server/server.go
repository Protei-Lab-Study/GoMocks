package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"git.protei.ru/protei-golang/common/ffsm/model"
	log "git.protei.ru/protei-golang/common/logger"
	"git.protei.ru/uc/core/services/push-firebase-mock/service"
	"github.com/gorilla/mux"
	"io"
	"mime"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Server struct {
	http.Handler
	Model         model.ActorModel
	sessionHolder *service.ActiveSessionHolder
}

type PushNotify struct {
	MulticastMessage MulticastMessage `json:"multicastMessage"`
}

type MulticastMessage struct {
	Tokens []string `json:"Tokens"`
}

const (
	MaxNotificationSize     = 1 * 1024 * 1024 // 1Mb
	PushNotificationsLogTag = "PUSH-NOTIFICATIONS-TAG"
	tokenKey                = "token"
)

func NewServer(model model.ActorModel, sessionHolder *service.ActiveSessionHolder) *Server {
	s := &Server{Model: model, sessionHolder: sessionHolder}

	router := mux.NewRouter()
	router.HandleFunc("/notifications", s.notificationHandler).Methods(http.MethodPost)
	router.HandleFunc("/notifications/subscribe", s.subscribeOnNotifications).Methods(http.MethodGet)
	router.HandleFunc(fmt.Sprintf("/notifications/subscribe/{%s:.+}", tokenKey), s.subscribeOnNotifications).Methods(http.MethodGet)
	s.Handler = router

	return s
}

func (s *Server) SetContentType(w http.ResponseWriter, contentType string) {
	w.Header().Set("Content-Type", contentType)
}

func (s *Server) subscribeOnNotifications(w http.ResponseWriter, r *http.Request) {
	subscribePushToken := mux.Vars(r)[tokenKey]
	rawResult, sendError := s.Model.SendMessageRequest(
		service.WsSessionName,
		&service.WsRequest{Request: &service.WsInitMsg{Response: w, Request: r, PushToken: subscribePushToken}},
		time.Minute)
	if sendError != nil {
		log.Errorf("Failed send message due to: %v", sendError)
	} else {
		if result, ok := rawResult.(*service.WsInitMsgResult); ok {
			if result.WsStatusNotify == service.WsSessionStatusConnected {
				log.Infof("Websocket %s connected", result.SessionId)
				return
			}
		}
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func (s *Server) notificationHandler(w http.ResponseWriter, r *http.Request) {
	if !ValidateContentType(w, r) {
		return
	}
	if !ValidateContentLength(w, r) {
		return
	}

	activeSessions := s.sessionHolder.GetActiveSessions()
	if notificationData, success := ExtractNotificationAndLog(w, r, len(activeSessions)); success {
		s.SendNotificationToSessions(w, notificationData, activeSessions)
	}
}

// ValidateContentLength проверяет размер тела сообщения и возвращает http ошибку, если размер не соответствует условиям
// return bool true - размер валидный
func ValidateContentLength(w http.ResponseWriter, r *http.Request) bool {
	contentLength := uint32(r.ContentLength)
	if contentLength <= 0 {
		contentLength2, err := strconv.ParseUint(r.Header.Get("Content-Length"), 10, 32) //nolint:govet
		if err != nil {
			log.Errorf("Can't find size of body")
			w.WriteHeader(http.StatusBadRequest)
			return false
		}
		contentLength = uint32(contentLength2)
	}
	if contentLength > MaxNotificationSize {
		log.Warnf("The user tries to proxy notification with a size greater than the maximum "+
			"allowed size: %d (bytes), maximum %d (bytes)", contentLength, MaxNotificationSize)
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		return false
	}
	return true
}

// ValidateContentType проверяет тип тела сообщения и возвращает http ошибку, если это не json
// return bool true - json является типом контента
func ValidateContentType(w http.ResponseWriter, r *http.Request) bool {
	contentType := r.Header.Get("Content-type")
	if !IsContentType(contentType, "application/json") {
		log.Warnf("Invalid content type: %s", contentType)
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return false
	}
	return true
}

// ExtractNotificationAndLog извлекает нотификацию из запроса и логирует ее в pretty json
// return ([]byte, bool) - (нотификация, успешность операции)
func ExtractNotificationAndLog(w http.ResponseWriter, r *http.Request, countOfActiveSessions int) ([]byte, bool) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Errorf("Failed read body due to: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return nil, false
	}
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, data, "", "\t")
	if err != nil {
		log.Errorf("Failed ident json due to: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return nil, false
	}
	log.SystemNamed(PushNotificationsLogTag).Infof("Notification send to %d sessions:\n%v",
		countOfActiveSessions, prettyJSON.String())
	return data, true
}

// SendNotificationToSessions отправить нотификацию во все активные сессии с ожиданием отправки по ws
func (s *Server) SendNotificationToSessions(w http.ResponseWriter, notificationData []byte, activeSessions []string) {
	var wg sync.WaitGroup
	pushNotify, err := s.parsePushNotification(notificationData)
	if err != nil {
		log.Warnf("Bad notification request: can't unmarshal push notification: %s", notificationData)
		w.WriteHeader(http.StatusBadRequest)
	}
	for _, wsId := range activeSessions {
		wg.Add(1)
		sessionId := wsId
		go func() {
			_, sendError := s.Model.SendMessageRequest(
				service.WsSessionName,
				&service.WsRequest{SessionId: sessionId,
					Request: &service.WsNotificationMsg{Notification: notificationData, Tokens: pushNotify.MulticastMessage.Tokens}},
				time.Minute) // signal that the routine has completed
			if sendError != nil {
				log.Errorf("Failed send message due to: %v", sendError)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	w.WriteHeader(http.StatusOK)
}

func (s *Server) parsePushNotification(notificationData []byte) (*PushNotify, error) {
	var notify PushNotify
	err := json.Unmarshal(notificationData, &notify)
	if err != nil {
		return nil, err
	}
	return &notify, nil
}

func IsContentType(contentType string, mimetype string) bool {
	t, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return false
	}
	if t != mimetype {
		return false
	}
	return true
}
