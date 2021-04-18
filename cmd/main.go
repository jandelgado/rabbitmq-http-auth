// rabbitmq-http-auth - exmple authentication service using a demo authenticator
// (c) copyright 2021 by Jan Delgadoauthenticator}
package main

import (
	"fmt"
	"net/http"
	"time"

	auth "github.com/jandelgado/rabbitmq-http-auth/pkg"
)

const httpReadTimeout = 10 * time.Second
const httpWriteTimeout = 10 * time.Second

func main() {
	a := DemoAuthenticator{}
	authenticator := auth.NewLogInterceptingAuthenticator(a)
	s := auth.NewAuthServer(authenticator)

	srv := &http.Server{
		Handler:      s.NewRouter(),
		Addr:         fmt.Sprintf(":%d", 8000),
		WriteTimeout: httpWriteTimeout,
		ReadTimeout:  httpReadTimeout,
	}

	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
