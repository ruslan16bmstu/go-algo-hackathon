build-server:
	go build

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
	    --openapiv2_out ./pb/gen                         \
	    ./pb/api.proto


clean:
	rm pocket-trader-backend


.PHONY: build-server run-server clean