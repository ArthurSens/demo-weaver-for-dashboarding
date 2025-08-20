package http

// Debug output completed - metric structure discovered:
// - metric.brief: Panel title/description
// - metric.instrument: "gauge", "counter", "histogram", "updowncounter"
// - metric.metric_name: Metric name for PromQL expressions

// Import the appropriate Grafana Foundation SDK packages
import (
	"encoding/json"
	"log"

	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
)

func main() {
	prometheusDataSourceRef := dashboard.DataSourceRef{
		Type: cog.ToPtr("prometheus"),
		Uid:  cog.ToPtr("prometheus"),
	}

	// Define our dashboard as strongly typed code
	builder := dashboard.NewDashboardBuilder("http").
		Uid("http-dashboard").
		Tags([]string{"built-with-weaver"}).
		Refresh("5m").
		Time("now-1h", "now").
		Timezone(common.TimeZoneBrowser).
		WithPanel(
			timeseries.NewPanelBuilder().
				Title("Number of active HTTP requests.").
				Datasource(prometheusDataSourceRef).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("sum(http.client.active_requests)").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("__auto"),
				),
		).
		WithPanel(
			timeseries.NewPanelBuilder().
				Title("The duration of the successfully established outbound HTTP connections.").
				Datasource(prometheusDataSourceRef).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("histogram_quantile(0.99, rate(http.client.connection.duration_bucket[$__rate_interval]))").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("p99"),
				).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("histogram_quantile(0.90, rate(http.client.connection.duration_bucket[$__rate_interval]))").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("p90"),
				).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("histogram_quantile(0.50, rate(http.client.connection.duration_bucket[$__rate_interval]))").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("p50"),
				).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("rate(http.client.connection.duration_sum[$__rate_interval]) / rate(http.client.connection.duration_count[$__rate_interval])").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("average"),
				),
		).
		WithPanel(
			timeseries.NewPanelBuilder().
				Title("Number of outbound HTTP connections that are currently active or idle on the client.").
				Datasource(prometheusDataSourceRef).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("sum(http.client.open_connections)").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("__auto"),
				),
		).
		WithPanel(
			timeseries.NewPanelBuilder().
				Title("Size of HTTP client request bodies.").
				Datasource(prometheusDataSourceRef).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("histogram_quantile(0.99, rate(http.client.request.body.size_bucket[$__rate_interval]))").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("p99"),
				).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("histogram_quantile(0.90, rate(http.client.request.body.size_bucket[$__rate_interval]))").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("p90"),
				).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("histogram_quantile(0.50, rate(http.client.request.body.size_bucket[$__rate_interval]))").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("p50"),
				).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("rate(http.client.request.body.size_sum[$__rate_interval]) / rate(http.client.request.body.size_count[$__rate_interval])").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("average"),
				),
		).
		WithPanel(
			timeseries.NewPanelBuilder().
				Title("Duration of HTTP client requests.").
				Datasource(prometheusDataSourceRef).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("histogram_quantile(0.99, rate(http.client.request.duration_bucket[$__rate_interval]))").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("p99"),
				).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("histogram_quantile(0.90, rate(http.client.request.duration_bucket[$__rate_interval]))").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("p90"),
				).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("histogram_quantile(0.50, rate(http.client.request.duration_bucket[$__rate_interval]))").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("p50"),
				).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("rate(http.client.request.duration_sum[$__rate_interval]) / rate(http.client.request.duration_count[$__rate_interval])").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("average"),
				),
		).
		WithPanel(
			timeseries.NewPanelBuilder().
				Title("Size of HTTP client response bodies.").
				Datasource(prometheusDataSourceRef).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("histogram_quantile(0.99, rate(http.client.response.body.size_bucket[$__rate_interval]))").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("p99"),
				).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("histogram_quantile(0.90, rate(http.client.response.body.size_bucket[$__rate_interval]))").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("p90"),
				).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("histogram_quantile(0.50, rate(http.client.response.body.size_bucket[$__rate_interval]))").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("p50"),
				).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("rate(http.client.response.body.size_sum[$__rate_interval]) / rate(http.client.response.body.size_count[$__rate_interval])").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("average"),
				),
		).
		WithPanel(
			timeseries.NewPanelBuilder().
				Title("Number of active HTTP server requests.").
				Datasource(prometheusDataSourceRef).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("sum(http.server.active_requests)").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("__auto"),
				),
		).
		WithPanel(
			timeseries.NewPanelBuilder().
				Title("Size of HTTP server request bodies.").
				Datasource(prometheusDataSourceRef).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("histogram_quantile(0.99, rate(http.server.request.body.size_bucket[$__rate_interval]))").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("p99"),
				).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("histogram_quantile(0.90, rate(http.server.request.body.size_bucket[$__rate_interval]))").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("p90"),
				).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("histogram_quantile(0.50, rate(http.server.request.body.size_bucket[$__rate_interval]))").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("p50"),
				).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("rate(http.server.request.body.size_sum[$__rate_interval]) / rate(http.server.request.body.size_count[$__rate_interval])").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("average"),
				),
		).
		WithPanel(
			timeseries.NewPanelBuilder().
				Title("Duration of HTTP server requests.").
				Datasource(prometheusDataSourceRef).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("histogram_quantile(0.99, rate(http.server.request.duration_bucket[$__rate_interval]))").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("p99"),
				).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("histogram_quantile(0.90, rate(http.server.request.duration_bucket[$__rate_interval]))").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("p90"),
				).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("histogram_quantile(0.50, rate(http.server.request.duration_bucket[$__rate_interval]))").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("p50"),
				).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("rate(http.server.request.duration_sum[$__rate_interval]) / rate(http.server.request.duration_count[$__rate_interval])").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("average"),
				),
		).
		WithPanel(
			timeseries.NewPanelBuilder().
				Title("Size of HTTP server response bodies.").
				Datasource(prometheusDataSourceRef).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("histogram_quantile(0.99, rate(http.server.response.body.size_bucket[$__rate_interval]))").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("p99"),
				).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("histogram_quantile(0.90, rate(http.server.response.body.size_bucket[$__rate_interval]))").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("p90"),
				).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("histogram_quantile(0.50, rate(http.server.response.body.size_bucket[$__rate_interval]))").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("p50"),
				).
				WithTarget(
					prometheus.NewDataqueryBuilder().
						Expr("rate(http.server.response.body.size_sum[$__rate_interval]) / rate(http.server.response.body.size_count[$__rate_interval])").
						Range().
						Format(prometheus.PromQueryFormatTimeSeries).
						LegendFormat("average"),
				),
		)

	// Build the dashboard - errors in configuration will be thrown here
	dashboard, err := builder.Build()
	if err != nil {
		log.Fatalf("failed to build dashboard: %v", err)
	}

	// Output the generated dashboard as JSON
	dashboardJson, err := json.MarshalIndent(dashboard, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal dashboard: %v", err)
	}
	// CALL Grafana API -X UPDATE DASHBOARD
	log.Printf(string(dashboardJson))
}
