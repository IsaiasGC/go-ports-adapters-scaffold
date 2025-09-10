package health

import (
	"time"

	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/domain/models"
)

type HealthCheckDTO struct {
	Status      string               `json:"status" example:"pass"`
	Description string               `json:"description,omitempty" example:"Kafka and SQS health check"`
	Version     string               `json:"version" example:"1.0"`
	Checks      []*ComponentCheckDTO `json:"checks"`
}

type ComponentCheckDTO struct {
	Name          string    `json:"componentName"`
	Type          string    `json:"componentType"`
	ObservedValue float64   `json:"observedValue,omitempty"`
	ObservedUnit  string    `json:"observedUnit,omitempty"`
	Time          time.Time `json:"time"`
	Status        string    `json:"status"`
	Output        string    `json:"output,omitempty"`
}

func toHealthDTO(model *models.HealthCheck) *HealthCheckDTO {
	checks := make([]*ComponentCheckDTO, 0, len(model.Checks))
	for _, c := range model.Checks {
		if c != nil {
			checks = append(checks, toComponentCheckDTO(*c))
		}
	}
	return &HealthCheckDTO{
		Status:      string(model.Status),
		Description: model.Description,
		Version:     model.Version,
		Checks:      checks,
	}
}

func toComponentCheckDTO(model models.ComponentCheck) *ComponentCheckDTO {
	return &ComponentCheckDTO{
		Name:          model.Name,
		Type:          string(model.Type),
		ObservedValue: model.ObservedValue,
		ObservedUnit:  model.ObservedUnit,
		Time:          model.Time,
		Status:        string(model.Status),
		Output:        model.Output,
	}
}
