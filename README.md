# go-algo-hackathon

## TODO

1. Сделать настроить генерацию из прото в grpc и http-proxy
2. Реализовать выгрузку данных из бд
3. Реализовать загрузку данных из ALGOPACK в бд
4. Реализовать отправку сообщений пользователю который подключил свое собственное api


- Install protobuf;

- To install `grpc-reverse-proxy-tools`
```
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

