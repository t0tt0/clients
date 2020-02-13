protoc -I ../uiprpc-base -I . uiprpc.proto --go_out=plugins=grpc:$GOPATH/src 
