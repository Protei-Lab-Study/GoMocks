package service

import (
	"git.protei.ru/protei-golang/common/ffsm"
	"git.protei.ru/protei-golang/common/ffsm/messages"
	"git.protei.ru/protei-golang/common/ffsm/model"
	log "git.protei.ru/protei-golang/common/logger"
	"git.protei.ru/protei-golang/common/pconfig"
	api "git.protei.ru/uc/core/services/GO-TEMPLATE-SERVICE/api/grpc"
	"strconv"
)

const TEST_SERVICE_NAME = "TEST_SERVICE_NAME"

type InitSessionMsg struct {
	UserId int64
}

type AbstractRequest struct {
	UserId  int64
	Request interface{}
}

func (i *InitSessionMsg) Hash() string {
	return strconv.FormatInt(i.UserId, 10)
}

func (m *InitSessionMsg) GetShortName() string {
	return "service.InitSessionMsg"
}

func (i *AbstractRequest) Hash() string {
	return strconv.FormatInt(i.UserId, 10)
}

func (m *AbstractRequest) GetShortName() string {
	return "service.AbstractRequest"
}

type TestServiceFsm struct {
	ffsm.Ffsm
	model  model.ActorModel
	userId int64
}

func TestServiceRouter(msg messages.IMsg) ffsm.RouterResult {
	if msg.GetDstDID() != "" {
		return ffsm.RouterResult{DialogId: msg.GetDstDID(), Route: ffsm.RouterResultAuto}
	}
	switch body := msg.GetBody().(type) {
	case *InitSessionMsg:
		return ffsm.RouterResult{DialogId: strconv.FormatInt(body.UserId, 10), Route: ffsm.RouterResultNew}
	case *AbstractRequest:
		return ffsm.RouterResult{DialogId: strconv.FormatInt(body.UserId, 10), Route: ffsm.RouterResultIn}
	default:
		return ffsm.RouterResult{Route: ffsm.RouterResultNone}
	}
}

func CreateTestServiceFsm(_ *pconfig.ActorSettings) ffsm.IFsmLight {
	fsm := &TestServiceFsm{}
	fsm.InitFfsm(fsm.Idle)
	fsm.model = model.GetInstanceActorModel()
	return fsm
}

func (s *TestServiceFsm) Idle(rawMsg messages.IMsg) ffsm.HandlerResult {
	log.Debugfw(s.DID(), "Msg: %v", rawMsg)
	switch msg := rawMsg.(type) {
	case *messages.Msg:
		switch msg.Body.(type) {
		case *InitSessionMsg:
			return s.Swap(s.Doing)
		default:
			return s.processUnhandled(rawMsg)
		}
	default:
		return s.processUnhandled(rawMsg)
	}
}

func (s *TestServiceFsm) Doing(rawMsg messages.IMsg) ffsm.HandlerResult {
	log.Debugfw(s.DID(), "Msg: %v", rawMsg)
	switch msg := rawMsg.GetBody().(type) {
	case *AbstractRequest:
		switch msg.Request.(type) {
		case *api.Method1Request:
			response := &api.Method1Response{}
			om := s.model.MakeReplyMessage(rawMsg, response)
			s.SendOut(om)
		case *api.Method2Request:
			response := &api.Method2Response{Results: []string{"result1", "result2"}}
			om := s.model.MakeReplyMessage(rawMsg, response)
			s.SendOut(om)
		}
	}
	return s.processUnhandled(rawMsg)
}

func (s *TestServiceFsm) processUnhandled(rawMsg messages.IMsg) ffsm.HandlerResult {
	log.Debugfw(s.DID(), "Unexpected message: %v", rawMsg)
	return s.Done()
}
