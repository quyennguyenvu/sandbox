if [ -z "${PROTOBUF_HOME}" ]; then
    GOOGLE_PROTOBUF_INCLUDE="$(dirname $(dirname $(which protoc)))/include"
else
    GOOGLE_PROTOBUF_INCLUDE="${PROTOBUF_HOME}/include"
fi
ENTITY=$1
protoc -I${GOOGLE_PROTOBUF_INCLUDE} \
  -I./pb/${ENTITY} \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:./pb/${ENTITY} \
  --grpc-gateway_out=logtostderr=true:./pb/${ENTITY} \
  --swagger_out=logtostderr=true:./pb/${ENTITY} \
  ./pb/${ENTITY}/*.proto
