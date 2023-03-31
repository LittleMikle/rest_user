Запуск базы через докер контейнер:

docker-compose up --build

Если не сработает:

docker run --name postgres -e POSTGRES_PASSWORD=mysecretpassword -d -p 5432:5432 postgres




Запуск самого Rest:

go run cmd/main.go


![Screenshot from 2023-03-31 04-57-57](https://user-images.githubusercontent.com/101155101/229003703-1cbc6355-5092-480a-9c67-e9fb44234950.png)


![Screenshot from 2023-03-31 04-59-22](https://user-images.githubusercontent.com/101155101/229003714-cb982d97-fa5c-4c91-a7fe-55099a864d88.png)
