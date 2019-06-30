package schedule

//go:generate protoc -I/usr/local/include -I. -I$GOPATH -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. schedule.proto

//go:generate protoc -I/usr/local/include -I. -I$GOPATH -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. schedule.proto
