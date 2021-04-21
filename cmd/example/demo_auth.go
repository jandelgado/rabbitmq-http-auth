// rabbitmq-http-auth - demo implementation of the Auth interface,
// which allows the user guest with ANY password to be authorized for
// everything.
// (c) copyright 2021 by Jan Delgado
package main

import (
	rabbitmqauth "github.com/jandelgado/rabbitmq-http-auth/pkg"
)

type DemoAuth struct{}

func (s DemoAuth) String() string {
	return "DemoAuth"
}

func (s DemoAuth) Resource(username, vhost, resource, name, permission string) rabbitmqauth.Decision {
	return rabbitmqauth.Allow
}

func (s DemoAuth) User(username, password string) (rabbitmqauth.Decision, string) {
	return username == "guest", "management administrator demo"
}

func (s DemoAuth) VHost(username, vhost, ip string) rabbitmqauth.Decision {
	return rabbitmqauth.Allow
}

func (s DemoAuth) Topic(username, vhost, resource, name, permission, routingKey string) rabbitmqauth.Decision {
	return rabbitmqauth.Allow
}
