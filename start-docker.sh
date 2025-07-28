#!/bin/bash

set -e

echo "Stopping and removing old containers..."
sudo docker-compose down

echo "Building Docker image and starting websocket-gochat service..."
sudo docker-compose up --build -d
sudo docker-compose logs -f

echo "WebSocket GoChat is running"