# Commands

if necessary:
source ~/.profile

```
go mod init github.com/riquefsouza/hefesto-fiber-gorm
go get -u github.com/gofiber/fiber/v2

go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite


go run main.go
```

```
curl http://localhost:3000/api/v1/admParameterCategory/1

curl -X POST http://localhost:3000/api/v1/admParameterCategory

curl -X POST -H "Content-Type: application/json" --data "{\"Description\": \"Nova Categoria\"}" http://localhost:3000/api/v1/admParameterCategory
curl http://localhost:3000/api/v1/admParameterCategory/2

curl -X DELETE http://localhost:3000/api/v1/admParameterCategory/1
```
