# RestAPI-todolist
[![forthebadge made-with-go](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://go.dev/)

## Dependency in this GO program
```sh
github.com/gin-gonic/gin
github.com/joho/godotenv
github.com/dgrijalva/jwt-go
github.com/golang/crypto
github.com/mashingan/smapping
github.com/google/uuid
github.com/lib/pq
```

## Installation

* git clone https://github.com/kelvinnich/RestAPI-todolist.git
* cd RestAPI-todolist
* Edit .env
* Set your database connection details
* go run main.go
 
# API Endpoints

| Method | Endpoints |    Description     |
| ------ | ------ | ------- |
| POST | /register | Register account must given name,email,password to body request
| POST | /login | login account must given email & password to body request
| GET | /users/profile | Get users
| GET |  /users/:id | Get Customer By Id
| POST | /todolist/addTodolist | Add todolist
| PUT | /todolist/updateTodo/:id | Update Todolist by id
| DELETE | /todolist/:id | Delete todolist by id
| GET | /todolist/ | Get All Todolist
| GET | /todolist/:name | Get Todolist By Name
| GET | /todolist/:status | Get Todolist By status
