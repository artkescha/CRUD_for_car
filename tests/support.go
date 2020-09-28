package tests

import (
	"encoding/json"
	"fmt"
	"github.com/manyminds/api2go"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
)

type testClient struct {
	baseURL string
	api     *api2go.API
}

func (c testClient) Do(req testRequest) (*httptest.ResponseRecorder, error) {
	request, err := http.NewRequest(req.method, c.baseURL, strings.NewReader(req.jsonBody))
	if err != nil {
		return &httptest.ResponseRecorder{}, err
	}
	response := httptest.NewRecorder()
	c.api.Handler().ServeHTTP(response, request)
	return response, err
}

type testRequest struct {
	method   string
	jsonBody string
}

type TestCase struct {
	Name     string
	request  testRequest
	response *httptest.ResponseRecorder
	err      error
}

func areEqualJSON(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	fmt.Println(o1)
	fmt.Println(o2)

	return reflect.DeepEqual(o1, o2), nil
}
