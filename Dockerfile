ARG registry
FROM golang:1.16 as builder

RUN mkdir -p /build
COPY . /build
RUN cd /build && make build

FROM scratch

COPY --from=builder /build/cmd/example/rabbitmq-http-auth /

WORKDIR /app
ENTRYPOINT ["/rabbitmq-http-auth"]
