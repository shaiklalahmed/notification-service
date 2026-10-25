package service

import database "github.com/shaiklalahmed/notification-service/db"

type DatabaseLayer[T any] interface {
	GetValue() []T
}

type Service[T int | string] struct {
	db DatabaseLayer[T]
}

func NewServiceLayer[T int | string](db database.DB[T]) Service[T] {
	return Service[T]{
		db: db,
	}
}

func (s Service[T]) GetValueSer() []T {
	return s.db.GetValue()
}

//testin phase
