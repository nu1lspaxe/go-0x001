package service_test

import (
	"context"
	"errors"
	_digimonServ "server/digimon/service"
	"server/domain"
	"server/domain/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetById(t *testing.T) {
	mockDigimonRepo := new(mocks.DigimonRepository)
	mockDigimon := domain.Digimon{
		Id:     "2e72c27e-0feb-44e1-89ef-3d58fd30a1b3",
		Name:   "Metrics",
		Status: "Good",
	}

	t.Run("Success", func(t *testing.T) {
		mockDigimonRepo.
			On("GetById", mock.Anything, mock.MatchedBy(func(value string) bool { return value == "e9addf2d-8739-427a-8b30-2383b9b045b1" })).
			Return(&mockDigimon, nil).Once()

		u := _digimonServ.NewDigimonService(mockDigimonRepo)
		aDigimon, err := u.GetById(context.Background(), "e9addf2d-8739-427a-8b30-2383b9b045b1")

		assert.NoError(t, err)
		assert.NotNil(t, aDigimon)

		mockDigimonRepo.AssertExpectations(t)
	})
	t.Run("Fail", func(t *testing.T) {
		mockDigimonRepo.
			On("GetById", mock.Anything, mock.MatchedBy(func(value string) bool { return value == "e9addf2d-8739-427a-8b30-2383b9b045b1" })).
			Return(nil, errors.New("Get error")).Once()

		u := _digimonServ.NewDigimonService(mockDigimonRepo)
		aDigimon, err := u.GetById(context.Background(), "e9addf2d-8739-427a-8b30-2383b9b045b1")

		assert.Error(t, err)
		assert.Nil(t, aDigimon)

		mockDigimonRepo.AssertExpectations(t)
	})
}

func TestStore(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockDigimonRepo := new(mocks.DigimonRepository)
		mockDigimon := domain.Digimon{
			Id:     "2e72c27e-0feb-44e1-89ef-3d58fd30a1b3",
			Name:   "Metrics",
			Status: "Good",
		}

		mockDigimonRepo.
			On("Store", mock.Anything, mock.MatchedBy(func(d *domain.Digimon) bool { return d == &mockDigimon })).
			Return(nil).Once()

		u := _digimonServ.NewDigimonService(mockDigimonRepo)
		err := u.Store(context.Background(), &mockDigimon)

		assert.NoError(t, err)

		mockDigimonRepo.AssertExpectations(t)
	})
	t.Run("Success. When 'Status' is blank, will set 'good'. When 'ID' is blank, will set random", func(t *testing.T) {
		mockDigimonRepo := new(mocks.DigimonRepository)
		mockDigimon := domain.Digimon{
			Name: "Metrics",
		}

		mockDigimonRepo.
			On("Store", mock.Anything, mock.MatchedBy(func(d *domain.Digimon) bool {
				return d.Id != "" && d.Status == "good" && d.Name == "Metrics"
			})).
			Return(nil).Once()

		u := _digimonServ.NewDigimonService(mockDigimonRepo)
		err := u.Store(context.Background(), &mockDigimon)

		assert.NoError(t, err)

		mockDigimonRepo.AssertExpectations(t)
	})
	t.Run("Fail", func(t *testing.T) {
		mockDigimonRepo := new(mocks.DigimonRepository)
		mockDigimon := domain.Digimon{
			Id:     "2e72c27e-0feb-44e1-89ef-3d58fd30a1b3",
			Name:   "Metrics",
			Status: "Good",
		}

		mockDigimonRepo.
			On("Store", mock.Anything, mock.MatchedBy(func(d *domain.Digimon) bool { return d == &mockDigimon })).
			Return(errors.New("Get error")).Once()

		u := _digimonServ.NewDigimonService(mockDigimonRepo)
		err := u.Store(context.Background(), &mockDigimon)

		assert.Error(t, err)

		mockDigimonRepo.AssertExpectations(t)
	})
}

func TestUpdateStatus(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockDigimonRepo := new(mocks.DigimonRepository)
		mockDigimon := domain.Digimon{
			Id:     "2e72c27e-0feb-44e1-89ef-3d58fd30a1b3",
			Status: "Good",
		}

		mockDigimonRepo.
			On("UpdateStatus", mock.Anything, mock.MatchedBy(func(d *domain.Digimon) bool {
				return d.Status == "Good"
			})).
			Return(nil).Once()

		u := _digimonServ.NewDigimonService(mockDigimonRepo)
		err := u.UpdateStatus(context.Background(), &mockDigimon)

		assert.NoError(t, err)

		mockDigimonRepo.AssertExpectations(t)
	})
	t.Run("Fail. When 'Status' is blank", func(t *testing.T) {
		mockDigimonRepo := new(mocks.DigimonRepository)
		mockDigimon := domain.Digimon{
			Id: "2e72c27e-0feb-44e1-89ef-3d58fd30a1b3",
		}

		u := _digimonServ.NewDigimonService(mockDigimonRepo)
		err := u.UpdateStatus(context.Background(), &mockDigimon)

		assert.Error(t, err)
		assert.Equal(t, "status is blank", err.Error())
	})
	t.Run("Fail", func(t *testing.T) {
		mockDigimonRepo := new(mocks.DigimonRepository)
		mockDigimon := domain.Digimon{
			Id:     "2e72c27e-0feb-44e1-89ef-3d58fd30a1b3",
			Name:   "Metrics",
			Status: "Good",
		}

		mockDigimonRepo.
			On("UpdateStatus", mock.Anything, mock.MatchedBy(func(d *domain.Digimon) bool { return d == &mockDigimon })).
			Return(errors.New("Get error")).Once()

		u := _digimonServ.NewDigimonService(mockDigimonRepo)
		err := u.UpdateStatus(context.Background(), &mockDigimon)

		assert.Error(t, err)

		mockDigimonRepo.AssertExpectations(t)
	})
}
