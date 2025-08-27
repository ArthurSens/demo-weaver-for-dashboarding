# demo-weaver-for-dashboarding

## Client SDKs

To generate the go client for use in the backend application run the following command:

```bash
weaver registry generate --templates=./templates --registry ./semconv go ./generated/client/go
```

You can add the `--param included_namespaces="db,http"` flag to limit generation to specific namespaces

## Grafana dashboards

To generate dashboards:

```bash
weaver registry generate --templates=./templates --registry ./semconv go ./generated/dashboards
```

The `included_namespaces` param also works for dashboards

## Demo applications

Run the demo backend application (requires generating the client library first):
```
go run ./applications/backend/main.go
```
