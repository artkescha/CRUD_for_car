package resource

import "github.com/artkescha/CRUD_for_car/model"

type Cars []*model.Car

func (c Cars) Len() int           { return len(c) }
func (c Cars) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c Cars) Less(i, j int) bool { return c[i].ID < c[j].ID }
