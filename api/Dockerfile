FROM golang:1.21-alpine AS builder
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN apk add --no-cache git build-base
RUN go mod download
COPY . .
RUN go build -o /usr/local/bin/api ./...

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /usr/local/bin/api /usr/local/bin/api
EXPOSE 8000
CMD ["/usr/local/bin/api"]
