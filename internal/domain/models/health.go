package models

import "time"

type HealthStatus string

const (
	StatusPass HealthStatus = "pass"
	StatusWarn HealthStatus = "warn"
	StatusFail HealthStatus = "fail"
)

type ComponentType string

const (
	TypeDatastore ComponentType = "datastore"
	TypeBroker    ComponentType = "broker"
)

type HealthCheck struct {
	Status      HealthStatus
	Description string
	Version     string
	Checks      []*ComponentCheck
}

type ComponentCheck struct {
	Name          string
	Type          ComponentType
	Status        HealthStatus
	ObservedValue float64
	ObservedUnit  string
	Time          time.Duration
	Output        string
}

func (c ComponentCheck) IsFail() bool {
	return c.Status != StatusPass
}

func (t ComponentType) IsFatal() bool {
	return t == TypeDatastore
}
