FROM golang:1.25.0-alpine3.22 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o /go/bin/webserver .
FROM scratch
COPY --from=builder /go/bin/webserver /webserver
ENTRYPOINT ["/webserver"]