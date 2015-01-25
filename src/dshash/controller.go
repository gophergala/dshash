package dshash

import (
	"appengine"
	"fmt"
	"net/http"
)

type WebContextHandler func(appengine.Context, http.ResponseWriter, *http.Request)

func HandlerWithContext(h WebContextHandler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		h(appengine.NewContext(r), w, r)
	}

	return http.HandlerFunc(f)
}

func init() {
	http.Handle("/", HandlerWithContext(handler))
}

func handler(c appengine.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
