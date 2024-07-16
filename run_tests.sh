#!/bin/bash

# Exit on any error
set -e

# Pull Docker images and start services
docker-compose up -d

# Run tests inside the Docker container
docker-compose exec -T api go test ./...

# Clean up
docker-compose down
