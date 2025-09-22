package store

import (
	"apiserver/model"
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
}
