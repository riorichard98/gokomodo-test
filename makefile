# Target: create_network
# Description: Creates the Docker network named "gokomodo" if it doesn't exist.
create_network:
	docker network create gokomodo || true

# Target: run_postgres
# Dependencies: create_network
# Description: Removes any existing "gokomodo-pg" container, creates a new PostgreSQL container named "gokomodo-pg" connected to the "gokomodo" network, and initializes it with the specified database and user credentials.
run_postgres: create_network
	docker rm -f gokomodo-pg || true
	docker run -d \
		--name gokomodo-pg \
		--network gokomodo \
		-e POSTGRES_DB=shop \
		-e POSTGRES_USER=gokomodo \
		-e POSTGRES_PASSWORD=12345 \
		-p 5432:5432 \
		-v $(PWD)/init.sql:/docker-entrypoint-initdb.d/init.sql \
		postgres:latest

# Target: build
# Description: Build the Docker image for the Go application.
build:
	docker build -t my-golang-app .

# Target: run_server
# Dependencies: build
# Description: Removes any existing "my-golang-app" container, creates a new container named "my-golang-app" connected to the "gokomodo" network, and runs the Go application inside it.
run_server: build create_network
	docker rm -f my-golang-app || true
	docker run -d \
		--name my-golang-app \
		--network gokomodo \
		-p 3001:3001 \
		my-golang-app

# Target: clean
# Description: Clean up Docker containers and images.
clean:
	docker rm -f gokomodo-pg || true
	docker rm -f my-golang-app || true
	docker rmi my-golang-app

run-all: run_postgres run_server

.PHONY: create_network run_postgres build run_server clean
