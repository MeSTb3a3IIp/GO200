package store

import (
	"apiserver/model"
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

func (s *Store) User() *UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
			users: make(map[int]*model.User),
		}
	}
	return s.userRepository
}

func (s *Store) Task() *TaskRepository {
	if s.taskRepository == nil {
		s.taskRepository = &TaskRepository{
			store: s,
			tasks: make(map[int]*model.Task),
		}
	}
	return s.taskRepository
}

func (s *Store) Solution() *SolutionRepository {
	if s.solutionRepository == nil {
		s.solutionRepository = &SolutionRepository{
			store:     s,
			solutions: make(map[int]*model.Solution),
		}
	}
	return s.solutionRepository
}

func (s *Store) Notification() *NotificationRepository {
	if s.notificationRepository == nil {
		s.notificationRepository = &NotificationRepository{
			store:         s,
			notifications: make(map[int]*model.Notification),
		}
	}
	return s.notificationRepository
}

func (s *Store) Comment() *CommentRepository {
	if s.taskRepository == nil {
		return &CommentRepository{
			store:    s,
			comments: make(map[int]*model.Comment),
		}
	}
	return s.commentRepository
}

func (s *Store) Datatest() *DatatestRepository {
	if s.datatestRepository != nil {
		return &DatatestRepository{
			store:     s,
			datatests: make(map[int]*model.Datatest),
		}
	}
	return s.datatestRepository
}

func (s *Store) Tag() *TagRepository {
	if s.tagRepository != nil {
		return &TagRepository{
			store: s,
			tags:  make(map[int]*model.Tag),
		}
	}
	return s.tagRepository
}
