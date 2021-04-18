ARG registry
FROM golang:1.16 as builder

RUN mkdir -p /build
COPY . /build
RUN cd /build/cmd && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o rabbitmq-http-auth .

FROM scratch

COPY --from=builder /build/cmd/rabbitmq-http-auth /

WORKDIR /app
ENTRYPOINT ["/rabbitmq-http-auth"]
