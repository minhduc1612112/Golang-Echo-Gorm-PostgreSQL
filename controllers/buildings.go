package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"echo-demo/models"
)

func GetBuildings(c echo.Context) error {
	models.GetRoomsByBuildingId(5)
	return c.String(http.StatusOK, "Buildings...")
}

func GetRoomsByBuildingId(c echo.Context) error {
	id, errParseUint := strconv.ParseUint(c.Param("id"), 10, 64)
	if(errParseUint != nil) {
		return c.JSON(http.StatusBadRequest, errParseUint.Error())
	}
	rooms, err := models.GetRoomsByBuildingId(id)
	if(err != nil) {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, rooms)
}