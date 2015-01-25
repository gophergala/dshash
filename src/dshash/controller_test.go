package dshash

import (
	"appengine/aetest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootRoute(t *testing.T) {
	assert := assert.New(t)

	w := httptest.NewRecorder()

	r, e := http.NewRequest("GET", "localhost", nil)
	assert.Nil(e)

	c, e := aetest.NewContext(nil)
	assert.Nil(e)
	defer c.Close()

	handler(c, w, r)
	assert.Equal(w.Code, 200)
}
