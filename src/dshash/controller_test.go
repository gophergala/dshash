package dshash

import (
	"appengine/aetest"
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootRoute(t *testing.T) {
	assert := assert.New(t)

	c, e := aetest.NewContext(nil)
	assert.Nil(e)
	defer c.Close()

	r, e := http.NewRequest("GET", "localhost?handler=chischaschos", nil)
	assert.Nil(e)

	w := httptest.NewRecorder()
	getHandler(c, w, r)
	assert.Equal(w.Code, 200)

	person := Person{}
	e = person.Unmarshal(w.Body.Bytes())

	assert.Nil(e)
	assert.Equal(person.Handler, "chischaschos")
	assert.Empty(person.Locations)

	person.Locations = []string{"Somewhere"}

	personBytes, e := person.Marshal()
	assert.Nil(e)

	r, e = http.NewRequest("POST", "localhost", bytes.NewBuffer(personBytes))
	assert.Nil(e)

	w = httptest.NewRecorder()
	postHandler(c, w, r)
	assert.Equal(w.Code, 200)

	person.Unmarshal(w.Body.Bytes())

	assert.Nil(e)
	assert.Equal(person.Handler, "chischaschos")
	assert.Equal(person.Locations, []string{"Somewhere"})

}
