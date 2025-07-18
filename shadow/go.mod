module github.com/vison888/iot-engine/shadow

go 1.20

require (
	github.com/BurntSushi/toml v1.3.2
	github.com/envoyproxy/protoc-gen-validate v1.0.2
	github.com/nats-io/nats.go v1.29.0
	// github.com/vison888/go-vkit v0.0.0-incompatible
	github.com/vison888/go-vkit 1.0.0
	google.golang.org/genproto/googleapis/api v0.0.0-20231016165738-49dd2c1f3d0b
	google.golang.org/grpc v1.59.0
	google.golang.org/protobuf v1.31.0
	gorm.io/gorm v1.25.5
)

require (
	github.com/Masterminds/squirrel v1.5.4
	github.com/robfig/cron/v3 v3.0.1
	github.com/spf13/cast v1.5.1
	github.com/taosdata/driver-go/v3 v3.5.0
	github.com/vison888/iot-engine/common v0.0.0-20231218092732-4206564b1e5a
	github.com/vison888/iot-engine/group v0.0.0-20231218092732-4206564b1e5a
	github.com/vison888/iot-engine/thing v0.0.0-20231218092732-4206564b1e5a
)

require (
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nats-io/nkeys v0.4.6 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto v0.0.0-20231012201019-e917dd12ba7a // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231012201019-e917dd12ba7a // indirect
	gorm.io/driver/mysql v1.5.1 // indirect
)

// replace github.com/vison888/go-vkit v0.0.0-incompatible => ../../../vison888/go-vkit
