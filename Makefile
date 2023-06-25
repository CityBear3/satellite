gen-proto:
	protoc --proto_path=protoc --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative protoc/**/v1/*.proto

test:
	go test -v $(go list ./... | grep -v pb)