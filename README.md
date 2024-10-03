# go-0x001

## Introduction
This repository implements gRPC protocol and make two services communicate: server (digimon), server_2 (weather).

## Set up
1. Specify the `.env` file, you can modify the template `.env.tmp` according to your needs.
2. Copy `docker-compose-tmp.yml` file and rename it as `docker-compose.yml`, then set postgresql password.
3. Test it!
   ```bash
   # start services
   docker compose up -d
   
   # run the test
   cd server/test
   go run grpc.go
   
   # close services
   docker compose down -v
   ```

## API Protocol
This repository implements these two transfer protocols:
- REST (branch: [digimon](https://github.com/nu1lspaxe/go-0x001/tree/digimon))
- gRPC (branch: main)

## Tool
- protoc
  ```bash
  protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative <proto_file_name>.proto
  ```
- [mockery](https://github.com/vektra/mockery)
  - Generate mock objects based on interface
  - Run the command in `server/domain` and the result is generated in `server/domain/mocks`
  ```bash
  mockery --all --keeptree
  ```
- [swagger](https://github.com/swaggo/swag)
  - Run the command in `server` and the result is generated in `server/docs`
  ```bash
  swag init -g cmd/main.go
  # update swagger just run again
  ```
  - `server/docs` contains: `docs.go`, `swagger.json`, `swagger.yaml`
  - Import packages and set router to apply swagger api
  ```go
  import (
    swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "go_0x001/server/docs"    # {root_path}/docs
  )

  router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
  ```

## Test
```bash
# Coverage test
go test -coverprofile cover.out ./...
# Show in browser
go tool cover -html="cover.out" -o cover.html 
```
