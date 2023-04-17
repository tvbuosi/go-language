package port
package port

type CreateVehicleRepository interface {
	Execute(vehicle *domain.Vehicle) (*string, error)
}