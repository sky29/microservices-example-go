protoc -I $GOPATH/src --go_out=$GOPATH/src $GOPATH/src/microservices-example-go/review/internal/proto-files/domain/review.proto
protoc -I $GOPATH/src --go_out=plugins=grpc:$GOPATH/src $GOPATH/src/microservices-example-go/review/internal/proto-files/service/review-service.proto
