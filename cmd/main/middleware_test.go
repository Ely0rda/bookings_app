package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNosurf(t *testing.T) {
	var myH myHandler
	h := NoSurf(&myH)
	switch v := h.(type) {
	case http.Handler:
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is %T", v))
	}
}

func TestLoadd(t *testing.T) {
	var myH myHandler
	h := SessionLoad(&myH)
	switch v := h.(type) {
	case http.Handler:
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is %T", v))
	}
}
