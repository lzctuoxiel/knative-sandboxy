module knative.dev/net-kourier

go 1.16

require (
	github.com/cncf/xds/go v0.0.0-20210922020428-25de7278fc84 // indirect
	github.com/envoyproxy/go-control-plane v0.9.10-0.20210928151155-81157c60e0e3
	github.com/google/go-cmp v0.5.6
	github.com/google/uuid v1.3.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	go.uber.org/zap v1.19.1
	google.golang.org/genproto v0.0.0-20211016002631-37fc39342514
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
	gotest.tools/v3 v3.0.3
	k8s.io/api v0.21.4
	k8s.io/apimachinery v0.21.4
	k8s.io/client-go v0.21.4
	k8s.io/code-generator v0.21.4
	knative.dev/hack v0.0.0-20211019034732-ced8ce706528
	knative.dev/networking v0.0.0-20211019132235-c8d647402afa
	knative.dev/pkg v0.0.0-20211019132235-ba2b2b1bf268
)
