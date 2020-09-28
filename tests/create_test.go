package tests

import (
	"bytes"
	"github.com/artkescha/CRUD_for_car/model"
	"github.com/artkescha/CRUD_for_car/model/status"
	"github.com/artkescha/CRUD_for_car/resource"
	"github.com/artkescha/CRUD_for_car/storage"
	"github.com/golang/mock/gomock"
	"github.com/manyminds/api2go"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

//tests
var createCarsCases = []TestCase{
	{
		Name: "create car ok case",
		request: testRequest{method: http.MethodPost, jsonBody: `{
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

func TestCarsHandler_CreateCar(t *testing.T) {

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

	mockRepo.EXPECT().Insert(car).Return(car.ID, nil)

	for _, testCase := range createCarsCases {
		response, err := client.Do(testCase.request)
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, testCase.response.Code, response.Code)
		isEqual, err := areEqualJSON(testCase.response.Body.String(), response.Body.String())
		assert.Nil(t, err, "error must be nil")
		assert.True(t, isEqual)
	}
}
