# Blog using golang and react

## Tech Stacks
- Golang 
- React 
- Postgres
- Docker
- Swagger for docs
- Golang migrate for migrations

## Reference

[The 12 Factor App](https://12factor.net/)
[Roy Fielding REST dissertation](https://ics.uci.edu/~fielding/pubs/dissertation/fielding_dissertation.pdf)
[Richardson Maturity Model](https://martinfowler.com/articles/richardsonMaturityModel.html)


## Command

```
go run *.go    
curl http://localhost:8080/v1/health
export PATH=$PATH:$(go env GOPATH)/bin    
air init 
docker exec -it postgres_db psql -U admin -d postgres
migrate create -seq -ext sql -dir ./cmd/migrate/migrations create_users
migrate create -seq -ext sql -dir ./cmd/migrate/migrations create_posts

migrate -path=./cmd/migrate/migrations -database="postgres://admin:adminpassword@localhost/social?sslmode=disable" up
migrate -path=./cmd/migrate/migrations -database="postgres://admin:adminpassword@localhost/social?sslmode=disable" down

migrate -path=./cmd/migrate/migrations -database="postgres://admin:adminpassword@localhost/social?sslmode=disable" force 4

http://localhost:8080/v1/swagger/index.html
```