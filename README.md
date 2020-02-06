# Usage

Deploy docker postgresql container

```
docker pull postgres:11.6
docker run --name postgres -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=test -e POSTGRES_DB=test -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data  postgres:11.6
```
Go into the postgresql container
```
docker exec -it postgres bash
```
Create a table
```
create table movie(
    id serial primary key,
    rating varchar (50),
    name varchar (50)
);
```

Install packages
```
go get github.com/gin-gonic/gin
go get github.com/lib/pq
```

Run application
```
go run main.go
```

Test API with curl

List all movies
```
curl localhost:8080/movie
```
Add movie
```
curl -X POST localhost:8080/movie -d '{"rating":"Excellent", "name":"Once upon a time in Hollywood"}'
```
Update movie
```
curl -X PUT localhost:8080/movie/1 -d '{"rating":"Bad", "name":"Once upon a time in Hollywood"}'
```
Delete movie
```
curl -X DELETE localhost:8080/movie/1
```
