mkdir -p vendor
PROTOC_OPTS='-I tensorflow -I serving --go_out=plugins=grpc:vendor'

eval "protoc $PROTOC_OPTS serving/tensorflow_serving/apis/*.proto"
eval "protoc $PROTOC_OPTS serving/tensorflow_serving/apis/internal/*.proto"
eval "protoc $PROTOC_OPTS serving/tensorflow_serving/core/*.proto"
eval "protoc $PROTOC_OPTS serving/tensorflow_serving/util/*.proto"
eval "protoc $PROTOC_OPTS serving/tensorflow_serving/config/*.proto"
eval "protoc $PROTOC_OPTS serving/tensorflow_serving/resources/*.proto"
eval "protoc $PROTOC_OPTS serving/tensorflow_serving/servables/hashmap/*.proto"
eval "protoc $PROTOC_OPTS serving/tensorflow_serving/servables/tensorflow/*.proto"
eval "protoc $PROTOC_OPTS serving/tensorflow_serving/session_bundle/oss/*.proto"
eval "protoc $PROTOC_OPTS serving/tensorflow_serving/sources/storage_path/*.proto"
