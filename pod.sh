#!/bin/bash

# Build the Docker image for the Go application using Podman
podman build -t my-go-app-image .

# Run the Go application container
podman run -d --name my-go-app-container my-go-app-image
