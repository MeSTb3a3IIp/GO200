package store

import (
	"database/sql"
)

type Store struct {
	db                     *sql.DB
	userRepository         *UserRepository
	taskRepository         *TaskRepository
	solutionRepository     *SolutionRepository
	datatestRepository     *DatatestRepository
	notificationRepository *NotificationRepository
	commentRepository      *CommentRepository
	tagRepository          *TagRepository
}

func New(db *sql.DB) Store {
	return Store{db: db}
}
