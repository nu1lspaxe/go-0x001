package main

import (
	"database/sql"
	"fmt"
	"net"

	_dietRepo "github.com/nu1lspaxe/go-0x001/server/diet/repository/postgresql"
	_dietServ "github.com/nu1lspaxe/go-0x001/server/diet/service"
	_digimonHandlerGrpcDelivery "github.com/nu1lspaxe/go-0x001/server/digimon/delivery/grpc"
	_digimonRepo "github.com/nu1lspaxe/go-0x001/server/digimon/repository/postgresql"
	_digimonServ "github.com/nu1lspaxe/go-0x001/server/digimon/service"
	_weatherRepo "github.com/nu1lspaxe/go-0x001/server/weather/repository/grpc"
	_weatherService "github.com/nu1lspaxe/go-0x001/server/weather/service"

	pb "github.com/nu1lspaxe/go-0x001/server_2/proto/weather"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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

//	@host		localhost:6000
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	logrus.Info("GRPC Server started")

	// RESTful
	// restfulHost := viper.GetString("RESTFUL_HOST")
	// restfulPort := viper.GetString("RESTFUL_PORT")

	grpcPort := viper.GetString("GRPC_PORT")
	grpcWeatherAddress := viper.GetString("GRPC_WEATHER_ADDRESS")

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

	// r := gin.Default()
	// _digimonHandlerHttpDelivery.NewDigimonHandler(r, digimonServ, dietServ)
	// logrus.Fatal(r.Run(restfulHost + ":" + restfulPort))

	conn, err := grpc.NewClient(grpcWeatherAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	weatherClient := pb.NewWeatherClient(conn)

	weatherRepo := _weatherRepo.NewgrpcWeatherRepository(weatherClient)
	digimonRepo := _digimonRepo.NewPostgresqlDigimonRepository(db)
	dietRepo := _dietRepo.NewPostgresqlDietRepository(db)

	weatherServ := _weatherService.NewWeatherService(weatherRepo)
	digimonServ := _digimonServ.NewDigimonService(digimonRepo)
	dietServ := _dietServ.NewDietService(dietRepo)

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		logrus.Fatal(err)
	}
	s := grpc.NewServer()

	_digimonHandlerGrpcDelivery.NewDigimonHandler(s, digimonServ, dietServ, weatherServ)

	logrus.Fatal(s.Serve(lis))
}
