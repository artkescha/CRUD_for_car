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
func TestCarsHandler_ReadAllCars(t *testing.T) {
	var readCarsCases = []TestCase{
		{
			Name:    "read all cars ok case",
			request: testRequest{method: "GET", jsonBody: ``},
			response: &httptest.ResponseRecorder{Code: http.StatusOK,

				Body: bytes.NewBuffer([]byte(`{
    "data": [
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
        }],
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

	mockRepo.EXPECT().Insert(car).Return(car.ID, nil)
	mockRepo.EXPECT().GetAll().Return([]*model.Car{car}, nil)

	//insert before read
	for _, testCase := range createCarsCases {
		client.Do(testCase.request)
	}

	for _, testCase := range readCarsCases {
		response, err := client.Do(testCase.request)
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, testCase.response.Code, response.Code)
		isEqual, err := areEqualJSON(testCase.response.Body.String(), response.Body.String())
		assert.Nil(t, err, "error must be nil")
		assert.True(t, isEqual)
	}
}
