package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nu1lspaxe/go-0x001/server/domain"
	"github.com/nu1lspaxe/go-0x001/server/domain/mocks"

	digimonHandlerHttpDelivery "github.com/nu1lspaxe/go-0x001/server/digimon/delivery/http"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"gotest.tools/assert"
)

func TestGetDigimonByDigimonId(t *testing.T) {
	mockDigimon := domain.Digimon{
		Id:     "8c862535-6de2-4da2-ad21-c853e4343bd7",
		Name:   "III",
		Status: "good",
	}
	mockDigimonMarshal, _ := json.Marshal(mockDigimon)
	mockDigimonServ := new(mocks.DigimonService)

	mockDigimonServ.On("GetById", mock.Anything, mockDigimon.Id).Return(&mockDigimon, nil)

	r := gin.Default()

	handler := digimonHandlerHttpDelivery.DigimonHandler{
		DigimonService: mockDigimonServ,
	}

	r.GET("/api/v1/digimons/:digimonId", handler.GetDigimonByDigimonId)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/digimons/8c862535-6de2-4da2-ad21-c853e4343bd7", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(mockDigimonMarshal), w.Body.String())
}

func TestPostToCreateDigimon(t *testing.T) {
	mockDigimonServ := new(mocks.DigimonService)

	mockDigimonServ.
		On("Store", mock.Anything, mock.AnythingOfType("*domain.Digimon")).
		Return(nil).
		Run(func(args mock.Arguments) {
			arg := args.Get(1).(*domain.Digimon)
			arg.Id = "e5d88876-e513-43e9-80ac-6b348d84d8b4"
			arg.Status = "good"
		})

	r := gin.Default()

	handler := digimonHandlerHttpDelivery.DigimonHandler{
		DigimonService: mockDigimonServ,
	}

	r.POST("/api/v1/digimons", handler.PostDigimon)

	jsonString := []byte(`{"name":"Agumon"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/digimons", bytes.NewBuffer(jsonString))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"id":"e5d88876-e513-43e9-80ac-6b348d84d8b4","name":"Agumon","status":"good"}`, w.Body.String())
}

func TestPostToFosterDigimon(t *testing.T) {
	mockDigimonServ := new(mocks.DigimonService)
	mockDietServ := new(mocks.DietService)

	mockDietServ.
		On("Store", mock.Anything, mock.MatchedBy(func(value *domain.Diet) bool {
			return value.UserId == "178744e4-a218-45ac-adfa-9023a3bf9699" && value.Name == "apple"
		})).
		Return(nil)

	mockDigimonServ.
		On("UpdateStatus", mock.Anything, mock.Anything).
		Return(nil)

	r := gin.Default()

	handler := digimonHandlerHttpDelivery.DigimonHandler{
		DigimonService: mockDigimonServ,
		DietService:    mockDietServ,
	}

	r.POST("/api/v1/digimons/:digimonId/foster", handler.PostFosterDigimon)

	jsonString := []byte(`{"food": {"name": "apple"}}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/digimons/178744e4-a218-45ac-adfa-9023a3bf9699/foster", bytes.NewBuffer(jsonString))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)
}
