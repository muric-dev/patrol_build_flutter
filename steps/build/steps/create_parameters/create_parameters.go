package create_parameters

import (
	"os"

	constants "patrol_install/steps/build/constants"
	bp "patrol_install/steps/build/models/build_parameters"
)

func BuildParametersFromEnv() (*bp.BuildParameters, error) {
	envMap := map[string]string{
		"platform":     os.Getenv(constants.Platform),
		"target":       os.Getenv(constants.Target),
		"buildType":    os.Getenv(constants.BuildType),
		"tags":         os.Getenv(constants.Tags),
		"excludedTags": os.Getenv(constants.ExcludedTags),
		"verbose":      os.Getenv(constants.IsVerbose),
		"isCovered":    os.Getenv(constants.IsCovered),
		"filePath":     os.Getenv(constants.FilePath),
	}

	// Final build
	return bp.NewBuildParameters(envMap)
}
