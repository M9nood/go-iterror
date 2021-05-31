# go-iterror

go-iterror is a package for manage HTTP errors which provides flexibility to use and customize.

### Installation
```bash
go get github.com/M9nood/go-iterror
```

### Example
```golang
package main

import (
	"fmt"

	iterror "github.com/M9nood/go-iterror"
)

func main() {
  err := iterror.NewError(iterror.BadRequestException, "Invalid Request")
  fmt.Println("error code :", err.GetCode())
  // error code : 400
  fmt.Println("error http :", err.GetHttpCode())
  // error http : 400
  fmt.Println("error name :", err.GetName())
  // error name : BadRequestException
  fmt.Println("error message :", err.Error())
  // error message : Invalid Request

  errCustom := iterror.NewErrorCustomHttpCode(iterror.BadRequestException, 200, "Invalid Request")
  fmt.Println("error custom code :", errCustom.GetCode())
  // error custom code : 400
  fmt.Println("error custom http :", errCustom.GetHttpCode())
  // error custom http : 200
  fmt.Println("error custom name :", errCustom.GetName())
  // error custom name : BadRequestException
  fmt.Println("error custom message :", errCustom.Error())
  // error custom message : Invalid Request
}
```

### Variables

"400" - BadRequestException  
"401" - UnauthorizedException  
"403" - ForbiddenException  
"404" - NotFoundException  
"408" - RequestTimeoutException  
"409" - ConflictException  
"429" - ManyRequestException  
"500" - InternalServerErrorException  

### Use case
In a project structure example as 
```
|_ controller
|   |_ sample.go
|_ service
|   |_ sample.go
|_ main.go
|_ go.mod
|_ go.sum
```

`/service/sample.go`
```golang
package service

import (
	iterror "github.com/M9nood/go-iterror"
)

func GetSampleService() (string, iterror.ErrorException) {
	// do something
	err := iterror.NewError(iterror.BadRequestException, "Invalid data")
	if err != nil {
		return "", err
	}
	return "Sample data", nil
}
```


`/controller/sample.go`
```golang
package controller

import (
	"m9error/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetSampleHandler(c echo.Context) error {
	sample, err := service.GetSampleService()
	if err != nil {
        // can use err.GetCode() for get error code
		return c.JSON(err.GetHttpCode(), err)
	}
	return c.JSON(http.StatusOK, sample)
}
```


`/main.go`
```golang
package main

import (
	"m9error/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", controller.GetSampleHandler)

	e.Logger.Fatal(e.Start(":4000"))
}
```



# License
[![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/M9nood/go-iterror/blob/master/LICENSE)
