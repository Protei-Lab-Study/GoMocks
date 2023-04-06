package server

import (
	"git.protei.ru/protei-golang/common/ffsm"
	"git.protei.ru/protei-golang/common/ffsm/model"
	"git.protei.ru/protei-golang/common/tests"
	"git.protei.ru/uc/core/services/push-firebase-mock/service"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type serverTestInstance struct {
	t              tests.TestingT
	model          model.ActorModel
	sessionHolder  *service.ActiveSessionHolder
	httpServerTest *httptest.Server
	server         *Server
}

func TestServerCreateAndCloseWsSession(t *testing.T) {
	s := newServerTest(t)
	defer s.teardown()

	wsURL := "ws" + strings.TrimPrefix(s.httpServerTest.URL, "http") + "/notifications/subscribe"
	testsCases := []struct {
		name  string
		wsURL string
	}{
		{"Тест создания сессии при подписки на все токены", "ws" + strings.TrimPrefix(s.httpServerTest.URL, "http") + "/notifications/subscribe"},
		{"Тест создания сессии при подписки на определенный токены", "ws" + strings.TrimPrefix(s.httpServerTest.URL, "http") + "/notifications/subscribe/test-token-1"},
	}
	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
			require.NoError(t, err)
			assert.Equal(t, 1, len(s.sessionHolder.GetActiveSessions()))
			_ = ws.Close()
			time.Sleep(time.Millisecond * 30)
			assert.Equal(t, 0, len(s.sessionHolder.GetActiveSessions()))
		})
	}
}

func TestServerProxyNotifications(t *testing.T) {
	s := newServerTest(t)
	defer s.teardown()

	wsURL := "ws" + strings.TrimPrefix(s.httpServerTest.URL, "http") + "/notifications/subscribe"
	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	require.NoError(t, err)
	defer func() { _ = ws.Close() }()

	var wsNotification string
	go func() {
		messageType, readBody, err := ws.ReadMessage()
		require.NoError(t, err)
		require.Equal(t, websocket.TextMessage, messageType)
		if readBody != nil {
			wsNotification = string(readBody)
		}
	}()

	testText := `{
        "multicastMessage": {
                "Tokens": [
                        "123123123:12312312-123123-43564645dfgd"
                ],
                "Data": {
                        "$type": "pushDataEventType",
                        "body": "{\"plaintext\":\"или нет?\",\"type\":\"TEXT\",\"$type\":\"pushChatEventType\"}",
                        "chat_avatar_url": "https://uc.protei.ru/avatars/ryabkov@protei-lab.ru",
                        "chat_event_id": "65214",
                        "chat_id": "7962670490631587100",
                        "chat_name": "Рябков Антон Николаевич",
                        "chat_type": "P2P",
                        "owner_id": "2162094504465089962",
                        "sender_id": "109",
                        "sender_name": "Рябков Антон Николаевич",
                        "sent_at": "2023-03-28T06:58:18.579+03:00"
                },
                "Notification": null,
                "Android": {
                        "ttl": "1s",
                        "priority": "high"
                },
                "Webpush": null,
                "APNS": null
        },
        "pushMethod": "pushChatEventType"
}`
	body := io.NopCloser(strings.NewReader(testText))
	r, _ := http.NewRequest(http.MethodPost, "/notifications", body)
	r.ContentLength = int64(len(testText))
	r.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.server.ServeHTTP(w, r)
	time.Sleep(time.Millisecond * 30)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testText, wsNotification)
}

