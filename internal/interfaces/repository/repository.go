package repository

import (
	"time"

	"authentication/internal/entities"
	"github.com/dgrijalva/jwt-go/v4"
)

type database interface {
	AddProfile(credentials entities.Credentials) error
	GetProfileByEmail(email string) (entities.Profile, error)
}

type driver struct {
	d         database
	secretKey string
}

func New(dbHandler database, secretKey string) *driver {
	return &driver{
		d:         dbHandler,
		secretKey: secretKey,
	}
}

func (driver driver) InsertProfile(credentials entities.Credentials) error {
	return driver.d.AddProfile(credentials)
}

func (driver driver) QueryProfileByEmail(email string) (entities.Profile, error) {
	return driver.d.GetProfileByEmail(email)
}

func (driver driver) GenerateAuthenticationToken(profileId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &entities.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(60 * 24 * time.Hour)),
			IssuedAt:  jwt.At(time.Now()),
		},
		UserId: profileId,
	})

	return token.SignedString([]byte(driver.secretKey))
}
