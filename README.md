# WIP - go football api

### Build and dependencies with go.mod

- Created a main.go file, imported `gin-gonic` and a basic example from its github page
- Ran a `go mod init go-football-project`, which created a `go.mod` file
- I then ran `go build` and this looked at the imports in the main.go file and then automatically added them into 
my `go.mod` file (in the 'require' section)
- The `go build` command also created a `go.sum` file (which keeps track of the exact dependency versions I now have)
- `go build` also created the `go-football-project` binary which I can execute via a `./go-football-project`. This runs
the entry point in the application, the main function within the main package.

### Tutorial followed:

- Have built some basic APIs via this medium post: https://medium.com/@cgrant/developing-a-simple-crud-api-with-go-gin-and-gorm-df87d98e6ed1

### TODOs

- Extract db connection credentials somewhere else
- Extract things in the main.go file so that everything is not in one place - https://www.alexedwards.net/blog/organising-database-access
- Figure how to get the DB accessible from more than just one place
- Get date of birth working player model with api request
- Get migrations running by not using the db.Automigrate function