
GO ?= go
GOTESTSUM ?= gotestsum

.PHONY: build
## Monta o plugin
build:
	@$(GO) mod tidy
	@$(GO) build -buildmode=plugin -o ../../cmd/plugin/user.sqlite.so