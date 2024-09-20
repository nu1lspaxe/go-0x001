package main

import (
	"database/sql"
	"fmt"

	_dietRepo "go_0x001/server/diet/repository/postgresql"
	_dietServ "go_0x001/server/diet/service"
	_digimonHandlerHttpDelivery "go_0x001/server/digimon/delivery/http"
	_digimonRepo "go_0x001/server/digimon/repository/postgresql"
	_digimonServ "go_0x001/server/digimon/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func init() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("dotenv")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalln("Fatal error config file:", err)
	}
}

//	@title			Swagger API
//	@version		1.0
//	@description	Digimon server repo.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:5000
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	logrus.Info("HTTP server started")

	restfulHost := viper.GetString("RESTFUL_HOST")
	restfulPort := viper.GetString("RESTFUL_PORT")
	dbHost := viper.GetString("DB_HOST")
	dbDatabase := viper.GetString("DB_DATABASE")
	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASSWORD")

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbDatabase),
	)
	if err != nil {
		logrus.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		logrus.Fatal(err)
	}

	r := gin.Default()

	digimonRepo := _digimonRepo.NewPostgresqlDigimonRepository(db)
	dietRepo := _dietRepo.NewPostgresqlDietRepository(db)

	digimonServ := _digimonServ.NewDigimonService(digimonRepo)
	dietServ := _dietServ.NewDietUsecase(dietRepo)

	_digimonHandlerHttpDelivery.NewDigimonHandler(r, digimonServ, dietServ)

	logrus.Fatal(r.Run(restfulHost + ":" + restfulPort))
}
