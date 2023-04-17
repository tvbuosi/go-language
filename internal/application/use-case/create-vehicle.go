package usecase

import (
	"hex-golang/internal/application/domain"

	"github.com/google/uuid"
)

type CreateVehicle struct {
}

func New() {
	return &CreateVehicle
}

func (cv *CreateVehicle) Execute(vehicle *domain.Vehicle) (*string, error) {
	return uuid.New()
}
