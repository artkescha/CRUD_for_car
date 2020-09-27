package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/artkescha/CRUD_for_car/model"
	"github.com/artkescha/CRUD_for_car/model/status"
	"github.com/artkescha/CRUD_for_car/resource"
	"github.com/artkescha/CRUD_for_car/storage"
	"github.com/golang/mock/gomock"
	"github.com/manyminds/api2go"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type testClient struct {
	baseURL string
	api     *api2go.API
}

func (c testClient) CreateCar(req testRequest) (*httptest.ResponseRecorder, error) {
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

//tests
type TestCase struct {
	Name     string
	request  testRequest
	response *httptest.ResponseRecorder
}

func TestCarsHandler_CreateCar(t *testing.T) {
	var createCarCases = []TestCase{
		{
			Name: "create car ok case",
			request: testRequest{method: "POST", jsonBody: `{
  "data" : {
  	"type" : "cars" , 
     "attributes": {
     	"vendor":"1",
     	"model" : "1",
     	"status": "in transit"
     }
    }
}`},
			response: &httptest.ResponseRecorder{Code: http.StatusCreated,

				Body: bytes.NewBuffer([]byte(`{
    "data": 
        {
            "type": "cars",
            "id": "",
            "attributes": {
                "model": "1",
                "vendor": "1",
                "price": 0,
                "status": "in transit",
                "mileage": 0
            }
        },
    "meta": {
        "author": "artem savarin",
        "license": "MIT",
        "license-url": "https://opensource.org/licenses/MIT"
    }
}`))},
		},
	}

	ctrl := gomock.NewController(t)
	mockRepo := storage.NewMockStorage(ctrl)

	api := api2go.NewAPIWithBaseURL("v1", "http://localhost:31415")
	api.AddResource(model.Car{}, resource.CarResource{Storage: mockRepo})

	client := testClient{
		baseURL: "/v1/cars",
		api:     api,
	}

	modelName := "1"
	vendor := "1"
	car := &model.Car{
		Model:   &modelName,
		Vendor:  &vendor,
		Price:   0,
		Status:  status.InTransit,
		Mileage: 0,
	}

	for _, testCase := range createCarCases {
		mockRepo.EXPECT().Insert(car).Return(car.ID, nil)

		response, err := client.CreateCar(testCase.request)
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, testCase.response.Code, response.Code)
		isEqual, err := areEqualJSON(testCase.response.Body.String(), response.Body.String())
		assert.Nil(t, err, "error must be nil")
		assert.True(t, isEqual)
	}
}
