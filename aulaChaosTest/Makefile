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

.PHONY: build-normal
## Executa o teste TestLocalDevOps sem habilitar o teste de caos, sem log de dados na saída padrão
build-normal:
	@rm -rf ./localDevOps/memory
	@$(GO) mod tidy
	@$(GO) clean -testcache
	@CHAOS_TEST=0 LOG=0 $(GO) test github.com/kempertrasdesclub/AulaTestes/aulaChaosTest/localDevOps -test.v -test.paniconexit0 -test.run TestLocalDevOps

.PHONY: build-chaos
## Executa o teste TestLocalDevOps habilitando o teste de caos, sem log de dados na saída padrão
build-chaos:
	@rm -rf ./localDevOps/memory
	@$(GO) mod tidy
	@$(GO) clean -testcache
	@CHAOS_TEST=1 LOG=0 $(GO) test github.com/kempertrasdesclub/AulaTestes/aulaChaosTest/localDevOps -test.v -test.paniconexit0 -test.run TestLocalDevOps

.PHONY: build-normal-log
## Executa o teste TestLocalDevOps sem habilitar o teste de caos, com log de dados na saída padrão
build-normal-log:
	@rm -rf ./localDevOps/memory
	@$(GO) mod tidy
	@$(GO) clean -testcache
	@CHAOS_TEST=0 LOG=1 $(GO) test github.com/kempertrasdesclub/AulaTestes/aulaChaosTest/localDevOps -test.v -test.paniconexit0 -test.run TestLocalDevOps

.PHONY: build-chaos-log
## Executa o teste TestLocalDevOps habilitando o teste de caos, com log de dados na saída padrão
build-chaos-log:
	@rm -rf ./localDevOps/memory
	@$(GO) mod tidy
	@$(GO) clean -testcache
	@CHAOS_TEST=1 LOG=1 $(GO) test github.com/kempertrasdesclub/AulaTestes/aulaChaosTest/localDevOps -test.v -test.paniconexit0 -test.run TestLocalDevOps
