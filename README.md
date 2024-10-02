# go-0x001

## Set up
Set up the setting in the following files :
1. `docker-compose-tmp.yml` to `docker-compose.yml`
2. `server/.env.tmp` to `server/.env`

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
