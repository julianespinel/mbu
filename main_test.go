package main_test

import (
	"testing"
	"net/http"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
)

func TestGetPing_OK(t *testing.T) {

	req, err := http.NewRequest("GET", "http://localhost:8080/mbu/admin/ping", nil)
	res, err := http.DefaultClient.Do(req)

	assert := assert.New(t)
	assert.Nil(err)
	statusCode := res.StatusCode
	assert.Equal(200, statusCode)

	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(err)
	assert.Equal("pong", string(body))
}
