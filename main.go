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
	"net/http"
	"os"
)

func main() {
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "81"
		zapLogger.Info("set default port value:", zap.String("", port))
	}

	api := api2go.NewAPIWithResolver("v1", &resolver.RequestURL{Port: port})
	api.AddResource(model.Car{}, resource.CarResource{Storage: storage.NewStorage()})

	zapLogger.Info("Listening on port:", zap.String("",port))

	handler, ok := api.Handler().(*httprouter.Router)
	if !ok {
		zapLogger.Fatal("router not implemented handler interface")
	}

	middleware := middlewares.AccessLogger{ZapLogger: zapLogger.Sugar()}
	api.UseMiddleware(middleware.AccessLogMiddleware)

	http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}
