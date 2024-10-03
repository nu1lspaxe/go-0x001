package postgresql

import (
	"context"
	"database/sql"

	"github.com/nu1lspaxe/go-0x001/server/domain"

	"github.com/sirupsen/logrus"
)

type postgresqlDigimonRepository struct {
	db *sql.DB
}

func NewPostgresqlDigimonRepository(db *sql.DB) domain.DigimonRepository {
	return &postgresqlDigimonRepository{db: db}
}

// Implement DigimonRepository: GetById, Store, UpdateStatus

func (p *postgresqlDigimonRepository) GetById(ctx context.Context, id string) (*domain.Digimon, error) {
	row := p.db.QueryRow("SELECT id, name, status FROM digimons WHERE id =$1", id)
	d := &domain.Digimon{}
	if err := row.Scan(&d.Id, &d.Name, &d.Status); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return d, nil
}

func (p *postgresqlDigimonRepository) Store(ctx context.Context, d *domain.Digimon) error {
	_, err := p.db.Exec(
		"INSERT INTO digimons (id, name, status) VALUES ($1, $2, $3)",
		d.Id, d.Name, d.Status,
	)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (p *postgresqlDigimonRepository) UpdateStatus(ctx context.Context, d *domain.Digimon) error {
	_, err := p.db.Exec(
		"UPDATE digimons SET status=$1, WHERE id=$2",
		d.Status, d.Id,
	)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
