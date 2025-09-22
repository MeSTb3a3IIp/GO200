package store

import "apiserver/model"

type NotificationRepository struct {
	store         *Store
	notifications map[int]*model.Notification
}
