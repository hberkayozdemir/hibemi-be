module gitlab.com/modanisatech/marketplace/service-template

go 1.14

require (
	github.com/gofiber/adaptor/v2 v2.1.2
	github.com/gofiber/fiber/v2 v2.14.0
	github.com/joho/godotenv v1.3.0
	github.com/pact-foundation/pact-go v1.4.4
	github.com/phayes/freeport v0.0.0-20180830031419-95f893ade6f2
	github.com/prometheus/client_golang v0.9.3
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/yudai/pp v2.0.1+incompatible
	gitlab.com/modanisatech/marketplace/shared/errors v0.0.3
	gitlab.com/modanisatech/marketplace/shared/httpkit v0.1.1
	gitlab.com/modanisatech/marketplace/shared/kafka v0.0.4
	gitlab.com/modanisatech/marketplace/shared/log v0.0.9
	gitlab.com/modanisatech/marketplace/shared/unleash v0.0.2
	go.uber.org/zap v1.16.0
)

replace github.com/coreos/etcd => github.com/coreos/etcd v3.3.24+incompatible

replace github.com/dgrijalva/jwt-go => github.com/golang-jwt/jwt v3.2.1+incompatible

replace gitlab.com/modanisatech/marketplace/shared/unleash => gitlab.com/modanisatech/marketplace/shared/unleash.git v0.0.2

replace gitlab.com/modanisatech/marketplace/shared/log => gitlab.com/modanisatech/marketplace/shared/log.git v0.0.9

replace gitlab.com/modanisatech/marketplace/shared/errors => gitlab.com/modanisatech/marketplace/shared/errors.git v0.0.3

replace gitlab.com/modanisatech/marketplace/shared/kafka => gitlab.com/modanisatech/marketplace/shared/kafka.git v0.0.4

replace gitlab.com/modanisatech/marketplace/kafka-retry => gitlab.com/modanisatech/marketplace/kafka-retry.git v0.0.2

replace gitlab.com/modanisatech/marketplace/shared/httpkit => gitlab.com/modanisatech/marketplace/shared/httpkit.git v0.1.1

replace github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2
