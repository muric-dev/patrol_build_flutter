package main

import (
	build "patrol_install/steps/build"
	export_artifacts "patrol_install/steps/export_artifacts"
	install_patrol_cli "patrol_install/steps/install_patrol_cli"
	validate "patrol_install/steps/validate"
	print "patrol_install/utils/print"
)

func main() {
	cliVersion, installError := install_patrol_cli.Run(&install_patrol_cli.InstallerRunner{})
	if installError != nil {
		print.Error("❌ Setup failed")
		print.Error(installError.Error())
		print.Error("Please check the logs for more details.")
	} else {
		print.Success("✅ Installing CLI Completed Successfully")
	}

	validatorParams := validate.ValidatorRunParams{
		Runner:     &validate.ValidatorRunner{},
		CliVersion: cliVersion,
	}

	validationError := validate.Run(validatorParams)
	if validationError != nil {
		print.Error("❌ Validation failed")
		print.Error(validationError.Error())
		print.Error("Please check the logs for more details.")
		return
	}

	buildError := build.Run(&build.BuilderRunner{})
	if buildError != nil {
		print.Error("❌ Build failed")
		print.Error(buildError.Error())
		print.Error("Please check the logs for more details.")
		return
	}

	exportError := export_artifacts.Run(&export_artifacts.ExporterRunner{})
	if exportError != nil {
		print.Error("❌ Export failed")
		print.Error(exportError.Error())
		print.Error("Please check the logs for more details.")
		return
	}

}
