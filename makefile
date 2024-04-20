# Define the run_postgres target to start the PostgreSQL container
run_postgres:
	docker rm -f postgresql || true
	docker run -d \
		--name postgresql \
		-e POSTGRES_DB=shop \
		-e POSTGRES_USER=gokomodo \
		-e POSTGRES_PASSWORD=12345 \
		-p 5432:5432 \
		-v $(PWD)/init.sql:/docker-entrypoint-initdb.d/init.sql \
		postgres:latest



