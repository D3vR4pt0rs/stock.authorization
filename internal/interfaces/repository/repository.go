package repository

import "authentication/internal/entities"

type driver interface {
	AddProfile(credentials entities.Credentials) error
	GetProfileByEmail(email string) (entities.Profile, error)
}

type database struct {
	d driver
}

func New(dbHandler driver) *database {
	return &database{
		d: dbHandler,
	}
}

func (driver database) InsertProfile(credentials entities.Credentials) error {
	return driver.d.AddProfile(credentials)
}

func (driver database) QueryProfileByEmail(email string) (entities.Profile, error) {
	return driver.d.GetProfileByEmail(email)
}
