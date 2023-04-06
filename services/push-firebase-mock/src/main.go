package main

import (
	"fmt"
	"git.protei.ru/protei-golang/common/ffsm"
	"git.protei.ru/protei-golang/common/ffsm/model"
	log "git.protei.ru/protei-golang/common/logger"
	"git.protei.ru/protei-golang/common/pconfig"
	"git.protei.ru/protei-golang/common/psync"
	"git.protei.ru/uc/core/services/push-firebase-mock/config"
	"git.protei.ru/uc/core/services/push-firebase-mock/server"
	"git.protei.ru/uc/core/services/push-firebase-mock/service"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cfg, err := config.InitializeConfiguration(os.Args[1:])
	if err == pconfig.ErrServiceInstanceNotFound {
		log.Warn(err)
	} else if err != nil {
		log.Fatal(err)
	}
	defer log.SyncZap()

	sessionHolder := &service.ActiveSessionHolder{}
	actorModel := model.GetInstanceActorModel()
	err = actorModel.RegisterService(service.WsSessionName,
		cfg.GetWorkersCount(service.WsSessionName, 1),
		ffsm.WithRouterArg(service.RouteWsSession),
		ffsm.WithCreateArg(func() ffsm.IFsmLight {
			return service.CreateWsSession(sessionHolder, actorModel)
		}))
	if err != nil {
		log.Fatalf("error during register service %s: %s", service.WsSessionName, err.Error())
	}
	if err = actorModel.Start(); err != nil {
		log.Fatalf("error during start services: %s", err.Error())
	}
	wg := psync.SpawnSafe(func() { runHttpServer(cfg, actorModel, sessionHolder) })
	wg.Wait()
}

func runHttpServer(cfg *config.PushFirebaseMockConfig, actorModel model.ActorModel,
	sessionHolder *service.ActiveSessionHolder) {
	s := server.NewServer(actorModel, sessionHolder)

	http.Handle("/", s.Handler)
	serverHost := fmt.Sprintf("%s:%d", cfg.ServerEndpoint.ServerHost, cfg.ServerEndpoint.ListenPort)
	err := http.ListenAndServeTLS(serverHost, cfg.TlsCertPath, cfg.TlsKeyPath, nil) //nolint:gosec
	if err != nil {
		log.Warnf("error during listen: %s", err.Error())
		return
	}
}
