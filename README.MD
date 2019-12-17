## Gin's RequestID Middleware

[Gin](https://github.com/gin-gonic/gin)'s is a middleware that injects a 'RequestID' into the context and header of each request with enhanced request logging.

## Contents

- [Installation](#installation)
- [Quick start](#quick-start)

## Installation

To install Gin package, you need to install Go and set your Go workspace first.

The first need [Go](https://golang.org/) installed (**version 1.11+ is required**), then you can use the below Go command to install Gin & the RequestID Middleware.

```sh
$ go get -v -u github.com/gin-gonic/gin
$ go get -v -u github.com/sumit-tembe/gin-requestid
```

## Quick start

###### RequestID Middleware with all default configs.

```sh
# assume the following code is in example.go file
$ cat example.go
```

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	requestid "github.com/sumit-tembe/gin-requestid"
)

func main() {
	// without any middlewares
	router := gin.New()

	// Middlewares
	{
		//recovery middleware
		router.Use(gin.Recovery())
		//middleware which injects a 'RequestID' into the context and header of each request.
		router.Use(requestid.RequestID(nil))
		//middleware which enhance Gin request logger to include 'RequestID'
		router.Use(gin.LoggerWithConfig(requestid.GetLoggerConfig(nil, nil, nil)))
	}

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Golang Otaku!")
	})

	router.Run(":8080")
}

```
```
# run example.go and visit 0.0.0.0:8080/ on browser
$ go run example.go
```

###### Output Logs:
```
[GIN-debug] 2019-12-16T18:50:49+05:30 [bzQg6wTpL4cdZ9bM] - "GET /"
[GIN-debug] 2019-12-16T18:50:49+05:30 [bzQg6wTpL4cdZ9bM] - [::1] "GET / HTTP/1.1 200 22.415µs" Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36
```
###### RequestID Middleware with user specified UUID generator function.
```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	requestid "github.com/sumit-tembe/gin-requestid"
	"utils"
)

func main() {
	// without any middlewares
	router := gin.New()

	// Middlewares
	{
		//recovery middleware
		router.Use(gin.Recovery())
		//middleware which injects a 'RequestID' into the context and header of each request.
		router.Use(requestid.RequestID(utils.GenerateULID))
		//middleware which enhance Gin request logger to include 'RequestID'
		router.Use(gin.LoggerWithConfig(requestid.GetLoggerConfig(nil, nil, nil)))
	}

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Golang Otaku!")
	})

	router.Run(":8080")
}
```
```
# run example.go and visit 0.0.0.0:8080/ on browser
$ go run example.go
```
###### Output Logs:
```
[GIN-debug] 2019-12-16T19:02:59+05:30 [01DW7EJYEKRXGJRGWERQ2QHWTZ] - "GET /"
[GIN-debug] 2019-12-16T19:02:59+05:30 [01DW7EJYEKRXGJRGWERQ2QHWTZ] - [::1] "GET / HTTP/1.1 200 22.415µs" Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36

```

##### Tip: It's a good practice to pass context to all functions so that we can log with RequestID.