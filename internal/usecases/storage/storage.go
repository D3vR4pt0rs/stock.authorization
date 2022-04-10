package storage

import "authentication/internal/entities"

type Controller interface {
	Registration(profile entities.Profile) error
	SignUp(profile entities.Profile) (entities.TokenInformation, error)
	RefreshToken(tokens entities.TokenInformation) (entities.TokenInformation, error)
}

type Repository interface {
	InsertProfile(profile entities.Profile) error
	QueryProfile(id string) (entities.Profile, error)
	InsertRefreshToken(session entities.RefreshSession) error
}

type application struct {
	repo Repository
}

func New(repo Repository) *application {
	return &application{
		repo: repo,
	}
}
