## Golang Api Gin Gonic
    Build API Auths, and Books User Use JWT, SQL, GIN GONIC

### Generate API DOCS
    pull swagger repo

    $ go get -v github.com/swaggo/swag/cmd/swag
    $ go get -v github.com/swaggo/gin-swagger
    $ go get -v github.com/swaggo/files

    Generate
    swag init -g main.go_files
    
    docs api :
    http://localhost:8080/swagger/index.html#

### API Docs by Postman
    https://documenter.getpostman.com/view/6097899/UVsLT7LB