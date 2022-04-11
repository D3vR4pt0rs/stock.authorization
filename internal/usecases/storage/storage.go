package storage

import (
	"time"

	"authentication/internal/entities"
	"github.com/D3vR4pt0rs/logger"
	"github.com/dgrijalva/jwt-go/v4"
	"golang.org/x/crypto/bcrypt"
)

type Controller interface {
	Registration(entities.Credentials) error
	SignUp(entities.Credentials) (string, error)
}

type Repository interface {
	InsertProfile(credentials entities.Credentials) error
	QueryProfileByEmail(email string) (entities.Profile, error)
}

type application struct {
	repo      Repository
	secretKey string
}

func New(repo Repository, secretKey string) *application {
	return &application{
		repo:      repo,
		secretKey: secretKey,
	}
}

func (app application) Registration(credentials entities.Credentials) error {
	_, err := app.repo.QueryProfileByEmail(credentials.Email)
	if err == nil {
		logger.Error.Println("Account exist")
		return AccountExistError
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error.Println(err.Error())
		return InternalError
	}

	credentials.Password = string(hash)
	err = app.repo.InsertProfile(credentials)

	if err != nil {
		logger.Error.Println(err.Error())
		return InternalError
	}
	return nil
}

func (app application) SignUp(credentials entities.Credentials) (string, error) {
	profile, err := app.repo.QueryProfileByEmail(credentials.Email)
	if err != nil {
		logger.Error.Println("Account didn't exist")
		return "", AccountNotFoundError
	}

	err = bcrypt.CompareHashAndPassword([]byte(profile.Password), []byte(credentials.Password))
	if err != nil {
		logger.Error.Println("Password is wrong")
		return "", WrongPasswordError
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &entities.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(60)),
			IssuedAt:  jwt.At(time.Now()),
		},
		UserId: profile.ID,
	})

	return token.SignedString([]byte(app.secretKey))
}
