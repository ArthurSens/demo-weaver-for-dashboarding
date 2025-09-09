PODMAN = $(shell if command -v podman &>/dev/null; then echo podman; else echo docker; fi)

.PHONY: clean
clean:
	rm -rf generated

.PHONY: all
all: clients dashboards

.PHONY: clients
clients: client-go client-python

.PHONY: client-go
client-go:
	$(PODMAN) run -it --rm -v $(PWD):/work:Z -u 0 --workdir /work otel/weaver registry generate -r ./semconv/ go ./generated/client/go 
.PHONY: client-python
client-python:
	$(PODMAN) run -it --rm -v $(PWD):/work:Z -u 0 --workdir /work otel/weaver registry generate -r ./semconv/ python ./generated/client/python 
.PHONY: dashboards
dashboards: foundation-dashboards
	go run ./dashboards/
.PHONY: foundation-dashboards
foundation-dashboards:
	$(PODMAN) run -it --rm -v $(PWD):/work:Z -u 0 --workdir /work otel/weaver registry generate -r ./semconv/ dashboards ./generated/dashboards


GITHUB_REPO=ArthurSens/demo-weaver-for-dashboarding
.PHONY: check
check:
	$(PODMAN) run -it --rm -v $(PWD):/work:Z -u 0 --workdir /work otel/weaver \
	registry check \
	-r ./semconv/ \
  --baseline-registry https://github.com/${GITHUB_REPO}/archive/refs/heads/main.zip[semconv] \
	--future --policy ./semconv/policies

.PHONY: watch-dashboards
watch-dashboards: dashboards
	grafanactl resources serve ./generated/dashboards_json --port 8181
