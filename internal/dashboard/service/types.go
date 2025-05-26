package service

import (
	"time"

	"github.com/docker/docker/client"
)

type ServiceStatus struct {
	Name         string    `json:"name"`
	DockerStatus string    `json:"docker_status"`
	IsRunning    bool      `json:"is_running"`
	HealthStatus string    `json:"health_status"`
	LastChecked  time.Time `json:"last_checked"`
	Error        string    `json:"error,omitempty"`
	Endpoint     string    `json:"endpoint"`
}

type ServiceControl struct {
	dockerClient *client.Client
}
