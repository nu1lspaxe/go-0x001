package postgresql

import (
	"context"
	"database/sql"
	"go_0x001/server/domain"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

type postgresqlDietRepository struct {
	db *sql.DB
}

func NewPostgresqlDietRepository(db *sql.DB) domain.DietRepository {
	return &postgresqlDietRepository{db}
}
func (p *postgresqlDietRepository) GetById(ctx context.Context, id string) (*domain.Diet, error) {
	row := p.db.QueryRow("SELECT id FROM diets WHERE id = $1", id)
	d := &domain.Diet{}
	if err := row.Scan(&d.Id, &d.UserId, &d.Name); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return d, nil
}

func (p *postgresqlDietRepository) Store(ctx context.Context, d *domain.Diet) error {
	if d.Id == "" {
		d.Id = uuid.Must(uuid.NewV4()).String()
	}
	_, err := p.db.Exec(
		"INSERT INTO diets (id, user_id, name) VALUES ($1, $2, $3)",
		d.Id, d.UserId, d.Name,
	)
	if err != nil {
		logrus.Error(err, d)
		return err
	}
	return nil
}
