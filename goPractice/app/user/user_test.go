package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Message struct {
	Message string
}

func TestSignUp(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/signup",
		strings.NewReader(`{"Email" : "qs553@nate.com" , "Password" : "ehdwk1", "Nickname" : "dongja2"}`))

	mux := MakeHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	var message Message

	err := json.NewDecoder(res.Body).Decode(&message)
	assert.Nil(err)
	assert.Equal("가입 성공", message.Message)

}

func TestSignIn(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/signin",
		strings.NewReader(`{"Email" : "qs553@nate.com" , "Password" : "ehdwk1"}`))

	mux := MakeHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
}
