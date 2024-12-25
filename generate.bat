@echo off
protoc ^
  --proto_path=protobuf "protobuf/buildrequest.proto" ^
  --go_out=services/common/buildrequest --go_opt=paths=source_relative ^
  --go-grpc_out=services/common/buildrequest --go-grpc_opt=paths=source_relative
