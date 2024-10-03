package service_test

import (
	"context"
	"errors"

	"github.com/nu1lspaxe/go-0x001/server/domain"
	"github.com/nu1lspaxe/go-0x001/server/domain/mocks"

	"testing"

	_dietServ "github.com/nu1lspaxe/go-0x001/server/diet/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetById(t *testing.T) {
	mockDietRepo := new(mocks.DietRepository)
	mockDiet := domain.Diet{
		Id:     "e9addf2d-8739-427a-8b30-2383b9b045b1",
		UserId: "ab18b1ba-48e1-48cf-88b5-48782874aa05",
		Name:   "Giuseppe",
	}

	t.Run("Success", func(t *testing.T) {
		mockDietRepo.
			On("GetById", mock.Anything, mock.MatchedBy(func(value string) bool { return value == "e9addf2d-8739-427a-8b30-2383b9b045b1" })).
			Return(&mockDiet, nil).Once()

		u := _dietServ.NewDietService(mockDietRepo)
		digimon, err := u.GetById(context.Background(), "e9addf2d-8739-427a-8b30-2383b9b045b1")

		assert.NoError(t, err)
		assert.NotNil(t, digimon)

		mockDietRepo.AssertExpectations(t)
	})

	t.Run("Fail", func(t *testing.T) {
		mockDietRepo.
			On("GetById", mock.Anything, mock.MatchedBy(func(value string) bool { return value == "e9addf2d-8739-427a-8b30-2383b9b045b1" })).
			Return(nil, errors.New("Get error")).Once()

		u := _dietServ.NewDietService(mockDietRepo)
		aDigimon, err := u.GetById(context.Background(), "e9addf2d-8739-427a-8b30-2383b9b045b1")

		assert.Error(t, err)
		assert.Nil(t, aDigimon)

		mockDietRepo.AssertExpectations(t)
	})
}

func TestStore(t *testing.T) {
	mockDietRepo := new(mocks.DietRepository)
	mockDiet := domain.Diet{
		Id:     "e9addf2d-8739-427a-8b30-2383b9b045b1",
		UserId: "ab18b1ba-48e1-48cf-88b5-48782874aa05",
		Name:   "Giuseppe",
	}

	t.Run("Success", func(t *testing.T) {
		mockDietRepo.
			On("Store", mock.Anything, mock.MatchedBy(func(d *domain.Diet) bool { return d == &mockDiet })).
			Return(nil).Once()

		u := _dietServ.NewDietService(mockDietRepo)
		err := u.Store(context.Background(), &mockDiet)

		assert.NoError(t, err)

		mockDietRepo.AssertExpectations(t)
	})
	t.Run("Fail", func(t *testing.T) {
		mockDietRepo.
			On("Store", mock.Anything, mock.MatchedBy(func(d *domain.Diet) bool { return d == &mockDiet })).
			Return(errors.New("Get error")).Once()

		u := _dietServ.NewDietService(mockDietRepo)
		err := u.Store(context.Background(), &mockDiet)

		assert.Error(t, err)

		mockDietRepo.AssertExpectations(t)
	})
}
