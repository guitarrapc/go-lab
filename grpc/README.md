# go-grpc

Minimum implementation to try gRPC Server/Client with golang.

## Local Execution

Use docker or local build.

```
# docker
docker compose build && docker compose up

# local build
cd server && go build && server.exe
cd client && go build && client.exe
```

# Appendix

## Reproduce how to create projects

```bash
mkdir grpc
go mod init github.com/guitarrapc/go-lab/server
mkdir cmd/grpcserver
mkdir cmd/grpcclient
# place code to cmd/grpcserver/main.go
# place code to cmd/grpcclient/main.go
go get ./cmd/grpcclient
go get ./cmd/grpcserver
go build ./cmd/grpcclient
go build ./cmd/grpcserver
```


## Docker and Compose

thanks https://zenn.dev/tatsurom/articles/golang-docker-environment
