![alt text](image.png)
![alt text](image-1.png)
![alt text](image-2.png)

## gin-package: https://github.com/gin-gonic/gin
1. first create a module: git mod init example.com/project
2. go get -u github.com/gin-gonic/gin


## Download sqlite3:
go get github.com/mattn/go-sqlite3

### To use the package:
`package db

import (
	"database/sql" // we will use this package to interact with the database
	_"github.com/mattn/go-sqlite3" // it must be imported, but will not be used directly, go will use it under the hood
)
`
* underscore tells that we will not use it directly.

## to install a package for password hashing:
> go get -u golang.org/x/crypto