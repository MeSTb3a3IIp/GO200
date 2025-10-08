package store

import "apiserver/model"

type TaskRepository struct {
	store *Store
	tasks map[int]*model.Task
}

func (taskRepository *TaskRepository) FindAll() ([]model.Task, error) {
	rows, err := taskRepository.store.db.Query("SELECT id, username, email, password, count_easy_task, count_medium_task, count_hard_task")
	if err != nil {
		return nil, err
	}
	var tasks []model.Task

	for rows.Next() {
		task := model.Task{}
		if err := rows.Scan(&task.Id, &task.Name, &task.Description, &task.Date, &task.Difficulty); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil

}

func (taskRepository *TaskRepository) Find(id int) (*model.Task, error) {
	task := &model.Task{}
	if err := taskRepository.store.db.QueryRow("SELECT id, name, discription, date, diffivulty FROM tasks WHERE id=$1", id).
		Scan(&task.Id, &task.Name, &task.Description, &task.Date, &task.Difficulty); err != nil {
		return nil, err
	}
	return task, nil
}
