module github.com/femitubosun/streaming-pipeline/processor

go 1.25.5

require (
	github.com/caarlos0/env/v11 v11.4.1
	github.com/twmb/franz-go v1.21.2
)

require (
	golang.org/x/crypto v0.51.0 // indirect
	golang.org/x/net v0.53.0 // indirect
	golang.org/x/sys v0.44.0 // indirect
	golang.org/x/text v0.37.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260226221140-a57be14db171 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

require (
	github.com/femitubosun/streaming-pipeline/proto v0.0.0-00010101000000-000000000000
	github.com/klauspost/compress v1.18.6 // indirect
	github.com/pierrec/lz4/v4 v4.1.26 // indirect
	github.com/twmb/franz-go/pkg/kadm v1.18.0
	github.com/twmb/franz-go/pkg/kmsg v1.13.1 // indirect
	google.golang.org/grpc v1.81.1
)

replace github.com/femitubosun/streaming-pipeline/proto => ../proto
