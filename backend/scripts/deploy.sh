#!/bin/bash

echo "Deploying application..."

# Function to build and deploy Docker container
deploy_docker() {
    echo "Building Docker image..."
    docker build -t game-image . || { echo "Docker build failed! Exiting."; exit 1; }

    echo "Running Docker container..."
    docker run -d -p 8080:8080 --name game-container game-image || { echo "Failed to start Docker container! Exiting."; exit 1; }

    echo "Docker container deployed successfully."
}

# Function to deploy to Kubernetes
deploy_kubernetes() {
    echo "Applying Kubernetes manifests..."
    kubectl apply -f kubernetes.yaml || { echo "Failed to apply Kubernetes manifests! Exiting."; exit 1; }

    echo "Kubernetes resources deployed successfully."
}

# Function to check prerequisites
check_prerequisites() {
    echo "Checking prerequisites..."
    
    if ! command -v docker &> /dev/null; then
        echo "Docker is not installed or not available in PATH. Exiting."
        exit 1
    fi
    
    if ! command -v kubectl &> /dev/null; then
        echo "kubectl is not installed or not available in PATH. Skipping Kubernetes deployment."
        USE_KUBERNETES=false
    else
        USE_KUBERNETES=true
    fi
    
    echo "All prerequisites satisfied."
}

# Main deployment logic
check_prerequisites

# Deploy Docker container
deploy_docker

# Deploy to Kubernetes if enabled
if [ "$USE_KUBERNETES" = true ]; then
    deploy_kubernetes
else
    echo "Kubernetes deployment skipped as kubectl is not available."
fi

echo "Application deployment completed."
