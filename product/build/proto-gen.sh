protoc -I $GOPATH/src --go_out=$GOPATH/src $GOPATH/src/microservices-example-go/product/internal/proto-files/domain/product.proto

protoc -I $GOPATH/src --go_out=plugins=grpc:$GOPATH/src $GOPATH/src/microservices-example-go/product/internal/proto-files/service/product-service.proto
