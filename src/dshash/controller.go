package dshash

import (
	"appengine"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

type WebContextHandler func(appengine.Context, http.ResponseWriter, *http.Request, httprouter.Params)

func HandlerWithContext(h WebContextHandler) httprouter.Handle {
	f := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		h(appengine.NewContext(r), w, r, ps)
	}

	return httprouter.Handle(f)
}
func init() {
	router := httprouter.New()
	router.GET("/locations/:handler", HandlerWithContext(getHandler))
	router.POST("/locations", HandlerWithContext(postHandler))

	http.Handle("/", router)
}

func getHandler(c appengine.Context, w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	person := &Person{}
	person.Handler = "chischaschos"
	bytes, err := person.Marshal()

	person.Locations = RetrieveLocations(person)

	if err != nil {
		panic(err)
	}

	_, err = w.Write(bytes)

	if err != nil {
		panic(err)
	}
}

func postHandler(c appengine.Context, w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	defer r.Body.Close()
	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	person := &Person{}
	err = person.Unmarshal(bodyBytes)

	if err != nil {
		panic(err)
	}

	storedPerson := StorePerson(person)

	bytes, err := storedPerson.Marshal()

	if err != nil {
		panic(err)
	}

	_, err = w.Write(bytes)

	if err != nil {
		panic(err)
	}
}
