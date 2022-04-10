package repository

import "authentication/internal/entities"

type driver interface {
	AddProfile(profile entities.Profile) error
	AddRefreshSession(profile entities.RefreshSession) error
	GetProfile(key, value string) (entities.Profile, error)
	GetRefreshSession(key, value string) (entities.RefreshSession, error)
}

type database struct {
	d driver
}

func New(dbHandler driver) *database {
	return &database{
		d: dbHandler,
	}
}
