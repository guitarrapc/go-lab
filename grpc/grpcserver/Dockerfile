# syntax=docker/dockerfile:1
FROM golang:1.18-alpine as dev

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download
EXPOSE 8080

CMD ["go", "run", "main.go"]


FROM golang:1.18-alpine as builder

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download

COPY . ${ROOT}
RUN GOOS=linux go build -o $ROOT/binary


FROM scratch as prod

ENV ROOT=/go/src/app
WORKDIR ${ROOT}
COPY --from=builder ${ROOT}/binary ${ROOT}

EXPOSE 50051
CMD ["/go/src/app/binary"]
