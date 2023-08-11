ADMIN_BINARY=adminApp
GATEWAY_BINARY=apiApp
CARTS_BINARY=cartsApp
PRODUCTS_BINARY=productsApp
USERS_BINARY=usersAPP

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"


## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"


## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_admin build_gateway build_carts build_products build_users
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

build_admin:
			@echo "Building admin service bianry..."
			cd ./moviesGo-admins-service && env GOOS=linux CGO_ENABLED=0 go build -o ${ADMIN_BINARY} ./cmd
			@echo "Done"

build_gateway:
			@echo "Building api gateway bianry..."
			cd ./moviesGo-api-gateway && env GOOS=linux CGO_ENABLED=0 go build -o ${GATEWAY_BINARY} ./cmd
			@echo "Done"

build_carts:
			@echo "Building carts service bianry..."
			cd ./moviesGo-carts-service && env GOOS=linux CGO_ENABLED=0 go build -o ${CARTS_BINARY} ./cmd
			@echo "Done"

build_products:
			@echo "Building products service bianry..."
			cd ./moviesGo-products-service && env GOOS=linux CGO_ENABLED=0 go build -o ${PRODUCTS_BINARY} ./cmd
			@echo "Done"

build_users:
			@echo "Building users service bianry..."
			cd ./moviesGo-users-service && env GOOS=linux CGO_ENABLED=0 go build -o ${USERS_BINARY} ./cmd
			@echo "Done"