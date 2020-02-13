protoc --go_out=plugins=grpc:$GOPATH/src -I ../uiprpc-base -I . wsrpc.proto
