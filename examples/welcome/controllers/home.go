package controllers

import (
	"fmt"
	. "github.com/xshifty/gonaut"
	"net/http"
)

type Home struct {
	BaseController
}

//gonaut:routing get /
func (h Home) GetIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	if _, err := w.Write(String("Index: Hello world!").Bytes()); err != nil {
		fmt.Printf("Can't write response: %s\n", err)
	}
}

//gonaut:routing get /login
func (h Home) GetLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	if _, err := w.Write(String("Login: Hello world!").Bytes()); err != nil {
		fmt.Printf("Can't write response: %s\n", err)
	}
}
