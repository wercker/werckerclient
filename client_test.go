package werckerclient

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wercker/werckerclient/credentials"
)

type testResult struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func TestClientMakeRequestGET400(t *testing.T) {
	result := []byte(`{"statusCode":400,"error":"Bad Request","message":"result must be one of aborted, unknown, passed, failed","details":[{"message":"result must be one of aborted, unknown, passed, failed","path":"result","type":"any.allowOnly","context":{"valids":"aborted, unknown, passed, failed","key":"result"}}]}`)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		w.Write(result)
	}))
	defer ts.Close()

	config := &Config{Endpoint: ts.URL}
	client := NewClient(config)
	_, err := client.makeRequest("GET", "/", nil)

	require.Error(t, err, "")
	assert.Equal(t, "result must be one of aborted, unknown, passed, failed", err.Error(), "")
}

func TestClientMakeRequestGET200Anon(t *testing.T) {
	result := []byte(`{"key": "some key","value":"some value"}`)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write(result)
	}))
	defer ts.Close()

	config := &Config{Endpoint: ts.URL}
	client := NewClient(config)
	body, err := client.makeRequest("GET", "/", nil)

	require.NoError(t, err, "")
	assert.Equal(t, result, body, "")
}

func TestClientMakeRequestGET200Token(t *testing.T) {
	result := []byte(`{"key": "some key","value":"some value"}`)
	tokenSet := false
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "Bearer secret_token" {
			tokenSet = true
		}
		w.WriteHeader(200)
		w.Write(result)
	}))
	defer ts.Close()

	config := &Config{Endpoint: ts.URL, Credentials: credentials.Token("secret_token")}
	client := NewClient(config)
	body, err := client.makeRequest("GET", "/", nil)

	require.NoError(t, err, "")
	assert.True(t, tokenSet, "")
	assert.Equal(t, result, body, "")
}

func TestClientMakeRequestGET200UsernamePassword(t *testing.T) {
	authSet := false
	result := []byte(`{"key": "some key","value":"some value"}`)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if user, pass, ok := r.BasicAuth(); ok && user == "secret username" && pass == "secret password" {
			authSet = true
		}
		w.WriteHeader(200)
		w.Write(result)
	}))
	defer ts.Close()

	config := &Config{Endpoint: ts.URL, Credentials: credentials.UsernamePassword("secret username", "secret password")}
	client := NewClient(config)
	body, err := client.makeRequest("GET", "/", nil)

	require.NoError(t, err, "")
	assert.True(t, authSet, "")
	assert.Equal(t, result, body, "")
}

func TestClientNoUrlModel(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(`{"hello": "i work"}`))
	}))
	defer ts.Close()
	config := &Config{Endpoint: ts.URL}
	client := NewClient(config)
	c := make(map[string]string)
	template := userTemplates["GetUser"]
	err := client.Do("GET", template, nil, nil, &c)
	assert.NoError(t, err, "should not error")
}
