package export_android_artifacts

import (
	"fmt"
	"os"
	"path/filepath"
	regex "patrol_install/constants"
	build_constants "patrol_install/steps/build/constants"
	constants "patrol_install/steps/export_artifacts/constants"
	export_artifacts_utils "patrol_install/steps/export_artifacts/utils"
	print "patrol_install/utils/print"
)

// CopyAndroidArtifacts finds the first test and app APKs and copies them to the artifacts directory.
func CopyAndroidArtifacts(artifactsPath, testPath, appPath string) error {
	platform := os.Getenv(build_constants.Platform)
	if !IsAndroidPlatform(platform) {
		print.Action("No Android builds were selected to build")
		return nil
	}

	apkFiles := make([]string, 0, 2)
	if testApk, err := FindFirstApkInDir(testPath); err != nil {
		return err
	} else if testApk != "" {
		apkFiles = append(apkFiles, testApk)
	}
	if appApk, err := FindFirstApkInDir(appPath); err != nil {
		return err
	} else if appApk != "" {
		apkFiles = append(apkFiles, appApk)
	}

	if len(apkFiles) == 0 {
		print.Error("No Android/Test APK files found.")
		return nil
	}

	if err := export_artifacts_utils.CreateFolder(artifactsPath); err != nil {
		return err
	}

	if err := export_artifacts_utils.CopyFilesToFolder(apkFiles, artifactsPath); err != nil {
		return err
	}

	return nil
}

// IsAndroidPlatform returns true if the platform is Android.
func IsAndroidPlatform(platform string) bool {
	return platform == "android" || platform == "both"
}

// AndroidApkPaths returns the test and app APK search paths for the given build type.
func AndroidApkPaths(isRelease bool) (testPath, appPath string) {
	if isRelease {
		return constants.AndroidTestPath + constants.ReleaseFolder, constants.AndroidAppPath + constants.ReleaseFolder
	}
	return constants.AndroidTestPath + constants.DebugFolder, constants.AndroidAppPath + constants.DebugFolder
}

// FindFirstApkInDir returns the first APK file found in the given directory, or an empty string if none found.
func FindFirstApkInDir(root string) (string, error) {
	var apkPath string
	rgx := regex.AndroidApk()
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if rgx.MatchString(info.Name()) {
			apkPath = path
			return filepath.SkipAll
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("error walking the path %s: %w", root, err)
	}
	return apkPath, nil
}
