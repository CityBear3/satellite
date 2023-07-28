MIGRATION = ""

gen-proto:
	protoc --proto_path=protoc --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative protoc/**/v1/*.proto

test:
	go test -v ./...

gen-schema:
	sqlboiler mysql -c database.toml -o ./adaptor/repository/mysql/shcema -p schema --no-tests --wipe

gen-migration:
	migrate create -ext sql -dir ./db/migration -seq ${MIGRATION}

migrate-up:
	migrate --path db/migration --database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}" --verbose up

migrate-down:
	migrate --path db/migration --database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}" --verbose down

migrate-drop:
	migrate --path db/migration --database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}" --verbose drop