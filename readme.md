# Alura Store Go project
## Set up
### Setup the database running the docker compose
```bash
docker compose up -d
```
### Migrations
In this project the migrations are done with Gorm AutoMigrate
```go
err = DB.AutoMigrate(&models.Student{})
```

## Tear down
### Stop the database
To stop the database container run the docker compose down command
```bash
docker compose down
```


