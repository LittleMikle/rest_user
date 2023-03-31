Запуск базы через докер контейнер:

docker-compose up --build

Если не сработает:

docker run --name postgres -e POSTGRES_PASSWORD=mysecretpassword -d -p 5432:5432 postgres




Запуск самого Rest:

go run cmd/main.go
