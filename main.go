package main

import (
	"fmt"
	"github.com/artkescha/CRUD_for_car/middlewares"
	"github.com/artkescha/CRUD_for_car/model"
	"github.com/artkescha/CRUD_for_car/resolver"
	"github.com/artkescha/CRUD_for_car/resource"
	"github.com/artkescha/CRUD_for_car/storage"
	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
)

func main() {
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "81"
		zapLogger.Info("set default port value")
	}

	api := api2go.NewAPIWithResolver("v1", &resolver.RequestURL{Port: port})
	api.AddResource(model.Car{}, resource.CarResource{Storage: storage.NewStorage()})

	log.Printf("Listening on port :%s", port)
	handler := api.Handler().(*httprouter.Router)

	middleware := middlewares.AccessLogger{ZapLogger: zapLogger.Sugar()}
	api.UseMiddleware(middleware.AccessLogMiddleware)

	http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}