func TestServerProxyNotificationsWithToken(t *testing.T) {
	s := newServerTest(t)
	defer s.teardown()

	wsURL := "ws" + strings.TrimPrefix(s.httpServerTest.URL, "http") + "/notifications/subscribe/test-token-1"
	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	require.NoError(t, err)
	defer func() { _ = ws.Close() }()

	var wsNotification string
	go func() {
		messageType, readBody, err := ws.ReadMessage()
		require.NoError(t, err)
		require.Equal(t, websocket.TextMessage, messageType)
		if readBody != nil {
			wsNotification = string(readBody)
		}
	}()

	testText := `{
        "multicastMessage": {
                "Tokens": [
                        "test-token-1",
                        "test-token-2"
                ],
                "Data": {
                        "$type": "pushDataEventType",
                        "body": "{\"plaintext\":\"или нет?\",\"type\":\"TEXT\",\"$type\":\"pushChatEventType\"}",
                        "chat_avatar_url": "https://uc.protei.ru/avatars/ryabkov@protei-lab.ru",
                        "chat_event_id": "65214",
                        "chat_id": "7962670490631587100",
                        "chat_name": "Рябков Антон Николаевич",
                        "chat_type": "P2P",
                        "owner_id": "2162094504465089962",
                        "sender_id": "109",
                        "sender_name": "Рябков Антон Николаевич",
                        "sent_at": "2023-03-28T06:58:18.579+03:00"
                },
                "Notification": null,
                "Android": {
                        "ttl": "1s",
                        "priority": "high"
                },
                "Webpush": null,
                "APNS": null
        },
        "pushMethod": "pushChatEventType"
}`
	body := io.NopCloser(strings.NewReader(testText))
	r, _ := http.NewRequest(http.MethodPost, "/notifications", body)
	r.ContentLength = int64(len(testText))
	r.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.server.ServeHTTP(w, r)
	time.Sleep(time.Millisecond * 30)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testText, wsNotification)
}

func TestServerProxyNotificationsWithNotSubscribeToken(t *testing.T) {
	s := newServerTest(t)
	defer s.teardown()

	wsURL := "ws" + strings.TrimPrefix(s.httpServerTest.URL, "http") + "/notifications/subscribe/test-token-not-subscribe"
	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	require.NoError(t, err)
	defer func() { _ = ws.Close() }()

	go func() {
		_, _, wsErr := ws.ReadMessage()
		require.Error(t, wsErr)
	}()

	testText := `{
        "multicastMessage": {
                "Tokens": [
                        "test-token-1",
                        "test-token-2"
                ],
                "Data": {
                        "$type": "pushDataEventType",
                        "body": "{\"plaintext\":\"или нет?\",\"type\":\"TEXT\",\"$type\":\"pushChatEventType\"}",
                        "chat_avatar_url": "https://uc.protei.ru/avatars/ryabkov@protei-lab.ru",
                        "chat_event_id": "65214",
                        "chat_id": "7962670490631587100",
                        "chat_name": "Рябков Антон Николаевич",
                        "chat_type": "P2P",
                        "owner_id": "2162094504465089962",
                        "sender_id": "109",
                        "sender_name": "Рябков Антон Николаевич",
                        "sent_at": "2023-03-28T06:58:18.579+03:00"
                },
                "Notification": null,
                "Android": {
                        "ttl": "1s",
                        "priority": "high"
                },
                "Webpush": null,
                "APNS": null
        },
        "pushMethod": "pushChatEventType"
}`
	body := io.NopCloser(strings.NewReader(testText))
	r, _ := http.NewRequest(http.MethodPost, "/notifications", body)
	r.ContentLength = int64(len(testText))
	r.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	s.server.ServeHTTP(w, r)
	time.Sleep(time.Millisecond * 300)
	assert.Equal(t, http.StatusOK, w.Code)
}

func newServerTest(t tests.TestingT) *serverTestInstance {
	test := &serverTestInstance{}
	test.setup(t)
	return test
}

func (s *serverTestInstance) setup(t tests.TestingT) {
	s.t = t
	s.sessionHolder = &service.ActiveSessionHolder{}

	s.model = model.GetInstanceActorModel()
	err := s.model.RegisterService(service.WsSessionName,
		1,
		ffsm.WithRouterArg(service.RouteWsSession),
		ffsm.WithCreateArg(func() ffsm.IFsmLight {
			return service.CreateWsSession(s.sessionHolder, s.model)
		}))
	require.NoError(t, err)
	err = s.model.Start()
	require.NoError(t, err)

	s.server = NewServer(s.model, s.sessionHolder)
	s.httpServerTest = httptest.NewServer(s.server)
}

func (s *serverTestInstance) teardown() {
	_ = s.model.UnregisterService(service.WsSessionName)
	s.httpServerTest.Close()
}
