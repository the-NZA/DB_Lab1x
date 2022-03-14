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
.PHONY: start
start:
	cd frontend && npm run dev

.PHONY: buildf
buildf:
	cd frontend && npm run build