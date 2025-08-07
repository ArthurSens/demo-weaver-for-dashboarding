package dotnet


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
  builder := dashboard.NewDashboardBuilder("dotnet").
    Uid("dotnet-dashboard").
	Tags([]string{"built-with-weaver"}).
	Refresh("5m").
	Time("now-1h", "now").
	Timezone(common.TimeZoneBrowser).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The number of .NET assemblies that are currently loaded.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(dotnet.assembly.count[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The number of exceptions that have been thrown in managed code.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(dotnet.exceptions[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The number of garbage collections that have occurred since the process has started.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(dotnet.gc.collections[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The *approximate* number of bytes allocated on the managed GC heap since the process has started. The returned value does not include any native allocations.
").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(dotnet.gc.heap.total_allocated[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The heap fragmentation, as observed during the latest garbage collection.
").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(dotnet.gc.last_collection.heap.fragmentation.size[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The managed GC heap size (including fragmentation), as observed during the latest garbage collection.
").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(dotnet.gc.last_collection.heap.size[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The amount of committed virtual memory in use by the .NET GC, as observed during the latest garbage collection.
").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(dotnet.gc.last_collection.memory.committed_size[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The total amount of time paused in GC since the process has started.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(dotnet.gc.pause.time[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The amount of time the JIT compiler has spent compiling methods since the process has started.
").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(dotnet.jit.compilation.time[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("Count of bytes of intermediate language that have been compiled since the process has started.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(dotnet.jit.compiled_il.size[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The number of times the JIT compiler (re)compiled methods since the process has started.
").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(dotnet.jit.compiled_methods[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The number of times there was contention when trying to acquire a monitor lock since the process has started.
").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(dotnet.monitor.lock_contentions[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The number of processors available to the process.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(dotnet.process.cpu.count[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("CPU time used by the process.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(dotnet.process.cpu.time[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The number of bytes of physical memory mapped to the process context.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(dotnet.process.memory.working_set[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The number of work items that are currently queued to be processed by the thread pool.
").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(dotnet.thread_pool.queue.length[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The number of thread pool threads that currently exist.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(dotnet.thread_pool.thread.count[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The number of work items that the thread pool has completed since the process has started.
").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(dotnet.thread_pool.work_item.count[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The number of timer instances that are currently active.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(dotnet.timer.count[$__rate_interval])").
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