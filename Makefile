.PHONY: build test run

build:
	cd cmd/example && \
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build -a -tags netgo -ldflags '-w -extldflags "-static"' \
				 -o rabbitmq-http-auth .
run:
	cd cmd/example && ./rabbitmq-http-auth

test:
	go test -v -cover -coverprofile=coverage.out github.com/jandelgado/rabbitmq-http-auth/pkg
	go tool cover -func=coverage.out

