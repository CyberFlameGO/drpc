module storj.io/drpc/examples/grpc_and_drpc

go 1.17

require (
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/grpc v1.56.3
	google.golang.org/protobuf v1.30.0
	storj.io/drpc v0.0.17
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/zeebo/errs v1.2.2 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
)

replace storj.io/drpc => ../..
