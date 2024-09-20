package domain

import "context"

type Digimon struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type DigimonRepository interface {
	GetById(ctx context.Context, id string) (*Digimon, error)
	Store(ctx context.Context, d *Digimon) error
	UpdateStatus(ctx context.Context, d *Digimon) error
}

type DigimonService interface {
	GetById(ctx context.Context, id string) (*Digimon, error)
	Store(ctx context.Context, d *Digimon) error
	UpdateStatus(ctx context.Context, d *Digimon) error
}
