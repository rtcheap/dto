package dto

import "fmt"

// Status constants
const (
	StatusHealty    ServiceStatus = "HEALTHY"
	StatusUnhealthy ServiceStatus = "UNHEALTHY"
)

// ServiceStatus status string.
type ServiceStatus string

// Service application instance metadata.
type Service struct {
	ID          string        `json:"id,omitempty"`
	Application string        `json:"application,omitempty"`
	Location    string        `json:"location,omitempty"`
	Port        int           `json:"port,omitempty"`
	Status      ServiceStatus `json:"status,omitempty"`
}

func (s Service) String() string {
	return fmt.Sprintf(
		"Service(id=%s, application=%s, location=%s, port=%d, status=%s)",
		s.ID,
		s.Application,
		s.Location,
		s.Port,
		s.Status,
	)
}
