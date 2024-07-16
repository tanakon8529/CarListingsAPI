#!/bin/bash

# Start Docker Compose
docker-compose up --build -d

# Give some time for services to start
sleep 3

# Run the tests
docker-compose exec api go test ./...

# Capture the test exit code
TEST_EXIT_CODE=$?

# Stop Docker Compose
docker-compose down

# Exit with the test exit code
exit $TEST_EXIT_CODE
