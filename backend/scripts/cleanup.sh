#!/bin/bash

echo "Cleaning up resources..."

# Function to clean up Docker containers
cleanup_docker() {
    echo "Stopping and removing Docker containers..."
    docker ps -aq | while read -r container_id; do
        echo "Stopping container $container_id..."
        docker stop "$container_id" > /dev/null 2>&1 || echo "Failed to stop container $container_id"
        
        echo "Removing container $container_id..."
        docker rm "$container_id" > /dev/null 2>&1 || echo "Failed to remove container $container_id"
    done

    echo "Pruning unused Docker resources..."
    docker system prune -f > /dev/null 2>&1
}

# Function to clean up Kubernetes resources
cleanup_kubernetes() {
    echo "Deleting Kubernetes resources..."
    kubectl delete pod,svc --all > /dev/null 2>&1 || echo "Failed to delete Kubernetes pods or services"
}

# Function to remove temporary project files or logs
cleanup_files() {
    echo "Removing temporary files and logs..."
    rm -rf ./tmp ./logs ./debug.log > /dev/null 2>&1 || echo "Failed to remove temporary files or logs"
}

# Cleanup Docker if Docker is enabled
if command -v docker &> /dev/null; then
    cleanup_docker
else
    echo "Docker is not installed or available in PATH. Skipping Docker cleanup."
fi

# Cleanup Kubernetes if Kubernetes is enabled
if command -v kubectl &> /dev/null; then
    cleanup_kubernetes
else
    echo "kubectl is not installed or available in PATH. Skipping Kubernetes cleanup."
fi

# Perform general file cleanup
cleanup_files

echo "Cleanup completed."
