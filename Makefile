.PHONY: build

BINARY_NAME=fin_app

build:
	go mod tidy && \
		templ generate && \
	go build -o ./bin/${BINARY_NAME} ./cmd/main.go && \
	air -c air.toml

clean:
	go clean

