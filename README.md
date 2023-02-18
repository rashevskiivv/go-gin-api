# API to work with books
This project is a pet project to try Gin. Here's a CRUD for 1 entity.


In project were used the following libraries:
* [Gin](https://gin-gonic.com/) - Routing
* [Viper](https://github.com/spf13/viper) - Managing the environment variables
* [GORM](https://gorm.io/) - ORM
* [MongoDB](https://www.mongodb.com/) - DB
* Shell - Script to initialize DB (start using and create user)
* Docker with volumes + Docker-compose - Containerization

## Docker
1. Clone this repository.
2. Run: ```docker compose up```
3. Stop in different window: ```docker compose down``` OR CTRL+C to terminate the process in the same window.