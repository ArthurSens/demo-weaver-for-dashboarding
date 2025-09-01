PODMAN = $(shell if command -v podman &>/dev/null; then echo podman; else echo docker; fi)

.PHONY: all
all: clients dashboards

.PHONY: clients
clients: client-go

.PHONY: client-go
client-go:
	$(PODMAN) run -it --rm -v $(PWD):/work:Z -u 0 --workdir /work otel/weaver registry generate -r ./semconv/ go ./generated/client/go 
.PHONY: dashboards
dashboards:
	$(PODMAN) run -it --rm -v $(PWD):/work:Z -u 0 --workdir /work otel/weaver registry generate -r ./semconv/ dashboards ./generated/dashboards
