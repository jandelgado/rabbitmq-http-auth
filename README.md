# RabbitMQ HTTP Auth Backend in Go

[![run tests](https://github.com/jandelgado/rabbitmq-http-auth/actions/workflows/test.yml/badge.svg)](https://github.com/jandelgado/rabbitmq-http-auth/actions/workflows/test.yml)
[![Coverage Status](https://coveralls.io/repos/github/jandelgado/rabbitmq-http-auth/badge.svg?branch=main)](https://coveralls.io/github/jandelgado/rabbitmq-http-auth?branch=main)

Package and example service to build a RabbitMQ HTTP Auth service for use with
the RabbitMQ "HTTP Auth Backend" (actually it is an AuthN/AuthZ backend).

For details see https://github.com/rabbitmq/rabbitmq-server/tree/master/deps/rabbitmq_auth_backend_http

<!-- vim-markdown-toc GFM -->

* [Build your own service](#build-your-own-service)
* [Test it](#test-it)
* [Test with RabbitMQ](#test-with-rabbitmq)
* [Author & License](#author--license)

<!-- vim-markdown-toc -->

## Build your own service

To build an RabbitMQ HTTP Auth Backend, you just need to implement the provided
`Authenticator` interface, which will be called by `POST` requests to the paths
`/auth/user`, `/auth/vhost`, `/auth/topic` and `/auth/resource`:

```go
type Decision bool

type Authenticator interface {
	// User authenticates the given user. In addition to the decision, the tags
	// associated with the user are returned.
	User(username, password string) (Decision, string)
	// VHost checks if the given user/ip combination is allowed to access the
	// vhosts
	VHost(username, vhost, ip string) Decision
	// Resource checks if the given user has access to the presented resource
	Resource(username, vhost, resource, name, permission string) Decision
	// Topic checks if the given user has access to the presented topic when
	// using topic authorization (https://www.rabbitmq.com/access-control.html#topic-authorisation)
	Topic(username, vhost, resource, name, permission, routingKey string) Decision
}
```

Start a web server using your authenticator and the http router provided
by the `rabbitmqauth.AuthServer.NewRouter()` function like:

```go
package main

import (
	"fmt"
	"net/http"
	"time"

	auth "github.com/jandelgado/rabbitmq-http-auth/pkg"
)

const httpReadTimeout = 10 * time.Second
const httpWriteTimeout = 10 * time.Second

func main() {
	authenticator := NewLogInterceptingAuthenticator(DemoAuthenticator{})
	s := auth.NewAuthServer(authenticator)

	srv := &http.Server{
		Handler:      s.NewRouter(),
		Addr:         fmt.Sprintf(":%d", 8000),
		WriteTimeout: httpWriteTimeout,
		ReadTimeout:  httpReadTimeout,
	}

	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
```

Have a look at the [example](cmd/example) for a complete example.

## Test it

Start the example by running `make build && make run` and then test the service
by issueing POST requests to the `User` endpoint , for example:

```sh
$ curl  -XPOST localhost:8000/auth/user -d "username=guest&password=test"
allow [management administrator demo]
$ curl  -XPOST localhost:8000/auth/user -d "username=john&password=test"
deny
```

Since the `DemoAuthenticator` only allows the `guest` user (but with any
password), this is the expected result.

## Test with RabbitMQ

A docker-compose file is provided which sets up a RabbitMQ broker with the
authentication service configured. To test it, run:

```sh
$ cd demo && docker-compose up
```

Then in another console, try to publish a message using [rabtap](TODO)
```sh
$  echo "hello" | rabtap pub --uri amqp://guest:123@localhost:5672 --exchange amq.topic --routingkey "#"
```

In the docker-compose log, should see the authenticator logging the request:
```
auth-http_1  | 2021/04/18 21:28:01 auth user(u=guest) -> allow [management administrator demo]
```

As the `DemoAuthenticator` allows any password for the guest user, you can 
try to change the password in the `rabtap` command or try to login on the 
[management console](http://localhost:15672) with any password.

## Author & License

(c) Copyright 2021 by Jan Delgado. Licence: MIT

