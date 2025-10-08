package store

import "apiserver/model"

type SolutionRepository struct {
	store     *Store
	solutions map[int]*model.Solution
}
