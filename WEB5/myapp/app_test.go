package myapp

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestIndex( t *testing.T)
{
	assert := assert.New(t)

	httptest.NewServer(NewHandler())

	defer ts.Close()

	http.Get(ts.URL)

	resp , err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(httpStatusOK, resp.StatusCode)
}