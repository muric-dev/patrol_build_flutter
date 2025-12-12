package export_artifacts

import (
	"patrol_install/utils/print"
)

type Exporter interface {
	FindAndExportAndroid() error
}

func Run(exporter Exporter) error {
	print.StepInitiated("--- Getting Patrol builds ---")

	if err := exporter.FindAndExportAndroid(); err != nil {
		return err
	}

	return nil
}
