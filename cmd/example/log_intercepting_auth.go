// rabbitmq-http-auth - log authentication requests
// (c) copyright 2021 by Jan Delgado
package main

import (
	"log"

	rabbitmqauth "github.com/jandelgado/rabbitmq-http-auth/pkg"
)

type LogInterceptingAuth struct {
	auth rabbitmqauth.Auth
}

func NewLogInterceptingAuth(auth rabbitmqauth.Auth) rabbitmqauth.Auth {
	return LogInterceptingAuth{auth}
}

func (s LogInterceptingAuth) String() string {
	return "LogInterceptingAuth"
}

func (s LogInterceptingAuth) User(username, password string) (rabbitmqauth.Decision, string) {
	res, tags := s.auth.User(username, password)
	log.Printf("auth user(u=%s) -> %v [%s]", username, res, tags)
	return res, tags
}

func (s LogInterceptingAuth) VHost(username, vhost, ip string) rabbitmqauth.Decision {
	res := s.auth.VHost(username, vhost, ip)
	log.Printf("auth vhost(u=%s,v=%s,i=%s) -> %v", username, vhost, ip, res)
	return res
}

func (s LogInterceptingAuth) Resource(username, vhost, resource, name, permission string) rabbitmqauth.Decision {
	res := s.auth.Resource(username, vhost, resource, name, permission)
	log.Printf("auth resource(u=%s,v=%s,r=%s,n=%s,p=%s) -> %v",
		username, vhost, resource, name, permission, res)
	return res
}

func (s LogInterceptingAuth) Topic(username, vhost, resource, name, permission, routingKey string) rabbitmqauth.Decision {
	res := s.auth.Topic(username, vhost, resource, name, permission, routingKey)
	log.Printf("auth topic(u=%s,v=%s,r=%s,n=%s,p=%s,k=%s) -> %v",
		username, vhost, resource, name, permission, routingKey, res)
	return res
}
