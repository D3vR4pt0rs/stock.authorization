package storage

import (
	"authentication/internal/entities"
	"github.com/D3vR4pt0rs/logger"
	"golang.org/x/crypto/bcrypt"
)

type Controller interface {
	Registration(entities.Credentials) error
	SignUp(entities.Credentials) (string, error)
}

type Repository interface {
	InsertProfile(credentials entities.Credentials) error
	QueryProfileByEmail(email string) (entities.Profile, error)
	GenerateAuthenticationToken(profileID int) (string, error)
}

type application struct {
	repo Repository
}

func New(repo Repository) *application {
	return &application{
		repo: repo,
	}
}

func (app application) Registration(credentials entities.Credentials) error {
	_, err := app.repo.QueryProfileByEmail(credentials.Email)
	if err == nil {
		return AccountExistError
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost)
	if err != nil {
		return InternalError
	}

	credentials.Password = string(hash)
	err = app.repo.InsertProfile(credentials)

	if err != nil {
		return InternalError
	}
	return nil
}

func (app application) SignUp(credentials entities.Credentials) (string, error) {
	profile, err := app.repo.QueryProfileByEmail(credentials.Email)
	if err != nil {
		return "", AccountNotFoundError
	}

	err = bcrypt.CompareHashAndPassword([]byte(profile.Password), []byte(credentials.Password))
	if err != nil {
		logger.Error.Println("Password is wrong")
		return "", WrongPasswordError
	}

	token, err := app.repo.GenerateAuthenticationToken(profile.ID)
	if err != nil {
		return "", InternalError
	}
	return token, nil
}
