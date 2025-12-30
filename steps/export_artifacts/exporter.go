package export_artifacts

import "patrol_install/utils/print"

type Exporter interface {
	FindAndExport() error
}

func Run(exporter Exporter) error {
	print.StepInitiated("--- Getting Patrol builds ---")
	return exporter.FindAndExport()
}
