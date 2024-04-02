ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder-go

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run-app .

FROM node:16 as builder-node
COPY web/ .
RUN npm install && npm run build

FROM debian:bookworm

COPY --from=builder-go /run-app ./
COPY --from=builder-node /dist ./dist
RUN apt update && apt install -y ca-certificates
CMD ["./run-app"]
