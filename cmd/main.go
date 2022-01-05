package main

import (
	"github.com/rea1shane/basexporter"
	"github.com/rea1shane/basexporter/util"
	"github.com/rea1shane/datalink_exporter/pkg"
)

func main() {
	logger := util.GetLogger()
	defaultPort := 9800
	datalinkExporter := basexporter.BuildExporter(
		"datalink",
		"datalink_exporter",
		defaultPort,
		"v1.0.0")
	basexporter.Start(logger, datalinkExporter, pkg.ParseArgs(defaultPort))
}
