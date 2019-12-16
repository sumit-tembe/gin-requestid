## Gin's RequestID Middleware

[Gin](https://github.com/gin-gonic/gin)'s is a middleware that injects a 'RequestID' into the context and header of each request with enhanced request logging.

## Usage:

RequestID Middleware with all default configs.
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
#####Output Logs:
```
[GIN-debug] 2019-12-16T18:50:49+05:30 [bzQg6wTpL4cdZ9bM] - "GET /"
[GIN-debug] 2019-12-16T18:50:49+05:30 [bzQg6wTpL4cdZ9bM] - [::1] "GET / HTTP/1.1 200 22.415µs" Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36
```
RequestID Middleware with user specified UUID generator function.
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
#####Output Logs:
```
[GIN-debug] 2019-12-16T19:02:59+05:30 [01DW7EJYEKRXGJRGWERQ2QHWTZ] - "GET /"
[GIN-debug] 2019-12-16T19:02:59+05:30 [01DW7EJYEKRXGJRGWERQ2QHWTZ] - [::1] "GET / HTTP/1.1 200 15.569µs" PostmanRuntime/7.20.1 

```