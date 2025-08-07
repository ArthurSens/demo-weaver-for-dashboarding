module github.com/ArthurSens/demo-weaver-for-dashboarding/dashboards/gen_ai

go 1.24.2

require (
   github.com/grafana/grafana-foundation-sdk/go v0.0.0-20250310114924-e8eb8530bc7c
   github.com/prometheus/client_golang v1.23.0
   github.com/ArthurSens/demo-weaver-for-dashboarding/dashboards/server v0.0.0-00010101000000-000000000000
   github.com/ArthurSens/demo-weaver-for-dashboarding/dashboards/server v0.0.0-00010101000000-000000000000
)
replace github.com/ArthurSens/demo-weaver-for-dashboarding/dashboards/server => ../server
replace github.com/ArthurSens/demo-weaver-for-dashboarding/dashboards/server => ../server