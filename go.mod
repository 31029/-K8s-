module k8sapi

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/gorilla/websocket v1.4.2
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/shenyisyn/goft-gin v0.5.0
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292
	gopkg.in/check.v1 v1.0.0-20200902074654-038fdea0a05b // indirect
	k8s.io/api v0.24.0
	k8s.io/apimachinery v0.24.0
	k8s.io/client-go v0.24.0
	k8s.io/metrics v0.24.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
