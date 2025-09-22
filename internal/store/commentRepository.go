package store

import "apiserver/model"

type CommentRepository struct {
	store    *Store
	comments map[int]*model.Comment
}
