package export

import (
	exporter "patrol_install/steps/export/find_and_export"
)

type ExporterRunner struct{}

func (p *ExporterRunner) FindAndExportBuilds() error {
	return exporter.FindAndExportBuilds()
}
