clean:
	docker-compose rm -vf

build:
	docker-compose build

up:
	docker-compose up -d
	migrate create -ext sql -dir deployment/migrations -seq init_schema

migrate_up:
	migrate -path /deployment/migrations/ -source file://deployment/migrations -database "mysql://root:password@tcp(127.0.0.1:3308)/order?charset=utf8mb4&parseTime=True&loc=Local" -verbose up