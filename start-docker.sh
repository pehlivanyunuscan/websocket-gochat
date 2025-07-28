#!/bin/bash

set -e

echo "Stopping and removing old containers..."
sudo docker-compose down
sudo docker-compose rm -f $(sudo docker-compose ps -aq)

echo "Pruning unused volumes and networks..."
sudo docker volume prune -f
sudo docker network prune -f

echo "Building Docker image and starting websocket-gochat service..."
sudo docker-compose up --build -d

echo "WebSocket GoChat is running"
