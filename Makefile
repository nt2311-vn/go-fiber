.PHONY: build

BINARY_NAME=fin_app
ifeq ($(OS),Windows_NT)
	BINARY_NAME=fin_app.exe
endif

build:
	go mod tidy && \
		templ generate && \
		go generate ./... && \
		bunx tailwindcss build -i static/css/style.css -o static/css/tailwindcss.css && \
	go build -o ./bin/${BINARY_NAME} ./cmd/web/fiber_main.go && \
	air -c air.toml

clean:
	go clean

