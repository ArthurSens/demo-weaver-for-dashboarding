package vcs


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
  builder := dashboard.NewDashboardBuilder("vcs").
    Uid("vcs-dashboard").
	Tags([]string{"built-with-weaver"}).
	Refresh("5m").
	Time("now-1h", "now").
	Timezone(common.TimeZoneBrowser).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The number of changes (pull requests/merge requests/changelists) in a repository, categorized by their state (e.g. open or merged).").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(vcs.change.count[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The time duration a change (pull request/merge request/changelist) has been in a given state.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(vcs.change.duration[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The amount of time since its creation it took a change (pull request/merge request/changelist) to get the first approval.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(vcs.change.time_to_approval[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The amount of time since its creation it took a change (pull request/merge request/changelist) to get merged into the target(base) ref.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(vcs.change.time_to_merge[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The number of unique contributors to a repository.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(vcs.contributor.count[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The number of refs of type branch or tag in a repository.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(vcs.ref.count[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The number of lines added/removed in a ref (branch) relative to the ref from the `vcs.ref.base.name` attribute.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(vcs.ref.lines_delta[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The number of revisions (commits) a ref (branch) is ahead/behind the branch from the `vcs.ref.base.name` attribute.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(vcs.ref.revisions_delta[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("Time a ref (branch) created from the default branch (trunk) has existed. The `ref.type` attribute will always be `branch`.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(vcs.ref.time[$__rate_interval])").
					Range().
					Format(prometheus.PromQueryFormatTimeSeries).
					LegendFormat("__auto"),
			),
	).
	WithPanel(
		timeseries.NewPanelBuilder().
			Title("The number of repositories in an organization.").
			Datasource(prometheusDataSourceRef).
			WithTarget(
				prometheus.NewDataqueryBuilder().
					Expr("rate(vcs.repository.count[$__rate_interval])").
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