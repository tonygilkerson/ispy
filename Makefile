NOW := $(shell echo "`date +%Y-%m-%d`")

#
# Display help
# 
define help_info
	@echo "\nUsage:\n"
	@echo ""
	@echo "  $$ make setVersion version=9.9.9     - Used to set the version number."
	@echo ""
	@echo "  $$ make devBuildApply                - Build dev image and load into kind."
	@echo ""
	@echo "  $$ make devApply                     - Install chart into kind."
	@echo ""


endef

help:
	$(call help_info)


setVersion:
	@echo Set version "$(version)";\
	ver="v$(version)" yq e '.image.tag = strenv(ver)' ./charts/ispy/values.yaml --inplace;\
	ver="v$(version)" yq e '.version = strenv(ver)' ./charts/ispy/Chart.yaml  --inplace;\
	ver="v$(version)" yq e '.appVersion = strenv(ver)' ./charts/ispy/Chart.yaml  --inplace;

devBuildApply:
	podman build -t ispy:dev .
	podman save -o .temp/ispy.tar localhost/ispy:dev
	kind load image-archive .temp/ispy.tar
	helm upgrade -i ispy charts/ispy --set image.repository=localhost/ispy --set image.tag=dev --set ingressClassName=nginx --set domain=127.0.0.1.nip.io 

devApply:
	helm upgrade -i ispy charts/ispy --set image.repository=localhost/ispy --set image.tag=dev --set ingressClassName=nginx --set domain=127.0.0.1.nip.io 
	
