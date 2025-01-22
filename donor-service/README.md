protoc --proto_path=. \
    --go_out=paths=source_relative:./pb \
    --go-grpc_out=paths=source_relative:./pb \
    donor.proto
