PROTO_SRC_DIR := pkg
PROTO_FILES := $(wildcard $(PROTO_SRC_DIR)/*.proto)

ifeq ($(PROTO),)
	PROTOS := $(wildcard $(PROTO_SRC_DIR)/*.proto)
else
	PROTOS := $(PROTO_SRC_DIR)/$(PROTO)
endif

all: run

up:
	docker compose up -d
stop:
	docker compose stop
down:
	docker compose down

run:
	go run application/*.go

docker-build:
	docker build -t ta-dd-dashboard-analytic .
docker-run:
	docker run -d -p 8080:8080 -p 9090:9090 ta-dd-dashboard-analytic

protoc: 
	protoc --go_out=. --go-grpc_out=. $(PROTOS)
