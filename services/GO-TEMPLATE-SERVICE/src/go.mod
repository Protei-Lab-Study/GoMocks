module git.protei.ru/uc/core/services/GO-TEMPLATE-SERVICE

go 1.16

require (
	git.protei.ru/protei-golang/common v1.3.25
	git.protei.ru/uc/core/services/golang v0.0.0
	github.com/golang/protobuf v1.5.2
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
)

replace git.protei.ru/uc/core/services/golang v0.0.0 => ../../golang
