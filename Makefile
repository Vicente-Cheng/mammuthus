SCRIPT_TARGETS := $(shell ls scripts | grep -v generate)

.dapper:
	@echo Downloading dapper
	@curl -sL https://releases.rancher.com/dapper/latest/dapper-$$(uname -s)-$$(uname -m) > .dapper.tmp
	@@chmod +x .dapper.tmp
	@./.dapper.tmp -v
	@mv .dapper.tmp .dapper

$(SCRIPT_TARGETS): .dapper
	./.dapper $@

# Code generation targets
generate: .dapper
	@echo "Generating Kubernetes client code using Dapper..."
	./.dapper generate

# Clean generated code
clean-generate:
	@echo "Cleaning generated code..."
	rm -rf pkg/client pkg/apis/freezeio.dev/v1beta1/zz_generated.deepcopy.go

# Regenerate everything
regenerate: clean-generate generate

.DEFAULT_GOAL := default

.PHONY: $(SCRIPT_TARGETS) generate clean-generate regenerate
