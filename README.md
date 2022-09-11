# movie_users-api
users api for movie app
# ## Major Technologies
- [ ] Go
- [ ] Gin/Gonic
- [ ] MySql
- [ ] Uber-Go/zap

### Install & Setup

To setup and install this sample Leaderboard project, follow the below steps:
- Clone this project by the command: 

```
$ git clone https://github.com/PreetSIngh8929/movie_users-api.git
```
- Then install dependencies using go get 

- Finally, run the below command to start the project.

```
go run main.go
```
# Endpoints

- POST ```/users``` create user
- GET ```"/users/:user_id"``` get user by user id
- PUT ```/users/:user_id``` update user
- POST ```/users/login``` login user
- GET ```/internal/users/search``` search user
- DELETE ```/users/:user_id``` Delete User
