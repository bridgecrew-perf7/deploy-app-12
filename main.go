package main

import (
	"fmt"
	"net/http"
	"path"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("Go Program")
	server := echo.New()
	server.GET(path.Join("/"), Version)
	
	server.Start(":1323")

}
func Version(context echo.Context) error {
	return context.JSON(http.StatusOK, map[string]interface{}{"version": 2})
}