package database

import (
	"github.com/Masterminds/squirrel"
	"github.com/storyofhis/auth-service/repositories/model"
)

type PGAuth struct {
	DB *DBTransaction
}

// CreateUser create user on database
func (pg *PGAuth) CreateUser(in *model.User) (err error) {
	if err = pg.DB.Builder.
		Insert("public.t_users").
		Columns("name", "email", "passw", "level_id").
		Values(in.Name, in.Email, in.Passw, in.Permission.ID).
		Suffix(`RETURNING "id"`).Scan(new(string)); err != nil {
		return err
	}

	return nil
}

// GetUser get data of user on database
func (pg *PGAuth) GetUser(email *string) (res *model.User, err error) {
	res = new(model.User)

	if err = pg.DB.Builder.Select(`
		TU.id AS user_id,
		TU.name AS user_name,
		TU.email AS email_user,
		TU.pass AS user_passw,
		TUL.id AS user_level_id,
		TUL.name AS user_level_name,
	`).
		From("public.t_users TU").
		Join("public.t_users_level TUL ON TUL.id = TU.level_id").
		Where(squirrel.Eq{
			"email": email,
		}).Limit(1).Scan(
		&res.Id,
		&res.Name,
		&res.Email,
		&res.Passw,
		&res.Permission.ID,
		&res.Permission.Name,
	); err != nil {
		return res, err
	}

	// Check user Level
	if res.Permission.ID == model.Permission["admin"] {
		res.Permission.IsAdmin = true
	}

	return res, nil
}
