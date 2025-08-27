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

## Demo applications

Run the demo backend application (requires generating the client library first):
```
go run ./applications/backend/main.go
```
