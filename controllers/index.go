package controllers

import (
	"github.com/echo"
	"net/http"
)

func IndexHandle(c echo.Context) error {
	return c.String(http.StatusOK,"index controller")
}
