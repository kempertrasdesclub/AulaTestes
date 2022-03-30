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

.PHONY: build-no-parallels
## Executa o teste normalMapNoParallels_test.go
build-no-parallels:
	@$(GO) mod tidy
	@$(GO) clean -testcache
	@$(GO) test github.com/kempertrasdesclub/AulaTestes/aulaTests -test.v -test.paniconexit0 -test.run TestNormalMapNoParallels

.PHONY: build-map-parallels-correct
## Executa o teste normalMapParallelsCorrect_test.go
build-map-parallels-correct:
	@$(GO) mod tidy
	@$(GO) clean -testcache
	@$(GO) test github.com/kempertrasdesclub/AulaTestes/aulaTests -test.v -test.paniconexit0 -test.run TestNormalMapParallelsCorrect

.PHONY: build-map-parallels-error
## Executa todos os testes em normalMapParallelsError_test.go
build-map-parallels-error:
	@$(GO) mod tidy
	@$(GO) clean -testcache
	@$(GO) test github.com/kempertrasdesclub/AulaTestes/aulaTests -test.v -test.paniconexit0 -test.run TestNormalMapParallelsError

.PHONY: build-parallels-syncmap
## Executa todos os testes em parallelsSyncMap_test.go
build-parallels-syncmap:
	@$(GO) mod tidy
	@$(GO) clean -testcache
	@$(GO) test github.com/kempertrasdesclub/AulaTestes/aulaTests -test.v -test.paniconexit0 -test.run TestParallelsSyncMap

.PHONY: build-pass-tests
## Executa todos os testes que passam
build-pass-tests:
	@$(GO) clean -testcache
	@$(MAKE) build-no-parallels
	@$(MAKE) build-map-parallels-correct
	@$(MAKE) build-parallels-syncmap

.PHONY: build-fail-tests
## Executa todos os testes que falham
build-fail-tests:
	@$(GO) clean -testcache
	@$(MAKE) build-map-parallels-error
