# demo-weaver-for-dashboarding

## Client SDKs

To generate the go client for use in the backend application run the following command:

```bash
make clients
```

## Grafana dashboards

To generate dashboards:

```bash
make dashboards
```

## Running the demo

### Infrastructure

To set up the demo infrastructure, use `docker-compose` (or `podman-compose`) to bring up an ephemeral deployment of Grafana & Prometheus:
```
docker-compose up -d
```

### Applications
Run the demo backend application (requires generating the client library first):
```
go run ./applications/backend/main.go
```
