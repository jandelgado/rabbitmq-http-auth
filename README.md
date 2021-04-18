# RabbitMQ HTTP Auth Backend in Go

Package and example service to build a RabbitMQ HTTP service for use with
the RabbitMQ "HTTP Auth Backend" (actually it is an AuthN/AuthZ backend).

For details see https://github.com/rabbitmq/rabbitmq-server/tree/master/deps/rabbitmq_auth_backend_http

## Build your own service

TODO

## Testing it

After starting the demo app either manually or by running 
`docker docker run --rm -ti -p8000:8000 rabbitmq-http-auth:latest`, we
can test the service by issueing POST requests to the `User` endpoint , 
for example:

```sh
$ curl  -XPOST localhost:8000/auth/user -d "username=guest&pasword=test"
allow [management administrator demo]
$ curl  -XPOST localhost:8000/auth/user -d "username=john&pasword=test"
deny
```

Since the `DemoAuthenticator` only allows the `guest` user, this is the
expected result.

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

