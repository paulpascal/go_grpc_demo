# generate proto buffer

protoc --proto_path=proto --go_out=proto --go_opt=paths=source_relative --go-grpc_out=proto --go-grpc_opt=paths=source_relative service.proto

# for

protoc --dart_out=grpc:proto -Iproto proto/service.proto
"# go_grpc_demo" 
