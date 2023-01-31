### setup project

`go mod init github.com/subhadip1203/GORM/project1`  => it's like package.json file

`go mod tidy`

### install packages

for SQL ORM : `go get -u gorm.io/gorm`

for PostgreSQL : `go get -u gorm.io/driver/postgres`

env file read : `go get -u github.com/joho/godotenv`