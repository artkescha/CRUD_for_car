package resource

import (
	"fmt"
	"github.com/artkescha/CRUD_for_car/model"
	"github.com/artkescha/CRUD_for_car/storage"
	"github.com/asaskevich/govalidator"
	"github.com/manyminds/api2go"
	"net/http"
	"sort"
	"strconv"
)

// UserResource for api2go routes
type CarResource struct {
	Storage storage.Storage
}

// FindAll method implements api2go.FindAll interface
func (c CarResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	cars, err := c.Storage.GetAll()
	if err != nil {
		if err == storage.ErrNotFound {
			return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
		}
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusInternalServerError)
	}

	sortField, ok := r.QueryParams["sort"]
	//sort only for id
	if ok {
		if ok && sortField[0] == "id" {
			sort.Sort(Cars(cars))
		} else {
			err := fmt.Errorf("sort field %s not allowed", sortField)
			return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusBadRequest)
		}
	}
	return &Response{Res: cars, Code: http.StatusOK}, nil
}

// PaginatedFindAll method implements api2go.PaginatedFindAll interface
func (c CarResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	offset, ok := r.Pagination["offset"]
	if !ok {
		err := fmt.Errorf("no offset set in request")
		return 0, &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusBadRequest)
	}
	limit, ok := r.Pagination["limit"]
	if !ok {
		err := fmt.Errorf("no limit set in request")
		return 0, &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusBadRequest)
	}
	cars, err := c.Storage.GetAll()
	if err != nil {

		if err == storage.ErrNotFound {
			return 0, &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
		}

		return 0, &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusInternalServerError)
	}

	limitI, err := strconv.ParseUint(limit, 10, 64)
	if err != nil {
		return 0, &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusBadRequest)
	}
	offsetI, err := strconv.ParseUint(offset, 10, 64)
	if err != nil {
		return 0, &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusBadRequest)
	}

	total := uint64(len(cars))

	sortField, ok := r.QueryParams["sort"]
	//sort only for id
	if ok {
		if ok && sortField[0] == "id" {
			sort.Sort(Cars(cars))
		} else {
			err := fmt.Errorf("sort field %s not allowed", sortField)
			return 0, &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusBadRequest)
		}
	}

	if offsetI+limitI <= total {
		return uint(total), &Response{Res: cars[offsetI : offsetI+limitI], Code: http.StatusOK}, nil
	}
	result := make([]model.Car, 0, limitI)
	for i := offsetI; i <= offsetI+limitI && i < total; i++ {
		result = append(result, *cars[i])
	}
	return uint(total), &Response{Res: result, Code: http.StatusOK}, nil
}

// FindOne method implements api2go.ResourceGetter interface
func (c CarResource) FindOne(id string, r api2go.Request) (api2go.Responder, error) {
	car, err := c.Storage.GetOne(id)
	if err != nil {
		if err == storage.ErrNotFound {
			return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
		}
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusInternalServerError)
	}
	return &Response{Res: car}, nil
}

// Create method implements api2go.ResourceCreator interface
func (c CarResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	car, ok := obj.(model.Car)
	if !ok {
		return &Response{}, api2go.NewHTTPError(fmt.Errorf("invalid instance given"), "invalid instance given", http.StatusBadRequest)
	}
	valid, err := govalidator.ValidateStruct(car)

	if !valid || err != nil {
		return &Response{}, api2go.NewHTTPError(fmt.Errorf("vendor and model is required"), "vendor and model is required", http.StatusBadRequest)
	}

	id, err := c.Storage.Insert(&car)
	if err != nil {
		return &Response{}, fmt.Errorf("insert car failed: %s", err)
	}
	car.SetID(id)

	return &Response{Res: car, Code: http.StatusCreated}, nil
}

// Delete method implements api2go.ResourceDeleter interface
func (c CarResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.Storage.Delete(id)
	if err != nil {
		if err == storage.ErrNotFound {
			return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
		}
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusInternalServerError)
	}
	return &Response{Code: http.StatusNoContent}, err
}

// Update implements api2go.ResourceUpdater interface
func (c CarResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	car, ok := obj.(model.Car)
	if !ok {
		return &Response{}, api2go.NewHTTPError(fmt.Errorf("invalid instance given"), "invalid instance given", http.StatusBadRequest)
	}
	valid, err := govalidator.ValidateStruct(car)

	if !valid || err != nil {
		return &Response{}, api2go.NewHTTPError(fmt.Errorf("vendor and model is required"), "vendor and model is required", http.StatusBadRequest)
	}

	err = c.Storage.Update(car)
	if err != nil {
		if err == storage.ErrNotFound {
			return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
		}
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusInternalServerError)
	}
	return &Response{Res: car, Code: http.StatusNoContent}, err
}
