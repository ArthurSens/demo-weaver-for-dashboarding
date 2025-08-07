package process


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
  builder := dashboard.NewDashboardBuilder("process").
    Uid("process-dashboard").
	Tags([]string{"built-with-weaver"}).
	Refresh("5m").
	Time("now-1h", "now").
	Timezone(common.TimeZoneBrowser).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("Number of times the process has been context switched.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(process.context_switches[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("Total CPU seconds broken down by different states.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(process.cpu.time[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("Difference in process.cpu.time since the last measurement, divided by the elapsed time and number of CPUs available to the process.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(process.cpu.utilization[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("Disk bytes transferred.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(process.disk.io[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The amount of physical memory in use.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(process.memory.usage[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The amount of committed virtual memory.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(process.memory.virtual[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("Network bytes transferred.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(process.network.io[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("Number of file descriptors in use by the process.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(process.open_file_descriptor.count[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("Number of page faults the process has made.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(process.paging.faults[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("Process threads count.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(process.thread.count[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The time the process has been running.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(process.uptime[$__rate_interval])").
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