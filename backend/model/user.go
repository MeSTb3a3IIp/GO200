package model

type User struct {
	Id                int
	Username          string
	Email             string
	Password          string
	Count_easy_task   int
	Count_medium_task int
	Count_hard_task   int
}
