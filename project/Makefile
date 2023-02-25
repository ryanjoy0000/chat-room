CHAT_BINARY = chatServiceApp
# -------------------------------

up: build_chat
	@echo "stopping docker containers if any..."
	docker-compose down
	@echo "building and starting all services..."
	docker-compose up --build

down:
	@echo "Stopping docker-compose..."
	docker-compose down

# -------------------------------

build_chat:
	@echo "current working directory"
	pwd
	@echo "building chat service -> binary ..."
	cd ../chat-service && env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ${CHAT_BINARY} .