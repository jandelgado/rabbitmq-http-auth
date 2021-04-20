// rabbitmq-http-auth - Authenicator interface
// (c) copyright 2021 by Jan Delgado
package rabbitmqauth

type Decision bool

const (
	Allow Decision = true
	Deny  Decision = false
)

// Authenticator instances make the actual decisions on authentication
// requests by RabboitMQ. Every function returns the authentication decision,
// which is always Allow or Deny.
//
// See https://github.com/rabbitmq/rabbitmq-server/tree/master/deps/rabbitmq_auth_backend_http
// for a detailed description.
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
