.PHONY: build

BINARY_NAME=fin_app
POCKETBASE_BINARY=fin_db
ifeq ($(OS),Windows_NT)
	BINARY_NAME=fin_app.exe
	POCKETBASE_BINARY=fin_db.exe
endif

build:
	go mod tidy && \
	bunx tailwindcss build -i static/css/style.css -o static/css/tailwindcss.css && \
	go build -o ./bin/${BINARY_NAME} ./cmd/web/fiber_main.go

build-db:
	go build -o ./bin/${POCKETBASE_BINARY} ./cmd/db/pb_main.go

build-all: build build-db

run-fiber:
	air -c air.toml

run-db:
	./bin/${POCKETBASE_BINARY} serve --http=127.0.0.1:8080 --dir=./pocketbase/pb_data/

run: run-fiber run-db

clean:
	go clean
	rm -rf ./bin/${BINARY_NAME} ./bin/${POCKETBASE_BINARY}

