package export_ios_artifacts

const (
	IOSArtifactsPath             = "patrol/ios"
	IOSAppUnderTestPathEnvKey    = "IOS_APP_UNDER_TEST"
	IOSTestInstrumentationEnvKey = "IOS_TEST_INSTRUMENTATION_APP"
	IOSRunnerFilePathEnvKey      = "IOS_RUNNER_FILE"
	IOSBuildExportsZipPathEnvKey = "IOS_BUILD_EXPORTS"

	IOSBuildProductsPath    = "build/ios_integ/Build/Products"
	IOSReleaseBuildDirName  = "Release-iphoneos"
	IOSDebugBuildDirName    = "Debug-iphonesimulator"
	IOSAppUnderTestName     = "Runner.app"
	IOSTestInstrumentation  = "RunnerUITests-Runner.app"
	IOSXCTestRunGlobPattern = "*.xctestrun"
	IOSExportsZipName       = "ios_tests.zip"
)
