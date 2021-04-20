package rabbitmqauth

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type MockAuthenticator struct{}

func (s MockAuthenticator) String() string {
	return "MockAuthenticator"
}

func (s MockAuthenticator) User(username, password string) (Decision, string) {
	return username == "a" && password == "b", "management"
}

func (s MockAuthenticator) VHost(username, vhost, ip string) Decision {
	return username == "a" && vhost == "b" && ip == "c"
}

func (s MockAuthenticator) Resource(username, vhost, resource, name, permission string) Decision {
	return username == "a" && vhost == "b" && resource == "c" && name == "d" && permission == "e"
}

func (s MockAuthenticator) Topic(username, vhost, resource, name, permission, routingKey string) Decision {
	return username == "a" && vhost == "b" && resource == "c" && name == "d" && permission == "e" && routingKey == "f"
}

// httpPost does a http POST request to the given url with the given request
// body
func httpPost(url, reqBody string) (string, error) {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", bytes.NewBufferString(reqBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

func TestPostHandlerReturns405OnGET(t *testing.T) {

	cases := []string{"/auth/user", "/auth/vhost", "/auth/resource", "/auth/topic"}

	for _, path := range cases {
		auth := NewAuthServer(MockAuthenticator{})
		handlerFunc := auth.NewRouter()

		req, err := http.NewRequest("GET", path, nil)
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		handlerFunc.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
		assert.Equal(t, "405 method not allowed", rr.Body.String())
	}
}

func TestUserRequestIsAllowedOnlyOnValidRequests(t *testing.T) {
	auth := NewAuthServer(MockAuthenticator{})
	srv := httptest.NewServer(auth.NewRouter())
	defer srv.Close()

	cases := []struct {
		request  string
		response string
	}{
		{"username=a&password=b", "allow [management]"},
		{"username=a&password=X", "deny"},
		{"", "deny"},
		{"invalid", "deny"},
	}
	for _, tc := range cases {
		resp, err := httpPost(srv.URL+"/auth/user", tc.request)
		assert.NoError(t, err)
		assert.Equal(t, tc.response, resp)
	}
}

func TestVHostRequestIsAllowedOnlyOnValidRequests(t *testing.T) {
	auth := NewAuthServer(MockAuthenticator{})
	srv := httptest.NewServer(auth.NewRouter())
	defer srv.Close()

	cases := []struct {
		request  string
		response string
	}{
		{"username=a&vhost=b&ip=c", "allow"},
		{"username=a&vhost=b&ip=X", "deny"},
		{"", "deny"},
		{"invalid", "deny"},
	}
	for _, tc := range cases {
		resp, err := httpPost(srv.URL+"/auth/vhost", tc.request)
		assert.NoError(t, err)
		assert.Equal(t, tc.response, resp, tc)
	}
}

func TestResourceRequestIsAllowedOnlyOnValidRequests(t *testing.T) {
	auth := NewAuthServer(MockAuthenticator{})
	srv := httptest.NewServer(auth.NewRouter())
	defer srv.Close()

	cases := []struct {
		request  string
		response string
	}{
		{"username=a&vhost=b&resource=c&name=d&permission=e", "allow"},
		{"username=a&vhost=b&resource=c&name=d&permission=X", "deny"},
		{"", "deny"},
		{"invalid", "deny"},
	}
	for _, tc := range cases {
		resp, err := httpPost(srv.URL+"/auth/resource", tc.request)
		assert.NoError(t, err)
		assert.Equal(t, tc.response, resp, tc)
	}
}

func TestTopicRequestIsAllowedOnlyOnValidRequests(t *testing.T) {
	auth := NewAuthServer(MockAuthenticator{})
	srv := httptest.NewServer(auth.NewRouter())
	defer srv.Close()

	cases := []struct {
		request  string
		response string
	}{
		{"username=a&vhost=b&resource=c&name=d&permission=e&routing_key=f", "allow"},
		{"username=a&vhost=b&resource=c&name=d&permission=e&routing_key=X", "deny"},
		{"", "deny"},
		{"invalid", "deny"},
	}
	for _, tc := range cases {
		resp, err := httpPost(srv.URL+"/auth/topic", tc.request)
		assert.NoError(t, err)
		assert.Equal(t, tc.response, resp, tc)
	}
}
