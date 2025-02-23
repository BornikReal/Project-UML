.PHONY: generate
generate:
	protoc \
    	--proto_path=api/v1/ \
    	--go_out=pkg \
    	--go-grpc_out=pkg \
    	--grpc-gateway_out pkg \
    	--openapiv2_out pkg \
    	--openapiv2_opt logtostderr=true \
    	--openapiv2_opt allow_merge=true \
    	--grpc-gateway_opt generate_unbound_methods=true \
    	api/v1/*.proto

.PHONY: build
build:
	mkdir -p bin && cd cmd && go build -o ../bin/storage_service

.PHONY: up
up:
	docker-compose build && docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: show_logs
show_logs:
	docker exec -it storage_service cat server.log