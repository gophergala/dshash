package dshash

import (
	"appengine"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

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
