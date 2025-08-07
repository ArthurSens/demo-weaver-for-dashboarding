package nodejs


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
  builder := dashboard.NewDashboardBuilder("nodejs").
    Uid("nodejs-dashboard").
	Tags([]string{"built-with-weaver"}).
	Refresh("5m").
	Time("now-1h", "now").
	Timezone(common.TimeZoneBrowser).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("Event loop maximum delay.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(nodejs.eventloop.delay.max[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("Event loop mean delay.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(nodejs.eventloop.delay.mean[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("Event loop minimum delay.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(nodejs.eventloop.delay.min[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("Event loop 50 percentile delay.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(nodejs.eventloop.delay.p50[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("Event loop 90 percentile delay.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(nodejs.eventloop.delay.p90[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("Event loop 99 percentile delay.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(nodejs.eventloop.delay.p99[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("Event loop standard deviation delay.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(nodejs.eventloop.delay.stddev[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("Cumulative duration of time the event loop has been in each state.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(nodejs.eventloop.time[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("Event loop utilization.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(nodejs.eventloop.utilization[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
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