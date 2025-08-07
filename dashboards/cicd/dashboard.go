package cicd


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
  builder := dashboard.NewDashboardBuilder("cicd").
    Uid("cicd-dashboard").
	Tags([]string{"built-with-weaver"}).
	Refresh("5m").
	Time("now-1h", "now").
	Timezone(common.TimeZoneBrowser).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The number of pipeline runs currently active in the system by state.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(cicd.pipeline.run.active[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("Duration of a pipeline run grouped by pipeline, state and result.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(cicd.pipeline.run.duration[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The number of errors encountered in pipeline runs (eg. compile, test failures).").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(cicd.pipeline.run.errors[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The number of errors in a component of the CICD system (eg. controller, scheduler, agent).").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(cicd.system.errors[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The number of workers on the CICD system by state.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(cicd.worker.count[$__rate_interval])").
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