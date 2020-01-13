protoc -I $GOPATH/src --go_out=$GOPATH/src $GOPATH/src/microservices-example-go/details/internal/proto-files/domain/details.proto

protoc -I $GOPATH/src --go_out=plugins=grpc:$GOPATH/src $GOPATH/src/microservices-example-go/details/internal/proto-files/service/details-service.proto
