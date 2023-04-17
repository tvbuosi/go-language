package port

type CreateVehicleUseCase struct {
	Execute(vehicle *domain.Vehicle) (*string, error)
}