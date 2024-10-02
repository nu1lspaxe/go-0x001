# go-0x001

## Set up
1. Specify the `.env` file, you can modify the template `.env.tmp` according to your needs.
2. Set postgresql password in `docker-compose-tmp.yml` file.
3. The main file is stored in `cmd/main.go`, run it in any way you like.

## API Protocol
This repository implements these two transfer protocols:
- REST
- gRPC

## Tool
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
