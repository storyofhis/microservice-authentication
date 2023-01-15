package services

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/storyofhis/auth-service/repositories"
	"github.com/storyofhis/auth-service/repositories/database"
	"github.com/storyofhis/auth-service/repositories/model"
	app "github.com/storyofhis/auth-service/repositories/proto"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(ctx context.Context, in *app.User) (err error) {
	var (
		pass        []byte
		transaction *database.DBTransaction
	)

	// Generate password
	if pass, err = bcrypt.GenerateFromPassword([]byte(in.Passw), 14); err != nil {
		return err
	}

	// give the params data
	var data = model.User{
		Name:  in.Name,
		Email: in.Email,
		Passw: string(pass),
		Permission: model.LevelUser{
			ID: in.Permission.Id,
		},
	}

	// Initializing transaction with database
	if transaction, err = database.OpenConnection(ctx, false); err != nil {
		return err
	}

	defer transaction.Rollback()

	var repo = repositories.NewRepositories(transaction)
	if err = repo.CreateUser(data); err != nil {
		return err
	}

	if err = transaction.Commit(); err != nil {
		return err
	}
	return nil
}

// Login
func Login(ctx context.Context, in *app.Credentials) (res *app.User, err error) {
	res = new(app.User)

	var (
		transaction *database.DBTransaction
	)
	transaction, err = database.OpenConnection(ctx, true)
	if err != nil {
		return nil, err
	}
	defer transaction.Rollback()

	var repo = repositories.NewRepositories(transaction)
	data, err := repo.GetUser(&in.Email)
	if err != nil {
		return nil, errors.New("User not found: " + err.Error())
	}

	if err = bcrypt.CompareHashAndPassword([]byte(data.Passw), []byte(in.Passw)); err != nil {
		return nil, err
	}

	if res.Token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, model.Session{
		Name: &data.Name,
		Permission: model.UserLevel{
			IsAdmin: &data.Permission.IsAdmin,
			ID:      &data.Permission.ID,
			Name:    &data.Permission.Name,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "storyofhis.auth",
			NotBefore: time.Now().Unix(),
		},
	}).SignedString([]byte(database.Get().SecretKey)); err != nil {
		return nil, errors.New("Could not generate token")
	}

	res.Id = data.Id
	res.Name = data.Name
	res.Email = data.Email
	res.Passw = ""
	return
}
