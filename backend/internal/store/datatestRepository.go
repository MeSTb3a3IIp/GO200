package store

import "apiserver/model"

type DatatestRepository struct {
	store     *Store
	datatests map[int]*model.Datatest
}
