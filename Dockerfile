ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder-go

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run-app .

FROM debian:bookworm

COPY --from=builder-go /run-app ./
RUN apt update && apt install -y ca-certificates
CMD ["./run-app"]
