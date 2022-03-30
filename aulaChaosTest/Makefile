GO ?= go

.PHONY: help
## Este comando de ajuda
help:
	@printf "Opções de comandos\n\n"
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "make %-30s ## %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.PHONY: build
## Executa todos os testes em tcserver_test.go
build:
	@$(GO) mod tidy
	@$(GO) clean -testcache
	@$(GO) test github.com/kempertrasdesclub/AulaTestes/aulaChaosTest/localDevOps -test.v -test.paniconexit0 -test.run TestLocalDevOps
