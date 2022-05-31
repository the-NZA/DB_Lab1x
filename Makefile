include .env
export

APP = dblab
BUILD_FLAGS = -v

# Backend commands
.PHONY: build
build:
	go build $(BUILD_FLAGS) ./backend/cmd/$(APP)

.PHONY: run
run:
	go run $(BUILD_FLAGS) ./backend/cmd/$(APP)

.PHONY: clean
clean:
	rm ./$(APP)

.DEFAULT_GOAL := build

# Frontend commands
.PHONE: frontend_install
frontend_install:
	cd frontend && npm install

.PHONY: start
start:
	cd frontend && npm run dev

.PHONY: buildf
buildf:
	cd frontend && npm run build

# Docker commands
.PHONY: docker_psql
docker_psql:
	docker exec -it lab_db psql -U ${POSTGRES_USER} ${POSTGRES_DB}

.PHONY: docker_start
docker_start:
	docker compose -f ${COMPOSE_FILE} up

.PHONY: docker_stop
docker_stop:
	docker compose -f ${COMPOSE_FILE} down
