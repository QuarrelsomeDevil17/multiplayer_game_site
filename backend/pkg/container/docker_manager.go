package container

import (
	"context"
	"fmt"

	//"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"

	//"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

// DockerManager manages Docker containers
type DockerManager struct {
	Client *client.Client
}

// NewDockerManager initializes a DockerManager
func NewDockerManager() (*DockerManager, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Docker client: %v", err)
	}
	return &DockerManager{Client: cli}, nil
}

// CreateContainer creates a new Docker container
func (d *DockerManager) CreateContainer(imageName, containerName string, envVars []string) (string, error) {
	ctx := context.Background()

	_, err := d.Client.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to pull image: %v", err)
	}

	resp, err := d.Client.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Env:   envVars,
	}, nil, nil, nil, containerName)
	if err != nil {
		return "", fmt.Errorf("failed to create container: %v", err)
	}

	return resp.ID, nil
}

// StartContainer starts a Docker container
func (d *DockerManager) StartContainer(containerID string) error {
	ctx := context.Background()
	if err := d.Client.ContainerStart(ctx, containerID, types.ContainerStartOptions{}); err != nil {
		return fmt.Errorf("failed to start container: %v", err)
	}
	return nil
}

// StopContainer stops a running Docker container
func (d *DockerManager) StopContainer(containerID string) error {
	ctx := context.Background()
	timeout := 10 * time.Second
	timeoutInt := int(timeout.Seconds()) // Convert time.Duration to seconds as int
	stopOptions := container.StopOptions{
		Timeout: &timeoutInt, // Timeout is now a pointer to an int
	}
	if err := d.Client.ContainerStop(ctx, containerID, stopOptions); err != nil {
		return fmt.Errorf("failed to stop container: %v", err)
	}
	return nil
}

// RemoveContainer removes a Docker container
func (d *DockerManager) RemoveContainer(containerID string) error {
	ctx := context.Background()
	if err := d.Client.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{Force: true}); err != nil {
		return fmt.Errorf("failed to remove container: %v", err)
	}
	return nil
}

// ListContainers lists all running containers
func (d *DockerManager) ListContainers() ([]types.Container, error) {
	ctx := context.Background()
	containers, err := d.Client.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return nil, fmt.Errorf("failed to list containers: %v", err)
	}
	return containers, nil
}
