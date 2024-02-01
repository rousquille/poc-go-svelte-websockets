version := "0.1.0"

PACKAGE := github.com/rousquille/poc-go-svelte-websockets/internal/cli

.PHONY: backend
backend:
	env CGO_ENABLED=0 go build -o ./poc-go-svelte-websockets -ldflags "-X $(PACKAGE).GlobalVersion=$(version) -w -s" ./cmd

.PHONY: frontend
frontend:
	npm install --prefix ./web
	npm run --prefix ./web build
	rm -rf ./internal/api/dist/
	mv ./web/dist ./internal/api/

.PHONY: build
build: frontend backend

.PHONY: clean
clean:
	rm ./poc-go-svelte-websockets
