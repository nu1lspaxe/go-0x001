package http

import (
	"go_0x001/server/domain"
	"go_0x001/server/swagger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "go_0x001/server/docs"
)

type DigimonHandler struct {
	DigimonService domain.DigimonService
	DietService    domain.DietService
}

func NewDigimonHandler(
	e *gin.Engine,
	digimonService domain.DigimonService,
	dietService domain.DietService,
) {
	handler := &DigimonHandler{
		DigimonService: digimonService,
		DietService:    dietService,
	}

	e.GET("/api/v1/digimons/:digimonId", handler.GetDigimonByDigimonId)
	e.POST("/api/v1/digimons", handler.PostDigimon)
	e.POST("/api/v1/digimons/:digimonId/foster", handler.PostFosterDigimon)

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// GetDigimonByDigimonId godoc
//
//	@Summary		Get Digimon Information
//	@Description	get digimon by Id
//	@Tags
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Digimon Id"
//	@Success	200	{object}	swagger.DigimonInfo
//	@Failure	500	{object}	swagger.ModelError
//	@Router		/api/v1/digimons/{digimonId} [get]
func (d *DigimonHandler) GetDigimonByDigimonId(c *gin.Context) {
	digimonId := c.Param("digimonId")

	digimon, err := d.DigimonService.GetById(c, digimonId)
	if err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal error. Query digimon error",
		})
		return
	}

	c.JSON(200, &swagger.DigimonInfo{
		Id:     digimon.Id,
		Name:   digimon.Name,
		Status: digimon.Status,
	})
}

func (d *DigimonHandler) PostDigimon(c *gin.Context) {
	var body swagger.DigimonInfoRequest
	if err := c.BindJSON(&body); err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal error. Parsing failed",
		})
		return
	}

	digimon := domain.Digimon{
		Name: body.Name,
	}

	if err := d.DigimonService.Store(c, &digimon); err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal error. Store failed",
		})
		return
	}

	c.JSON(200, swagger.DigimonInfo{
		Id:     digimon.Id,
		Name:   digimon.Name,
		Status: digimon.Status,
	})
}

func (d *DigimonHandler) PostFosterDigimon(c *gin.Context) {
	digimonId := c.Param("digimonId")

	var body swagger.FosterRequest
	if err := c.BindJSON(&body); err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal error. Parsing failed",
		})
		return
	}

	if err := d.DietService.Store(c, &domain.Diet{
		UserId: digimonId,
		Name:   body.Food.Name,
	}); err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal error. Store failed",
		})
		return
	}

	if err := d.DigimonService.UpdateStatus(c, &domain.Digimon{
		Id:     digimonId,
		Status: "good",
	}); err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal error. Store failed",
		})
		return
	}
	c.JSON(204, nil)
}
