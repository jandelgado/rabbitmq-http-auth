// rabbitmq-http-auth - log authentication requests
// (c) copyright 2021 by Jan Delgado
package rabbitmqauth

import (
	"log"
)

type LogInterceptingAuthenticator struct {
	authenticator Authenticator
}

func NewLogInterceptingAuthenticator(authenticator Authenticator) Authenticator {
	return LogInterceptingAuthenticator{authenticator}
}

func (s LogInterceptingAuthenticator) String() string {
	return "LogInterceptingAuthenticator"
}

func (s LogInterceptingAuthenticator) User(username, password string) (Decision, string) {
	res, tags := s.authenticator.User(username, password)
	log.Printf("auth user(u=%s) -> %v [%s]", username, res, tags)
	return res, tags
}

func (s LogInterceptingAuthenticator) VHost(username, vhost, ip string) Decision {
	res := s.authenticator.VHost(username, vhost, ip)
	log.Printf("auth vhost(u=%s,v=%s,i=%s) -> %v", username, vhost, ip, res)
	return res
}

func (s LogInterceptingAuthenticator) Resource(username, vhost, resource, name, permission string) Decision {
	res := s.authenticator.Resource(username, vhost, resource, name, permission)
	log.Printf("auth resource(u=%s,v=%s,r=%s,n=%s,p=%s) -> %v",
		username, vhost, resource, name, permission, res)
	return res
}

func (s LogInterceptingAuthenticator) Topic(username, vhost, resource, name, permission, routing_key string) Decision {
	res := s.authenticator.Topic(username, vhost, resource, name, permission, routing_key)
	log.Printf("auth topic(u=%s,v=%s,r=%s,n=%s,p=%s,k=%s) -> %v",
		username, vhost, resource, name, permission, routing_key, res)
	return res
}
