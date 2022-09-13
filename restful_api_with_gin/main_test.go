package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPinPong 測試伺服器有無回應
func TestPinPong(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/pin")
	if err != nil {
		t.Fatalf("TestPinPong error:[%v]", err)
		t.FailNow()
	}
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode, "http status code")

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	type respPong struct {
		Message string `json:"message"`
	}
	var pong respPong
	err = json.Unmarshal(body, &pong)
	assert.Nil(t, err)
	// fmt.Printf("message: %v\n", pong.Message)
	assert.Equal(t, "pong", pong.Message)
}

// TestGetUser1 測試能否正確取得user1的資料
func TestGetUser1(t *testing.T) {
	url_expected := map[string]string{
		"http://localhost:8080/user/1": `{"name":"mike"}`,
		"http://localhost:8080/user/2": `{"name":"joe"}`,
	}

	for url, expected := range url_expected {
		resp, err := http.Get(url)
		assert.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, 200, resp.StatusCode)

		body, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)
		assert.Equal(t, expected, string(body))
	}

}