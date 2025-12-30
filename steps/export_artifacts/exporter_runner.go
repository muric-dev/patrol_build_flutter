package export_artifacts

import (
	"os"
	build_constants "patrol_install/steps/build/constants"
	export_android_artifacts "patrol_install/steps/export_artifacts/export_android_artifacts"
)

type ExporterRunner struct{}

func (p *ExporterRunner) FindAndExportAndroid() error {
	isRelease := os.Getenv(build_constants.BuildType) == "release"
	artifactsPath := export_android_artifacts.AndroidArtifactsPath
	testPath, appPath := export_android_artifacts.AndroidApkPaths(isRelease)
	return export_android_artifacts.CopyAndroidArtifacts(artifactsPath, testPath, appPath)
}
