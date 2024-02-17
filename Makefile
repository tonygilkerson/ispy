NOW := $(shell echo "`date +%Y-%m-%d`")

#
# Display help
# 
define help_info
	@echo "\nUsage:\n"
	@echo ""
	@echo "  $$ make test       - Run all KUTTL tests locally, this will create a kwok-cluster"
	@echo ""


endef

help:
	$(call help_info)

# export KUBECONFIG := ./testing/kwok-kubeconfig.yaml
ociBuild:
	podman build . --file Dockerfile --tag ispy:dev

	