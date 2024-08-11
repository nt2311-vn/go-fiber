.PHONY: build

BINARY_NAME=fin_app

build:
	go mod tidy && \
		templ generate && \
		bunx tailwindcss build -i static/css/style.css -o static/css/tailwindcss.css && \
	go build -o ./bin/${BINARY_NAME} ./cmd/main.go && \
	air -c air.toml

clean:
	go clean

