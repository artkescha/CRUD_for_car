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
func TestCarsHandler_DeleteCar(t *testing.T) {
	var deleteCarsCases = []TestCase{
		{
			Name:    "delete car ok case",
			request: testRequest{method: http.MethodDelete, jsonBody: ``},
			response: &httptest.ResponseRecorder{Code: http.StatusNoContent,

				Body: bytes.NewBuffer([]byte(``))},
		},
	}

	ctrl := gomock.NewController(t)
	mockRepo := storage.NewMockStorage(ctrl)

	api := api2go.NewAPIWithBaseURL("v1", "http://localhost:31415")
	api.AddResource(model.Car{}, resource.CarResource{Storage: mockRepo})

	client := testClient{
		baseURL: "/v1/cars/0",
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

	//insert before read
	for _, testCase := range createCarsCases {
		client.Do(testCase.request)
	}

	for _, testCase := range deleteCarsCases {
		mockRepo.EXPECT().Delete("0").Return(testCase.err)

		response, err := client.Do(testCase.request)
		assert.Nil(t, err, "error must be nil")
		assert.Equal(t, testCase.response.Code, response.Code)
	}
}
