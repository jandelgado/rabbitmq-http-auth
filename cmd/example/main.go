// rabbitmq-http-auth - exmaple rabbitmq http auth backend service
// (c) copyright 2021 by Jan Delgado
package main

import (
	"fmt"
	"net/http"
	"time"

	rabbitmqauth "github.com/jandelgado/rabbitmq-http-auth/pkg"
)

const httpReadTimeout = 10 * time.Second
const httpWriteTimeout = 10 * time.Second

func main() {
	auth := NewLogInterceptingAuth(DemoAuth{})
	service := rabbitmqauth.NewAuthService(auth)

	server := &http.Server{
		Handler:      service.NewRouter(),
		Addr:         fmt.Sprintf(":%d", 8000),
		WriteTimeout: httpWriteTimeout,
		ReadTimeout:  httpReadTimeout,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
