module go-gin-api

go 1.16

require (
	github.com/gin-gonic/gin v1.7.2
	github.com/go-playground/locales v0.13.0
	github.com/go-playground/universal-translator v0.17.0
	github.com/go-playground/validator/v10 v10.4.1
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/raylin666/go-utils v0.0.0-20210529030518-1030944dc1e0
	google.golang.org/grpc v1.38.0 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	gorm.io/gorm v1.21.10
)

replace github.com/raylin666/go-utils => ../go-utils
