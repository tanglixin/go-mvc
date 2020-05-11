package person

import (
	"github.com/echo"
	"net/http"
)

func PersonHandle(c echo.Context )error  {
	return c.String(http.StatusOK,"person list")
}
