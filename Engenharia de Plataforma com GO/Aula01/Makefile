APP_NAME = configparser
BIN_DIR = bin
SRC = ./main.go
CONFIG_FILE = ./config.yaml

.PHONY: all build test run clean docker-build docker-run

all: build

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) $(SRC)

test:
	go test ./...

run-parse:
	$(BIN_DIR)/$(APP_NAME) parse --file $(CONFIG_FILE)

run-server:
	$(BIN_DIR)/$(APP_NAME) server --file $(CONFIG_FILE)

clean:
	rm -rf $(BIN_DIR)

docker-build:
	docker build --no-cache -t $(APP_NAME) .

docker-run-parse:
	docker run --rm -v "$(PWD)/config.yaml:/app/config.yaml" $(APP_NAME) parse --file /app/config.yaml

docker-run-server:
	docker run --rm -v "$(PWD)/config.yaml:/app/config.yaml" $(APP_NAME) server --file /app/config.yaml