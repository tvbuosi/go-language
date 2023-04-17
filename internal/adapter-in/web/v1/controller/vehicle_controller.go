package controller

import (
	"errors"
	"net/http"

	"hex-golang/internal/application/domain"
	"hex-golang/internal/application/port"

	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
)

type VehicleController struct {
	CreateVehicleUseCase port.CreateVehicleUseCase
}

func NewVehicleController() *VehicleController {
	return &VehicleController{}
}

const (
	base_path = "/vehicles"
)

func (vc VehicleController) Register(server *echo.Echo) {
	v1 := server.Group("v1")
	v1.POST(base_path, vc.Insert)
}

func (vc *VehicleController) Insert(c *echo.Context) (*uuid.UUID, error) {
	vehicle := domain.Vehicle{}

	if err := c.Bind(&vehicle); err != nil {
		return nil, errors.New("Failed to parse vehicle input")
	}

	id, err := vc.CreateVehicleUseCase.Execute(&vehicle)
	if err != nil {
		return nil, errors.New("Failed to create vehicle input")
	}

	c.JSON(http.StatusCreated, id)
}
