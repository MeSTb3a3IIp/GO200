package store

import "apiserver/model"

type TagRepository struct {
	store *Store
	tags  map[int]*model.Tag
}
