build-server:
	go mod tidy && go build

run-server:
	./pocket-trader-backend

gen-proto:
	protoc                                               \
	    -I./pb                                           \
	    -I//home/maksim/Compilers/protoc/include         \
	    --go_out=./pb/gen                                \
	    --go_opt paths=source_relative                   \
	    --go-grpc_out=./pb/gen                           \
	    --go-grpc_opt paths=source_relative              \
	    --grpc-gateway_out=./pb/gen                      \
	    --grpc-gateway_opt paths=source_relative         \
	    --grpc-gateway_opt generate_unbound_methods=true \
	    --openapiv2_out ./swagger                        \
	    ./pb/api.proto

build-docker:
	docker build --pull -f ./Dockerfile -t moex-hack-backend .

run-docker:
	docker run \
		-p 8000:8000 \
		-p 8001:8001 \
		--name=moex-hack-backend-instance \
		moex-hack-backend

clean:
	rm pocket-trader-backend


.PHONY: build-server run-server gen-proto clean