package storage

import (
	"errors"
	"github.com/artkescha/CRUD_for_car/model"
	"github.com/satori/go.uuid"
	"sync"
)

var (
	ErrNotFound = errors.New("not found")
)

type storage struct {
	syn  sync.RWMutex
	cars map[string]*model.Car
}

func NewStorage() *storage {
	return &storage{cars: make(map[string]*model.Car)}
}

func (s storage) GetAll() ([]*model.Car, error) {
	target := make([]*model.Car, len(s.cars))
	s.syn.RLock()
	defer s.syn.RUnlock()
	// Copy from the original map to the target map
	index := 0
	for _, car := range s.cars {
		target[index] = car
		index++
	}
	return target, nil
}

func (s storage) GetOne(id string) (model.Car, error) {
	s.syn.RLock()
	defer s.syn.RUnlock()
	car, exist := s.cars[id]
	if !exist {
		return model.Car{}, ErrNotFound
	}
	return *car, nil
}

func (s *storage) Insert(car *model.Car) (string, error) {
	s.syn.Lock()
	defer s.syn.Unlock()
	id := uuid.NewV4().String()
	s.cars[id] = car
	return id, nil
}

func (s *storage) Delete(id string) error {
	s.syn.Lock()
	defer s.syn.Unlock()
	_, exists := s.cars[id]
	if !exists {
		return ErrNotFound
	}
	delete(s.cars, id)
	return nil
}

func (s *storage) Update(c model.Car) error {
	s.syn.Lock()
	defer s.syn.Unlock()
	_, exists := s.cars[c.ID]
	if !exists {
		return ErrNotFound
	}
	s.cars[c.ID] = &c

	return nil
}
