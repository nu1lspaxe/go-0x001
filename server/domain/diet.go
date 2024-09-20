package domain

import "context"

type Diet struct {
	Id     string
	UserId string
	Name   string
}

type DietRepository interface {
	GetById(ctx context.Context, id string) (*Diet, error)
	Store(ctx context.Context, d *Diet) error
}

type DietService interface {
	GetById(ctx context.Context, id string) (*Diet, error)
	Store(ctx context.Context, d *Diet) error
}
