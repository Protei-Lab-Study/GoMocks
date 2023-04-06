package config

import (
	"encoding/json"
	"git.protei.ru/protei-golang/common/pconfig"
	ucConfig "git.protei.ru/uc/core/services/golang/config"
	"time"
)

// GO_TEMPLATE_SERVICE_Config - структура с описанием конфигурации текущего микросервиса
type GO_TEMPLATE_SERVICE_Config struct {
	*pconfig.BaseInfluxConfig
	*pconfig.BaseTelemetryConfig
	*pconfig.ActorConfig
	// Endpoint собственного сервиса, т.е. тот endpoint на котором он доступен
	ServerEndpoint *pconfig.Endpoint `endpoint:"server-endpoint"`
	// Тестовый endpoint для внешнего сервиса, с которым работает наш сервис
	TestOnlyExternalServiceEndpoint *pconfig.Endpoint `endpoint:"test-only-external-service-endpoint"`
	// Настройка коннекции до Amqp
	Amqp *pconfig.Amqp `amqp:"amqp-broker"`
	// Любые другие параметры
	Key1 string        `mapstructure:"key1"`
	Key2 bool          `mapstructure:"key2"`
	Key3 time.Duration `mapstructure:"key3"`
}

func (a *GO_TEMPLATE_SERVICE_Config) GetAmqp() *pconfig.Amqp {
	return a.Amqp
}

// GetServiceName - метод возвращает имя данного микросервиса, которое используется в файле конфигурации.
func (a *GO_TEMPLATE_SERVICE_Config) GetServiceName() string {
	return "GO-TEMPLATE-SERVICE"
}

func (a *GO_TEMPLATE_SERVICE_Config) String() string {
	data, _ := json.MarshalIndent(a, " ", " ")
	return string(data)
}

// DefaultConfig - метод возвращает конфиг по умолчанию для данного микросервиса
func DefaultConfig() *GO_TEMPLATE_SERVICE_Config {
	return &GO_TEMPLATE_SERVICE_Config{
		BaseTelemetryConfig:             ucConfig.DefaultTelemetryConfig(),
		BaseInfluxConfig:                ucConfig.DefaultInfluxConfig(),
		ServerEndpoint:                  &pconfig.Endpoint{ServerHost: "0.0.0.0", ListenPort: 8310},
		TestOnlyExternalServiceEndpoint: &pconfig.Endpoint{ClientHost: "127.0.0.1", Port: 8304},
		Amqp:                            &pconfig.Amqp{Url: "amqp://127.0.0.1:5672"},
		Key1:                            "value1",
		Key2:                            true,
		Key3:                            time.Second,
		ActorConfig:                     pconfig.NewActorConfig(),
	}
}

// InitializeConfiguration - метод выполняет начальную инициализацию конфигурации микросервиса из файла конфигурации.
func InitializeConfiguration(commandLineArgs []string) (*GO_TEMPLATE_SERVICE_Config, error) {
	defaultConfig := DefaultConfig()
	err := pconfig.InitializeConfiguration(commandLineArgs, defaultConfig)
	return defaultConfig, err
}

// CurrentConfig - метод возвращает типизированную конфигурацию текущего микросервиса
func CurrentConfig() *GO_TEMPLATE_SERVICE_Config {
	return pconfig.GetServiceConfig().(*GO_TEMPLATE_SERVICE_Config)
}
