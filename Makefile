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

compile_proto:
	@echo "Compile proto files..."	
	@( \
		cd ./internal/auth/infrastructure/grpc/proto && \
		protoc --go_out=../generated --go-grpc_out=../generated \
		--go_opt=paths=source_relative --go-grpc_opt=paths=source_relative ./registration.proto \
	)
	@echo "Done!"	

encrypt:
	@echo "Encrypt .env File"
	sops --encrypt --pgp 38600674253871759CE3C9CC8BFA5F0299EAA8FF .env > .env.enc
	@echo "Done!"

encrypt_secret:
	@echo "Encrypt .env File"
	cd kubernetes && sops --encrypt --pgp 38600674253871759CE3C9CC8BFA5F0299EAA8FF postgres-secret.yml > postgres-secret.yml.enc
	@echo "Done!"	

decrypt:
	@echo "Decrypt .env File"	
	sops --decrypt .env.enc > .env
	@echo "Done!"

decrypt_secret:
	@echo "Decrypt .env File"	
	cd kubernetes && sops --decrypt postgres-secret.yml.enc > postgres-secret.yml
	@echo "Done!"	

get_private_key:
	@echo "Generate private key for export to Github secrets..."	
	gpg --export-secret-keys --armor 38600674253871759CE3C9CC8BFA5F0299EAA8FF > sops-private.asc
	@echo "Done!"
