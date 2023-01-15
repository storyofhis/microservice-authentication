package repositories

import (
	"github.com/storyofhis/auth-service/repositories/database"
	"github.com/storyofhis/auth-service/repositories/model"
)

type Repositories struct {
	pg *database.PGAuth
}

func NewRepositories(db *database.DBTransaction) *Repositories {
	return &Repositories{
		pg: &database.PGAuth{
			DB: db,
		},
	}
}

func (repo *Repositories) CreateUser(in model.User) error {
	return repo.pg.CreateUser(&in)
}

func (repo *Repositories) GetUser(email *string) (*model.User, error) {
	return repo.pg.GetUser(email)
}
