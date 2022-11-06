prepare:
  # https://developers.google.com/protocol-buffers/docs/proto3
  # In general you should set the --proto_path flag to the root of your project and use fully qualified names for all imports.
	protoc --proto_path=. \
	  --go_out=. --go_opt=paths=source_relative \
	  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/greet.proto