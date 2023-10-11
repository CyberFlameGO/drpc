module storj.io/drpc/examples/grpc_and_drpc

go 1.17

require (
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.27.1
	storj.io/drpc v0.0.17
)

require (
	github.com/golang/protobuf v1.5.0 // indirect
	github.com/zeebo/errs v1.2.2 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
)

replace storj.io/drpc => ../..
