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

# License
[![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/finnomena/go-fnerror/blob/master/LICENSE)
