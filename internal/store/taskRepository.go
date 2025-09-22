package store

import "apiserver/model"

type TaskRepository struct {
	store *Store
	tasks map[int]*model.Task
}
