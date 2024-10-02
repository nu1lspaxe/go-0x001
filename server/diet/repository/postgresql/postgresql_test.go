package postgresql_test

import (
	"fmt"
	"reflect"
	dietPostgresqlRepo "server/diet/repository/postgresql"
	"server/domain"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"golang.org/x/net/context"
	"gotest.tools/assert"
)

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockDiet := &domain.Diet{
		Id:     "ad1de98f-d5ec-4976-867e-531429a28cda",
		UserId: "318f5799-e2f9-4ace-96e0-6658fe603d10",
		Name:   "Apple",
	}

	rows := sqlmock.NewRows([]string{"id", "user_id", "name"}).
		AddRow(mockDiet.Id, mockDiet.UserId, mockDiet.Name)

	query := "SELECT id FROM diets WHERE id = ?"

	mock.ExpectQuery(query).WithArgs("ad1de98f-d5ec-4976-867e-531429a28cda").WillReturnRows(rows)

	d := dietPostgresqlRepo.NewPostgresqlDietRepository(db)
	diet, _ := d.GetById(context.Background(), "ad1de98f-d5ec-4976-867e-531429a28cda")
	assert.Equal(t, mockDiet.UserId, diet.UserId)
	fmt.Println("Deep Equal?", reflect.DeepEqual(mockDiet, diet))
}

func TestStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockDiet := &domain.Diet{
		Id:     "ad1de98f-c0cd-46c6-9349-eec80f4b1a12",
		UserId: "318f5799-e2f9-4ace-96e0-6658fe603d10",
		Name:   "Apple",
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO diets").
		WithArgs(mockDiet.Id, mockDiet.UserId, mockDiet.Name).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	d := dietPostgresqlRepo.NewPostgresqlDietRepository(db)
	tx, _ := db.Begin()
	d.Store(context.Background(), mockDiet)
	tx.Commit()

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
