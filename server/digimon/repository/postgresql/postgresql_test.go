package postgresql_test

import (
	"context"
	"fmt"
	"reflect"
	digimonPostgresqlRepo "server/digimon/repository/postgresql"
	"server/domain"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gotest.tools/assert"
)

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockDigimon := &domain.Digimon{
		Id:     "69770f2d-933e-474d-8357-a2f8a9c874df",
		Name:   "Customer",
		Status: "Good",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "status"}).
		AddRow(mockDigimon.Id, mockDigimon.Name, mockDigimon.Status)

	query := "SELECT id, name, status FROM digimons WHERE id =?"

	mock.ExpectQuery(query).WithArgs("69770f2d-933e-474d-8357-a2f8a9c874df").WillReturnRows(rows)
	d := digimonPostgresqlRepo.NewPostgresqlDigimonRepository(db)
	diet, _ := d.GetById(context.Background(), "69770f2d-933e-474d-8357-a2f8a9c874df")
	assert.Equal(t, mockDigimon.Name, diet.Name)
	fmt.Println("Deep Equal?", reflect.DeepEqual(mockDigimon, diet))
}

func TestStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockDigimon := &domain.Digimon{
		Id:     "69770f2d-933e-474d-8357-a2f8a9c874df",
		Name:   "Customer",
		Status: "Good",
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO digimons").
		WithArgs(mockDigimon.Id, mockDigimon.Name, mockDigimon.Status).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	d := digimonPostgresqlRepo.NewPostgresqlDigimonRepository(db)

	tx, _ := db.Begin()
	d.Store(context.Background(), mockDigimon)
	tx.Commit()

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockDigimon := &domain.Digimon{
		Id:     "69770f2d-933e-474d-8357-a2f8a9c874df",
		Name:   "Customer",
		Status: "Good",
	}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE digimons").
		WithArgs(mockDigimon.Status, mockDigimon.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	d := digimonPostgresqlRepo.NewPostgresqlDigimonRepository(db)

	tx, _ := db.Begin()
	d.UpdateStatus(context.Background(), mockDigimon)
	tx.Commit()

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
