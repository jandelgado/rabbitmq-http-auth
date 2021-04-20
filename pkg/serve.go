// rabbitmq-http-auth - http router
// (c) copyright 2021 by Jan Delgado
package rabbitmqauth

import (
	"fmt"
	"net/http"
)

type AuthServer struct {
	authenticator Authenticator
}

func NewAuthServer(authenticator Authenticator) AuthServer {
	return AuthServer{authenticator}
}

func (s Decision) String() string {
	if s {
		return "allow"
	}
	return "deny"
}

func validatePostArgs(args []string, r *http.Request) map[string]string {
	result := map[string]string{}
	for _, s := range args {
		result[s] = r.PostFormValue(s)
	}
	return result
}

func (s *AuthServer) userHandler(w http.ResponseWriter, r *http.Request) {
	args := validatePostArgs([]string{"username", "password"}, r)
	res, tags := s.authenticator.User(args["username"], args["password"])
	if res {
		fmt.Fprintf(w, "%s [%s]", res, tags)
	} else {
		fmt.Fprintf(w, "%s", res)
	}
}

func (s *AuthServer) vhostHandler(w http.ResponseWriter, r *http.Request) {
	args := validatePostArgs([]string{"username", "vhost", "ip"}, r)
	res := s.authenticator.VHost(args["username"], args["vhost"], args["ip"])
	fmt.Fprintf(w, "%s", res)
}

func (s *AuthServer) topicHandler(w http.ResponseWriter, r *http.Request) {
	args := validatePostArgs([]string{"username", "vhost", "resource", "name", "permission", "routing_key"}, r)
	res := s.authenticator.Topic(args["username"], args["vhost"],
		args["resource"], args["name"], args["permission"], args["routing_key"])
	fmt.Fprintf(w, "%s", res)
}

func (s *AuthServer) resourceHandler(w http.ResponseWriter, r *http.Request) {
	args := validatePostArgs([]string{"username", "vhost", "resource", "name", "permission"}, r)
	res := s.authenticator.Resource(args["username"], args["vhost"],
		args["resource"], args["name"], args["permission"])
	fmt.Fprintf(w, "%s", res)
}

func postHandler(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handler(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			_, _ = w.Write([]byte("405 method not allowed"))
		}
	}
}

func (s *AuthServer) NewRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/auth/user", postHandler(s.userHandler))
	router.HandleFunc("/auth/vhost", postHandler(s.vhostHandler))
	router.HandleFunc("/auth/resource", postHandler(s.resourceHandler))
	router.HandleFunc("/auth/topic", postHandler(s.topicHandler))
	return router
}
