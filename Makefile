clean:
	docker-compose rm -vf
	rm proto/*/*.pb.go

build:
	docker-compose build

up:
	docker-compose up -d
	migrate create -ext sql -dir deployment/migrations -seq init_schema

dev:
	go get -u google.golang.org/grpc
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go

migrate_up:
	migrate -path /deployment/migrations/ -source file://deployment/migrations -database "mysql://root:password@tcp(127.0.0.1:3308)/order?charset=utf8mb4&parseTime=True&loc=Local" -verbose up

compile_protos:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/*/*.proto