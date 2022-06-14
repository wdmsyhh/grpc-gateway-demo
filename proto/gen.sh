protoc -I. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative trip.proto

protoc -I. --grpc-gateway_out=paths=source_relative,grpc_api_configuration=trip.yaml:. trip.proto
