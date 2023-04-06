package config

import (
	"encoding/json"
	"git.protei.ru/protei-golang/common/pconfig"
	"git.protei.ru/uc/core/services/push-firebase-mock/service"
)

// PushFirebaseMockConfig - структура с описанием конфигурации текущего микросервиса
type PushFirebaseMockConfig struct {
	*pconfig.ActorConfig
	// Endpoint собственного сервиса, т.е. тот endpoint на котором он доступен
	ServerEndpoint *pconfig.Endpoint `endpoint:"server-endpoint"`
	TlsCertPath    string            `mapstructure:"tls-cert-path"`
	TlsKeyPath     string            `mapstructure:"tls-key-path"`
}

// GetServiceName - метод возвращает имя данного микросервиса, которое используется в файле конфигурации.
func (a *PushFirebaseMockConfig) GetServiceName() string {
	return "push-firebase-mock"
}

func (a *PushFirebaseMockConfig) String() string {
	data, _ := json.MarshalIndent(a, " ", " ")
	return string(data)
}

// DefaultConfig - метод возвращает конфиг по умолчанию для данного микросервиса
func DefaultConfig() *PushFirebaseMockConfig {
	return &PushFirebaseMockConfig{
		ServerEndpoint: &pconfig.Endpoint{ServerHost: "0.0.0.0", ListenPort: 8319},
		TlsCertPath:    "./cert/server.crt",
		TlsKeyPath:     "./cert/server.key",
		ActorConfig: pconfig.NewActorConfig().
			WithActorSettings(service.WsSessionName, &pconfig.ActorSettings{WorkerCount: 4}),
	}
}

// InitializeConfiguration - метод выполняет начальную инициализацию конфигурации микросервиса из файла конфигурации.
func InitializeConfiguration(commandLineArgs []string) (*PushFirebaseMockConfig, error) {
	defaultConfig := DefaultConfig()
	err := pconfig.InitializeConfiguration(commandLineArgs, defaultConfig)
	return defaultConfig, err
}

// CurrentConfig - метод возвращает типизированную конфигурацию текущего микросервиса
func CurrentConfig() *PushFirebaseMockConfig {
	return pconfig.GetServiceConfig().(*PushFirebaseMockConfig)
}
