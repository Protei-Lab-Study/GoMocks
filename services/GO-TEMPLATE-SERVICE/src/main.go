package main

import (
	"context"
	"fmt"
	"git.protei.ru/protei-golang/common/ffsm"
	"git.protei.ru/protei-golang/common/ffsm/amqp"
	"git.protei.ru/protei-golang/common/ffsm/model"
	grpclib "git.protei.ru/protei-golang/common/grpc-lib"
	log "git.protei.ru/protei-golang/common/logger"
	"git.protei.ru/protei-golang/common/pconfig"
	"git.protei.ru/protei-golang/common/pmetrics"
	"git.protei.ru/protei-golang/common/pshell"
	"git.protei.ru/protei-golang/common/psync"
	api "git.protei.ru/uc/core/services/GO-TEMPLATE-SERVICE/api/grpc"
	cfg "git.protei.ru/uc/core/services/GO-TEMPLATE-SERVICE/config"
	"git.protei.ru/uc/core/services/GO-TEMPLATE-SERVICE/service"
	"git.protei.ru/uc/core/services/golang/shellcommands"
	"google.golang.org/grpc"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	config, err := cfg.InitializeConfiguration(os.Args[1:])
	if err == pconfig.ErrServiceInstanceNotFound {
		log.Warn(err)
	} else if err != nil {
		log.Fatal(err)
	}
	defer log.SyncZap()
	err = pmetrics.InitGlobalMetricsReporter(config)
	if err != nil {
		log.Fatalf("error during metrics init: %s", err.Error())
	}
	actorModel := model.GetInstanceActorModel()
	err = amqp.RegisterAmqpClient(config)
	if err != nil {
		log.Fatalf("error during register service %s: %s", amqp.AmqpClientName, err.Error())
	}
	err = actorModel.RegisterService(service.TEST_SERVICE_NAME,
		config.GetWorkersCount(service.TEST_SERVICE_NAME, 1),
		ffsm.WithRouterArg(service.TestServiceRouter),
		ffsm.WithCreateArg(func() ffsm.IFsmLight {
			return service.CreateTestServiceFsm(config.GetActorSettings(service.TEST_SERVICE_NAME))
		}))
	if err != nil {
		log.Fatalf("error during register service %s: %s", service.TEST_SERVICE_NAME, err.Error())
	}
	if err = actorModel.Start(); err != nil {
		log.Fatalf("error during start services: %s", err.Error())
	}
	// Только для теста!
	_ = actorModel.SendMessage(service.TEST_SERVICE_NAME, actorModel.MakeMessage(service.TEST_SERVICE_NAME, &service.InitSessionMsg{UserId: 1000}))
	err = pshell.StartServer(shellcommands.DefaultCommandFactory())
	if err != nil {
		log.Errorf("Can't start shell command server cause: %v", err)
	}
	wg := psync.SpawnSafe(func() { runGRPCServer(config) }, func() { runGRPCClient(config) })
	wg.Wait()
}

// Запуск gRPC сервиса
func runGRPCServer(config *cfg.GO_TEMPLATE_SERVICE_Config) {
	// Set up a connection to the server.
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.ServerEndpoint.ServerHost, config.ServerEndpoint.ListenPort))
	if err != nil {
		log.Errorf("failed to listen: %v", err)
	}
	s, err := grpclib.NewServer(config, config.ServerEndpoint)
	if err != nil {
		log.Errorf("failed to listen: %v", err)
	}
	api.RegisterGoTemplateServiceServer(s, &service.GoTemplateServiceServerImpl{Model: model.GetInstanceActorModel()})
	if err := s.Serve(lis); err != nil {
		log.Errorf("failed to serve: %v", err)
	}
	log.Infof(" ...server stopped")
}

// Запуск gRPC клиента (ТЕСТОВОГО) для своего сервиса
func runGRPCClient(config *cfg.GO_TEMPLATE_SERVICE_Config) {
	target := fmt.Sprintf("%s:%d", config.ServerEndpoint.ClientHost, config.ServerEndpoint.Port)
	conn, err := grpc.Dial(target, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Errorf("did not connect: %v", err)
	}
	client := api.NewGoTemplateServiceClient(conn)
	log.Infof(" client started")
	log.Infof("Make TestMethod1")
	result1, error1 := client.TestMethod1(context.Background(), &api.Method1Request{
		Arg11: "111",
		Arg12: false,
	})
	time.Sleep(500 * time.Millisecond)
	log.Infof("TestMethod2 result: %#v, error: %#v", result1, error1)
	log.Infof("Make TestMethod2")
	result2, error2 := client.TestMethod2(context.Background(), &api.Method2Request{
		Arg21: "111",
		Arg22: 23423,
	})
	log.Infof("TestMethod2 result: %#v, error: %#v", result2, error2)
}
