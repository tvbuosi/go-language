package domain

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
)

type Vehicle struct {
	ID          uuid.UUID           `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Type        string              `json:"type"`
	Rate        float64             `json:"rate"`
	CreatedAt   timestamp.Timestamp `json:"created_at"`
	CreatedBy   timestamp.Timestamp `json:"created_by"`
	UpdatedAt   timestamp.Timestamp `json:"updated_at"`
	UpdatedBy   timestamp.Timestamp `json:"updated_by"`
}
