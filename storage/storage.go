package storage

import "github.com/artkescha/CRUD_for_car/model"

//go:generate mockgen -destination=./repository_mock.go -package=storage . Storage

type Storage interface {
	GetAll() ([]*model.Car, error)
	GetOne(id string) (model.Car, error)
	Insert(c *model.Car) (string, error)
	Delete(id string) error
	Update(c model.Car) error
}
