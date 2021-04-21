// rabbitmq-http-auth - demo implementation of the Auth interface,
// which allows the user guest with ANY password to be authorized for
// everything.
// (c) copyright 2021 by Jan Delgado
package main

import (
	auth "github.com/jandelgado/rabbitmq-http-auth/pkg"
)

type DemoAuth struct{}

func (s DemoAuth) String() string {
	return "DemoAuth"
}

func (s DemoAuth) Resource(username, vhost, resource, name, permission string) auth.Decision {
	return true
}

func (s DemoAuth) User(username, password string) (auth.Decision, string) {
	return username == "guest", "management administrator demo"
}

func (s DemoAuth) VHost(username, vhost, ip string) auth.Decision {
	return true
}

func (s DemoAuth) Topic(username, vhost, resource, name, permission, routingKey string) auth.Decision {
	return true
}
