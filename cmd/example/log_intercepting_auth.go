// rabbitmq-http-auth - log authentication requests
// (c) copyright 2021 by Jan Delgado
package main

import (
	"log"

	auth "github.com/jandelgado/rabbitmq-http-auth/pkg"
)

type LogInterceptingAuthenticator struct {
	authenticator auth.Authenticator
}

func NewLogInterceptingAuthenticator(authenticator auth.Authenticator) auth.Authenticator {
	return LogInterceptingAuthenticator{authenticator}
}

func (s LogInterceptingAuthenticator) String() string {
	return "LogInterceptingAuthenticator"
}

func (s LogInterceptingAuthenticator) User(username, password string) (auth.Decision, string) {
	res, tags := s.authenticator.User(username, password)
	log.Printf("auth user(u=%s) -> %v [%s]", username, res, tags)
	return res, tags
}

func (s LogInterceptingAuthenticator) VHost(username, vhost, ip string) auth.Decision {
	res := s.authenticator.VHost(username, vhost, ip)
	log.Printf("auth vhost(u=%s,v=%s,i=%s) -> %v", username, vhost, ip, res)
	return res
}

func (s LogInterceptingAuthenticator) Resource(username, vhost, resource, name, permission string) auth.Decision {
	res := s.authenticator.Resource(username, vhost, resource, name, permission)
	log.Printf("auth resource(u=%s,v=%s,r=%s,n=%s,p=%s) -> %v",
		username, vhost, resource, name, permission, res)
	return res
}

func (s LogInterceptingAuthenticator) Topic(username, vhost, resource, name, permission, routingKey string) auth.Decision {
	res := s.authenticator.Topic(username, vhost, resource, name, permission, routingKey)
	log.Printf("auth topic(u=%s,v=%s,r=%s,n=%s,p=%s,k=%s) -> %v",
		username, vhost, resource, name, permission, routingKey, res)
	return res
}
