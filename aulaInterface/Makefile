
GO ?= go

.PHONY: build-plugin-user-mongodb
## Gera o binário do plugin user mongodb
build-plugin-user-mongodb:
	@$(GO) mod tidy
	@$(GO) build -buildmode=plugin -o ./externalPlugin/user.mongodb.so ./plugin/mongodb.user

.PHONY: build-plugin-user-sqlite
## Gera o binário do plugin user sqlite
build-plugin-user-sqlite:
	@$(GO) mod tidy
	@$(GO) build -buildmode=plugin -o ./externalPlugin/user.sqlite.so ./plugin/sqlite.user

.PHONY: build-site
## Gera o binário do site
build-site:
	@$(GO) build -o ./main.so ./main.go
	@chmod +X ./main.so
	./main.so

.PHONY: build
## Build completo
build: build-plugin-user-mongodb build-plugin-user-sqlite build-site
