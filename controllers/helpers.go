package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

func pastForm(r *http.Request, dst interface{}) error {
	// r.ParseForm() populates r.PostForm
	if err := r.ParseForm(); err != nil {
		return err
	}
	dec := schema.NewDecoder()
	if err := dec.Decode(dst, r.PostForm); err != nil {
		return err
	}
	return nil
}
