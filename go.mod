module knative.dev/net-kourier

go 1.15

require (
	github.com/envoyproxy/go-control-plane v0.9.9-0.20210217033140-668b12f5399d
	github.com/google/go-cmp v0.5.5
	github.com/google/uuid v1.2.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	go.uber.org/zap v1.16.0
	google.golang.org/genproto v0.0.0-20210416161957-9910b6c460de
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
	gotest.tools/v3 v3.0.3
	k8s.io/api v0.19.7
	k8s.io/apimachinery v0.19.7
	k8s.io/client-go v0.19.7
	knative.dev/hack v0.0.0-20210426064739-88c69cd1eca7
	knative.dev/networking v0.0.0-20210426163440-2696fca263ec
	knative.dev/pkg v0.0.0-20210426101439-2a0fc657a712
)
