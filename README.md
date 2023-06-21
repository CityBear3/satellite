# Satellite

The service hub to call iot camera from Slack, Discord or CLI etc.

## Development

### build proto file

```shell
protoc --proto_path=protoc --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative protoc/**/v1/*.proto
```