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

###REST:
- ```curl -i -X GET http://localhost:8080/players```
- ```curl -i -X GET http://localhost:8080/players/1```
- ```curl -i -X DELETE http://localhost:8080/players/32```
- ```curl -i -X POST http://localhost:8080/players -d '{"first_name": "harry", "last_name": "kane", "age": 25}'```
- ```curl -i -X PUT http://localhost:8080/players -d '{"first_name": "harry", "last_name": "kane", "age": 25}'```

# TODO:

- Extract db connection credentials somewhere else
- Extract things in the main.go file so that everything is not in one place - https://www.alexedwards.net/blog/organising-database-access
- Figure how to get the DB accessible from more than just one place
- Get date of birth working player model with api request
- Get migrations running by not using the db.Automigrate function
- Implement the following to start with:

```
team
	- id
	- name
	- stadium (foreign key) (one to one)

players
	- id 
	- first name
	- last name
	- date of birth
	- current team (foreign key) (one to many)
	- last team - (one team can have many ex-players)can be null or a foreign key

managers
	- id
	- first name
	- last name
	- current club
	- former club

stadium
	- id
	- name
	- capacity
	- team (foreign key) (one to one)

fixtures:
	- id
	- date
	- team 1
	- team 2
	- stadium
	- referee

referee:
	- id
	- first name
	- last name
```