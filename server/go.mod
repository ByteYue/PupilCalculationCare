module my-app

go 1.13

require (
	db v0.0.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kyleconroy/sqlc v1.6.0 // indirect
	github.com/lib/pq v1.9.0
	github.com/spf13/viper v1.7.0
	github.com/ugorji/go v1.2.1 // indirect
	go.mongodb.org/mongo-driver v1.4.1
	go.uber.org/zap v1.15.0
	golang.org/x/crypto v0.0.0-20201208171446-5f87f3452ae9 // indirect
	golang.org/x/sys v0.0.0-20201211002650-1f0c578a6b29 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace db => ./db/sqlc
