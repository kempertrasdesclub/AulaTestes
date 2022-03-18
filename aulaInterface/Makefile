GO ?= go

.PHONY: help
## Este comando de ajuda
help:
	@printf "Opções de comandos\nmake commando\n\n"
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "make %-30s ## %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.PHONY: build-plugins
## Gera todos os plugins
build-plugins:
	@$(MAKE) build-plugin-user-fake
	@$(MAKE) build-plugin-user-mongodb
	@$(MAKE) build-plugin-user-sqlite

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

.PHONY: build-plugin-user-fake
## Gera o binário do plugin user fake
build-plugin-user-fake:
	@$(GO) mod tidy
	@$(GO) build -buildmode=plugin -o ./externalPlugin/user.fake.so ./plugin/fake.user

.PHONY: build-site
## Gera o binário do site
build-site:
	@$(GO) build -o ./main.so ./main.go
	@chmod +X ./main.so
	./main.so

.PHONY: build
## Build completo
build:
	@$(MAKE) build-plugins
	@$(MAKE) build-site

.PHONY: clean-sqlite
## Apaga o banco de dados SQLite
clean-sqlite:
	@rm ./database.sqlite

.PHONY: clean-binaries
## Limpa os arquivos binários
clean-binaries:
	@rm ./externalPlugin/user.sqlite.so
	@rm ./externalPlugin/user.mongodb.so
	@rm ./main.so

.PHONY: clean-all
## Limpa os arquivos de teste
clean-all:
	@$(MAKE) clean-binaries
	@$(MAKE) clean-sqlite
