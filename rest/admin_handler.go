package rest

import (
	"auth-microservice/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	service domain.AdminService
}

func NewAdminHandler(s domain.AdminService) *AdminHandler {
	return &AdminHandler{service: s}
}

func (h AdminHandler) HandleGetAdmins(c echo.Context) error {
	list, err := h.service.ListAdmins(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"Error getting Admins": err.Error()})
	}
	return c.JSON(http.StatusOK, list)
}

func (h AdminHandler) HandleSaveAdmin(c echo.Context) error {
	var r domain.SaveParams

	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error:": err.Error()})
	}

	if err := h.service.SaveAdmin(c.Request().Context(), r); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error_message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Admin saved!"})
}

func (h AdminHandler) HandleDeleteAdmin(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error deleting Admin:": err.Error()})
	}

	if err := h.service.RemoveAdmin(c.Request().Context(), domain.AdminID(id)); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error while removing": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Admin deleted"})
}
