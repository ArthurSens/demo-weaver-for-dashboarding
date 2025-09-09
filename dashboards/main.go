package main

import (
	"encoding/json"
	"log/slog"
	"os"
	"path"

	"github.com/ArthurSens/demo-weaver-for-dashboarding/dashboards/index"
	_ "github.com/ArthurSens/demo-weaver-for-dashboarding/generated/dashboards/http"
	_ "github.com/ArthurSens/demo-weaver-for-dashboarding/generated/dashboards/library"
)

func main() {
	os.MkdirAll("./generated/dashboards_json/", 0755)
	for _, dashboard := range index.Dashboards {
		p := path.Join("./generated/dashboards_json/", dashboard.Metadata.Name+".json")
		f, err := os.Create(p)
		if err != nil {
			slog.Error("creating output file", "err", err, "path", p)
			os.Exit(1)
		}
		if err := json.NewEncoder(f).Encode(dashboard); err != nil {
			slog.Error("encoding dashboard to json", "err", err)
			os.Exit(1)
		}
		slog.Info("generated dashboard", "path", p)
	}
}
