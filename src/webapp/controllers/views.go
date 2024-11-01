package controllers

import (
	"net/http"
	"webapp/components"
	"webapp/utils"

	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {
	return utils.Render(c, http.StatusOK, components.Home())
}
