include .env

migrate:
	@echo "Create migration files..."
	migrate create -ext sql -dir ./migrations -seq create_users
	migrate create -ext sql -dir ./migrations -seq create_profiles
	@echo "Done!"

migrate_up:
	@echo "Starting migration up..."	
	migrate -database ${POSTGRES_URI}?sslmode=disable -path ./migrations up
	@echo "Done!"	

migrate_down:
	@echo "Starting migration down..."	
	migrate -database ${POSTGRES_URI}?sslmode=disable -path ./migrations down
	@echo "Done!"		

docker_up:
	@echo "Stopping all containers..."
	cd Docker && docker compose down 
	@echo "Starting all containers..."
	cd Docker && docker compose up --build
	@echo "Done!"	

docker_dev_up:
	@echo "Stopping all containers..."
	cd Docker && docker compose down 
	@echo "Starting only postgres container..."
	cd Docker && docker compose up -d postgres
	@echo "Done!"	

docker_dev_build:
	@echo "Stopping all containers..."
	cd Docker && docker compose down 
	@echo "Starting only postgres container..."
	cd Docker && docker compose up -d --build 
	@echo "Done!"		

docker_down:
	@echo "Stopping all containers..."
	cd Docker && docker compose down --volumes
	@echo "Done!"	

encrypt:
	@echo "Encrypt .env File"
	sops --encrypt --pgp 87CBFE717AF74872F700D544D2511A4B048A790D .env > .env.enc
	@echo "Done!"

decrypt:
	@echo "Deccrypt .env File"	
	sops --decrypt .env.enc > .env
	@echo "Done!"