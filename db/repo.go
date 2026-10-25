package database

type DB[T any] struct {
	val []T
}

func NewService[T int | string](val ...T) DB[T] {
	return DB[T]{
		val: val,
	}
}

func (db DB[T]) GetValue() []T {
	return db.val
}
