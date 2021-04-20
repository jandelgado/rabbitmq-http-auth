// rabbitmq-http-auth - demo implementation of the Authenticator interface,
// which allows the user guest with ANY password to be authorized for
// everything.
// (c) copyright 2021 by Jan Delgado
package main

import (
	auth "github.com/jandelgado/rabbitmq-http-auth/pkg"
)

type DemoAuthenticator struct{}

func (s DemoAuthenticator) String() string {
	return "DemoAuthenticator"
}

func (s DemoAuthenticator) Resource(username, vhost, resource, name, permission string) auth.Decision {
	return true
}

func (s DemoAuthenticator) User(username, password string) (auth.Decision, string) {
	return username == "guest", "management administrator demo"
}

func (s DemoAuthenticator) VHost(username, vhost, ip string) auth.Decision {
	return true
}

func (s DemoAuthenticator) Topic(username, vhost, resource, name, permission, routingKey string) auth.Decision {
	return true
}
