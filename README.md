# WIP - go football api

### Build and dependencies with go.mod

- Created a main.go file, imported `gin-gonic` and a basic example from its github page
- Ran a `go mod init go-football-project`, which created a `go.mod` file
- I then ran `go build` and this looked at the imports in the main.go file and then automatically added them into 
my `go.mod` file (in the 'require' section)
- The `go build` command also created a `go.sum` file (which keeps track of the exact dependency versions I now have)
- `go build` also created the `go-football-project` binary which I can execute via a `./go-football-project`. This runs
the entry point in the application, the main function within the main package.

