// rabbitmq-http-auth - http router
// (c) copyright 2021 by Jan Delgado
package rabbitmqauth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const httpReadTimeout = 15 * time.Second
const httpWriteTimeout = 45 * time.Second

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

func (s *AuthServer) NewRouter() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/auth/user", s.userHandler).Methods("POST")
	router.HandleFunc("/auth/vhost", s.vhostHandler).Methods("POST")
	router.HandleFunc("/auth/resource", s.resourceHandler).Methods("POST")
	router.HandleFunc("/auth/topic", s.topicHandler).Methods("POST")
	return router
}
