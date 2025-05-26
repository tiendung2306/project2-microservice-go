package service

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func NewServiceControl() (*ServiceControl, error) {
	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithVersion("1.47"),
	)
	if err != nil {
		return nil, fmt.Errorf("không thể kết nối tới Docker: %v", err)
	}

	return &ServiceControl{
		dockerClient: cli,
	}, nil
}

func (sc *ServiceControl) checkServiceHealth(serviceName string) (string, error) {
	userServicePort := os.Getenv("USER_SERVICE_PORT")
	authServicePort := os.Getenv("AUTH_SERVICE_PORT")
	taskServicePort := os.Getenv("TASK_SERVICE_PORT")
	notificationServicePort := os.Getenv("NOTIFICATION_SERVICE_PORT")
	dashboardServicePort := os.Getenv("DASHBOARD_SERVICE_PORT")
	servicePorts := map[string]string{
		"user-service":         userServicePort,
		"auth-service":         authServicePort,
		"task-service":         taskServicePort,
		"notification-service": notificationServicePort,
		"dashboard-service":    dashboardServicePort,
	}

	port, exists := servicePorts[serviceName]
	if !exists {
		return "", fmt.Errorf("không tìm thấy port cho service %s", serviceName)
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(fmt.Sprintf("http://%s:%s/api/health", serviceName, port))
	if err != nil {
		return "unhealthy", fmt.Errorf("không thể kết nối tới service: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "unhealthy", fmt.Errorf("service trả về status code: %d", resp.StatusCode)
	}

	return "healthy", nil
}

func (sc *ServiceControl) GetAllServicesStatus() ([]ServiceStatus, error) {
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "localhost"
	}

	// Danh sách các service cố định
	fixedServices := map[string]string{
		"user-service":         os.Getenv("USER_SERVICE_PORT"),
		"auth-service":         os.Getenv("AUTH_SERVICE_PORT"),
		"task-service":         os.Getenv("TASK_SERVICE_PORT"),
		"notification-service": os.Getenv("NOTIFICATION_SERVICE_PORT"),
	}

	// Lấy danh sách containers đang chạy
	containers, err := sc.dockerClient.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("lỗi khi lấy danh sách containers: %v", err)
	}

	// Tạo map để dễ dàng kiểm tra container đang chạy
	runningContainers := make(map[string]types.Container)
	for _, container := range containers {
		serviceName := container.Names[0][1:]
		runningContainers[serviceName] = container
	}

	var services []ServiceStatus
	// Duyệt qua danh sách service cố định
	for serviceName, port := range fixedServices {
		status := ServiceStatus{
			Name:        serviceName,
			Endpoint:    fmt.Sprintf("http://%s:%s", baseURL, port),
			LastChecked: time.Now(),
		}

		// Kiểm tra xem service có đang chạy không
		if container, exists := runningContainers[serviceName]; exists {
			status.DockerStatus = container.Status
			status.IsRunning = true
			// Check health nếu service đang chạy
			healthStatus, err := sc.checkServiceHealth(serviceName)
			if err != nil {
				status.HealthStatus = "unhealthy"
				status.Error = err.Error()
			} else {
				status.HealthStatus = healthStatus
			}
		} else {
			status.DockerStatus = "stopped"
			status.IsRunning = false
			status.HealthStatus = "stopped"
		}

		services = append(services, status)
	}

	return services, nil
}

func (sc *ServiceControl) RestartService(serviceName string) error {
	err := sc.dockerClient.ContainerRestart(context.Background(), serviceName, container.StopOptions{})
	if err != nil {
		return fmt.Errorf("lỗi khi khởi động lại service %s: %v", serviceName, err)
	}
	return nil
}

func (sc *ServiceControl) StopService(serviceName string) error {
	timeout := 10
	err := sc.dockerClient.ContainerStop(context.Background(), serviceName, container.StopOptions{
		Timeout: &timeout,
	})
	if err != nil {
		return fmt.Errorf("lỗi khi dừng service %s: %v", serviceName, err)
	}
	return nil
}

func (sc *ServiceControl) StartService(serviceName string) error {
	err := sc.dockerClient.ContainerStart(context.Background(), serviceName, container.StartOptions{})
	if err != nil {
		return fmt.Errorf("lỗi khi bật service %s: %v", serviceName, err)
	}
	return nil
}
