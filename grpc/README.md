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
mkdir server
cd server
go mod init github.com/guitarrapc/go-grpc/server
go get
go build
```

```bash
mkdir client
cd client
go mod init github.com/guitarrapc/go-grpc/client
go get
go build
```

## Docker and Compose

thanks https://zenn.dev/tatsurom/articles/golang-docker-environment
