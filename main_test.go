package main_test

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetPing_OK(t *testing.T) {

	req, err := http.NewRequest("GET", "http://localhost:7007/mbu/admin/ping", nil)
	res, err := http.DefaultClient.Do(req)

	assert := assert.New(t)
	assert.Nil(err)
	statusCode := res.StatusCode
	assert.Equal(200, statusCode)

	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(err)
	assert.Equal("pong", string(body))
}
