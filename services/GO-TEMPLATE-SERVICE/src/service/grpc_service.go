package service

import (
	"context"
	"git.protei.ru/protei-golang/common/ffsm/model"
	log "git.protei.ru/protei-golang/common/logger"
	"git.protei.ru/uc/core/services/GO-TEMPLATE-SERVICE/api/grpc"
	"time"
)

type GoTemplateServiceServerImpl struct {
	api.UnimplementedGoTemplateServiceServer
	Model model.ActorModel
}

func (g *GoTemplateServiceServerImpl) TestMethod1(_ context.Context, req *api.Method1Request) (response *api.Method1Response, responseError error) {
	response = &api.Method1Response{}
	defer LogFunction("TestMethod1", req, response)()
	_, responseError = g.Model.SendMessageRequest(TEST_SERVICE_NAME, g.Model.MakeMessage(TEST_SERVICE_NAME, &AbstractRequest{UserId: 1000, Request: req}), 2*time.Second)
	return response, responseError
}

func (g *GoTemplateServiceServerImpl) TestMethod2(_ context.Context, req *api.Method2Request) (response *api.Method2Response, responseError error) {
	response = &api.Method2Response{}
	defer LogFunction("TestMethod2", req, response)()
	rawResult, responseError := g.Model.SendMessageRequest(TEST_SERVICE_NAME, g.Model.MakeMessage(TEST_SERVICE_NAME, &AbstractRequest{UserId: 1000, Request: req}), 2*time.Second)
	switch result := rawResult.(type) {
	case *api.Method2Response:
		response.Results = result.Results
	}
	return response, responseError
}

func LogFunction(name string, req interface{}, resp interface{}) func() {
	id := time.Now().UnixNano()
	log.Debugf("%d: %s.Req: %v", id, name, req)
	return func() { log.Debugf("%d: %s.Resp: %v", id, name, resp) }
}
