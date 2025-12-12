package install_patrol_cli

import (
	v "github.com/Masterminds/semver/v3"

	"patrol_install/utils/print"
)

type Installer interface {
	GetPatrolCLIVersion() (*v.Version, error)
	InstallPatrolCLI() error
}

func Run(installer Installer) (*v.Version, error) {
	print.StepInitiated("--- Checking if Patrol CLI is already installed ---")

	version, err := installer.GetPatrolCLIVersion()
	if err != nil {
		print.Warning("CLI is not installed, attempting installation...")
		if err := installer.InstallPatrolCLI(); err != nil {
			print.Error("❌ Installation failed: " + err.Error())
			return nil, err
		}

		version, err = installer.GetPatrolCLIVersion()
		if err != nil {
			print.Error("❌ Failed to verify version after install: " + err.Error())
			return nil, err
		}

		print.StepCompleted("✅ PATROL CLI installed successfully. Version: " + version.String() + "\n")
		return version, nil
	}

	print.StepCompleted("✅ Tool already installed. Version: " + version.String() + "\n")
	return version, nil
}
