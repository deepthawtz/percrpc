example gRPC server/client
==========================

super dumb percentage calculation service client/server service implemented w/ gRPC

### Run it

`go run server/main.go`

and in separate window/tab

`go run client/main.go -o DIVIDEND -n DIVISOR`

### Protocol Buffers

See [proto definition](percentage/percentage.proto)
