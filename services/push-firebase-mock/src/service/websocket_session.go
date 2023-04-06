package service

import (
	"git.protei.ru/protei-golang/common/ffsm"
	"git.protei.ru/protei-golang/common/ffsm/messages"
	"git.protei.ru/protei-golang/common/ffsm/model"
	log "git.protei.ru/protei-golang/common/logger"
	"git.protei.ru/protei-golang/common/psync"
	"git.protei.ru/protei-golang/common/utils"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/thoas/go-funk"
	"net/http"
)

const (
	WsSessionName            = "WebSocketSessionName"
	WsSessionStatusConnected = iota + 1
	WsSessionStatusDisconnected
)

type WsRequest struct {
	Request   interface{}
	SessionId string
}

type WsInitMsg struct {
	Response  http.ResponseWriter
	Request   *http.Request
	PushToken string
}

type WsDisconnectedMgs struct {
}

type WsNotificationMsg struct {
	Tokens       []string
	Notification []byte
}

type WsInitMsgResult struct {
	SessionId      string
	WsStatusNotify wsSessionStatus
}

type wsSessionStatus int

type WsSession struct {
	wsConn        *websocket.Conn
	sessionHolder *ActiveSessionHolder
	actorModel    model.ActorModel
	upgrader      websocket.Upgrader
	// функция показывает, нужно или нет посылать переданную Push нотификацию
	needSendPushFunc func(tokens []string) (needSend bool)
	ffsm.Ffsm
}

func (msg *WsRequest) Hash() string {
	if msg.SessionId == "" {
		msg.SessionId = uuid.New().String()
	}

	return msg.SessionId
}

func (msg *WsRequest) GetShortName() string {
	return "service.WsRequest"
}

func CreateWsSession(sessionHolder *ActiveSessionHolder, actorModel model.ActorModel) *WsSession {
	server := &WsSession{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		sessionHolder: sessionHolder,
		actorModel:    actorModel,
	}
	server.InitFfsm(server.onIdle)
	return server
}

func (s *WsSession) String() string {
	formatter := utils.InitStructFormatterWithSeparator("NOTIFICATION PROXY WEBSOCKET SESSION", "\n\n")
	formatter.AppendString("FFSM", s.Ffsm.String())
	formatter.AppendStringer("RemoteAddr", s.wsConn.RemoteAddr())
	formatter.AppendStringer("LocalAddr", s.wsConn.LocalAddr())
	formatter.AppendString("SubProtocol", s.wsConn.Subprotocol())
	return formatter.String()
}

func RouteWsSession(msg messages.IMsg) ffsm.RouterResult {
	if msg.GetDstDID() != "" {
		return ffsm.RouterResult{DialogId: msg.GetDstDID(), Route: ffsm.RouterResultAuto}
	}

	switch body := msg.GetBody().(type) {
	case *WsRequest:
		return ffsm.RouterResult{DialogId: body.SessionId, Route: ffsm.RouterResultAuto}
	default:
		return ffsm.RouterResult{Route: ffsm.RouterResultNone}
	}
}

func (s *WsSession) Terminate(_ interface{}) []messages.IMsg {
	s.sessionHolder.RemoveSession(s.DID())
	if s.wsConn != nil {
		if err := s.wsConn.Close(); err != nil {
			log.Error("error during close ws connection: ", err)
		}
	}
	return nil
}

func (s *WsSession) onIdle(rawMsg messages.IMsg) ffsm.HandlerResult {
	log.Debug("ws session: ", rawMsg)
	switch msg := rawMsg.GetBody().(type) {
	case *WsRequest:
		switch request := msg.Request.(type) {
		case *WsInitMsg:
			var err error
			s.wsConn, err = s.upgrader.Upgrade(request.Response, request.Request, nil)
			if err != nil {
				log.Errorf("Failed to connect to websocket due to: %v", err)
				s.SendOut(model.MakeReplyMessage(rawMsg, &WsInitMsgResult{WsStatusNotify: WsSessionStatusDisconnected}))
				return s.Stop("failed connect")
			}
			if request.PushToken != "" {
				s.needSendPushFunc = func(tokens []string) (needSend bool) {
					return funk.ContainsString(tokens, request.PushToken)
				}
			} else {
				s.needSendPushFunc = func([]string) bool {
					return true
				}
			}
			s.sessionHolder.AddSession(s.DID())
			psync.SpawnSafe(s.listenWS)
			s.SendOut(model.MakeReplyMessage(rawMsg, &WsInitMsgResult{WsStatusNotify: WsSessionStatusConnected,
				SessionId: s.DID()}))
			return s.Swap(s.onActive)
		}
	}
	return s.processUnhandled(rawMsg)
}

func (s *WsSession) onActive(rawMsg messages.IMsg) ffsm.HandlerResult {
	log.Debugfw(s.DID(), "rawMsg: %v", rawMsg)
	switch msg := rawMsg.GetBody().(type) {
	case *WsRequest:
		switch request := msg.Request.(type) {
		case *WsNotificationMsg:
			if s.needSendPushFunc(request.Tokens) {
				err := s.wsConn.WriteMessage(websocket.TextMessage, request.Notification)
				if err != nil {
					log.Errorfw(s.DID(), "error during write to ws: %v", err)
				}
			}
			s.SendOut(model.MakeReplyMessage(rawMsg, &struct{}{}))
			return s.Done()
		case *WsDisconnectedMgs:
			return s.Stop("Websocket closed")
		}
	}
	return s.processUnhandled(rawMsg)
}

func (s *WsSession) listenWS() {
	for {
		if s.wsConn == nil {
			return
		}
		_, _, err := s.wsConn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure,
				websocket.CloseGoingAway, websocket.CloseNoStatusReceived) {
				log.Infofw(s.DID(), "Websocket disconnected")
			} else {
				log.Errorfw(s.DID(), "websocket error %v", err)
			}
			break
		}
	}
	err := s.actorModel.SendMessage(WsSessionName, &WsRequest{SessionId: s.DID(), Request: &WsDisconnectedMgs{}})
	if err != nil {
		log.Errorfw(s.DID(), "failed send message due to: %v", err)
	}
}

func (s *WsSession) processUnhandled(msg messages.IMsg) ffsm.HandlerResult {
	switch msg.GetBody().(type) {
	case *WsRequest:
		log.Warnfw(s.DID(), "Unexpected message: %v at %s state", msg.GetBody(), s.GetCurrentStageName())
		s.SendOut(model.MakeReplyMessage(msg, &WsInitMsgResult{WsStatusNotify: WsSessionStatusDisconnected}))
	case *ffsm.Enter, *ffsm.Leave:
		return s.Done()
	case ffsm.FfsmSystemMessage:
		log.Debugf("unhandled msg: %v", msg)
	}
	return s.Done()
}
