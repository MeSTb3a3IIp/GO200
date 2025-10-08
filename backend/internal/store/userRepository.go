package store

import (
	"apiserver/model"
	"database/sql"
	"errors"
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
}

func (userRepository *UserRepository) Create(user *model.User) error {
	//проверить на валидность
	if userTemp, err := userRepository.Find(user.Username); err != nil {
		_ = userTemp
		return err
	}
	return userRepository.store.db.QueryRow("INSERT INTO users VALUES ($1, $2, $3, $4, $5, $,6) RETURNING id", user.Username, user.Email, user.Password, user.Count_easy_task, user.Count_medium_task, user.Count_hard_task).Scan(&user.Id)

}

func (userRepository *UserRepository) FindByEmail(email string) (*model.User, error) {

	user := model.User{}

	if err := userRepository.store.db.QueryRow("SELECT username, email, password, count_easy_task, count_medium_task, count_hard_task FROM users WHERE email=$1", email).
		Scan(&user.Username, &user.Email, &user.Password, &user.Count_easy_task, &user.Count_medium_task, &user.Count_hard_task); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("record not found")
		}
		return nil, err
	}
	return &user, nil
}

// пока что нужно только для проверки уникальности юзернейма при создании или редактировании, код грязный
func (userRepository *UserRepository) Find(username string) (*model.User, error) {

	user := model.User{}

	if err := userRepository.store.db.QueryRow("SELECT username, email, password, count_easy_task, count_medium_task, count_hard_task FROM users WHERE username=$1", username).
		Scan(&user.Username, &user.Email, &user.Password, &user.Count_easy_task, &user.Count_medium_task, &user.Count_hard_task); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("record not found")
		}
		return nil, err
	}
	return &user, nil
}
