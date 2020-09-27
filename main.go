package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/artkescha/CRUD_for_car/model"
	"github.com/artkescha/CRUD_for_car/resolver"
	"github.com/artkescha/CRUD_for_car/resource"
	"github.com/artkescha/CRUD_for_car/storage"
	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go"
)

func main() {
	port := 81
	api := api2go.NewAPIWithResolver("v1", &resolver.RequestURL{Port: port})
	api.AddResource(model.Car{}, resource.CarResource{Storage: storage.NewStorage()})

	log.Printf("Listening on :%d", port)
	handler := api.Handler().(*httprouter.Router)

	http.ListenAndServe(fmt.Sprintf(":%d", port), handler)

}
