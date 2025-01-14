module wrpc.io/examples/go/envoy-filter

go 1.22.2

require wrpc.io/go v0.1.0

require (
	cel.dev/expr v0.16.0 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.1.0 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240814211410-ddb44dafa142 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142 // indirect
)

require (
	github.com/cncf/xds/go v0.0.0-20241223141626-cff3c89139a3
	github.com/envoyproxy/envoy v1.32.3
	google.golang.org/protobuf v1.35.2
)

replace wrpc.io/go v0.1.0 => /source/go
