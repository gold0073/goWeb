package myapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	//indexHandler(res, req)
	mux := NewHTTPHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	data, _ := ioutil.ReadAll(res.Body)

	assert.Equal("Hello World", string(data))

	/*
		data, _ := ioutil.ReadAll(res.Body)


			if res.Code != http.StatusOK {
				t.Fatal("Failed!!", res.Code)
			}
	*/

}

func TestBarPathHandler_WithOutName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	//barHandler(res, req)
	mux := NewHTTPHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	data, _ := ioutil.ReadAll(res.Body)

	assert.Equal("Hello barPark", string(data))
}

func TestFooHandler_WithOutJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)

	//barHandler(res, req)
	mux := NewHTTPHandler()
	mux.ServeHTTP(res, req)

	//assert.Equal(http.StatusOK, res.Code)
	assert.Equal(http.StatusBadRequest, res.Code)

}

func TestFooHandler_WithJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("post", "/foo",
		strings.NewReader(`{"first_name":"jy","last_name":"park","email":"aa.@a.a"}`))

	//barHandler(res, req)
	mux := NewHTTPHandler()
	mux.ServeHTTP(res, req)

	//assert.Equal(http.StatusOK, res.Code)
	//assert.Equal(http.StatusBadRequest, res.Code)
	assert.Equal(http.StatusCreated, res.Code)

	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)

	assert.Nil(err)
	assert.Equal("jy", user.FirstName)
	assert.Equal("park", user.LastName)
	assert.Equal("aa.@a.a", user.Email)

}

func TestNewHTTPPathHandler_WithName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name=Park", nil)

	mux := NewHTTPHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	data, _ := ioutil.ReadAll(res.Body)

	assert.Equal("Hello Park", string(data))
}
