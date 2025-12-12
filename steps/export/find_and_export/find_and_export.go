package find_and_export

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"patrol_install/commands"
	regex "patrol_install/constants"
	build_constants "patrol_install/steps/build/constants"
	print "patrol_install/utils/print"
)

var androidApkPath = "build/app/outputs/apk"
var iosAppPath = "build/ios_integ/Build/Products"

func FindAndExportBuilds() error {
	platform := os.Getenv(build_constants.Platform)
	err := findAndMoveAndroidBuilds(platform)
	if err != nil {
		return err
	}
	err = findAndMoveIOSBuilds(platform)
	if err != nil {
		return err
	}
	return nil
}

func findAndMoveAndroidBuilds(platform string) error {
	if platform == "ios" {
		print.Action("No Android builds were selected to build")
		return nil
	}

	apkFiles, err := findAndroidApks(androidApkPath)
	if err != nil {
		return err
	}

	if len(apkFiles) == 0 {
		print.Error("No Android APK files found.")
		return nil
	}

	// Create patrol directory if it doesn't exist
	if _, err := os.Stat("patrol"); os.IsNotExist(err) {
		command := commands.CreatePatrolFolder
		_, err := exec.Command(command.Name, command.Args...).Output()
		if err != nil {
			return err
		}
	}

	// Move the found APK files to the patrol directory
	for _, apkFile := range apkFiles {
		commands := commands.CopyBuildsToFolder.CopyWith(nil, []string{apkFile, "patrol"})

		_, err := exec.Command(commands.Name, commands.Args...).Output()
		if err != nil {
			print.Error(fmt.Sprintf("Error moving %s: %v\n", apkFile, err))
			return err
		}
		print.Success(fmt.Sprintf("Moved %s to patrol/\n", apkFile))
	}

	return nil
}

func findAndroidApks(root string) ([]string, error) {
	var files []string
	regex := regex.AndroidApk()

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if regex.MatchString(info.Name()) {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		print.Error(fmt.Sprintf("error walking the path %s: %v", root, err))
		return nil, err
	}

	return files, nil
}

func findAndMoveIOSBuilds(platform string) error {
	if platform == "android" {
		print.Action("No iOS builds were selected to build")
		return nil
	}

	appPattern := filepath.Join(iosAppPath, "Release-iphoneos", "*.app")
	testrunPattern := filepath.Join(iosAppPath, "*.xctestrun")

	// Find .app files
	appFiles, err := filepath.Glob(appPattern)
	if err != nil {
		print.Error(fmt.Sprintf("error finding .app files with pattern %s: %v", appPattern, err))
		return err
	}

	// Find .xctestrun files
	testrunFiles, err := filepath.Glob(testrunPattern)
	if err != nil {
		print.Error(fmt.Sprintf("error finding .xctestrun files with pattern %s: %v", testrunPattern, err))
		return err
	}

	if len(appFiles) == 0 {
		err := errors.New("build failed: no .app file found to zip")
		print.Error(err.Error())
		return err
	}
	if len(testrunFiles) == 0 {
		err := errors.New("build failed: no .xctestrun file found to zip")
		print.Error(err.Error())
		return err
	}

	sourceFiles := append(appFiles, testrunFiles...)

	command := commands.CompressIOSFiles
	command.Args = append(command.Args, sourceFiles...)

	print.Action(fmt.Sprintf("Zip command: %s %s\n", command.Name, strings.Join(command.Args, " ")))

	output, err := exec.Command(command.Name, command.Args...).CombinedOutput()

	if err != nil {
		print.Action(string(output))
		print.Error(fmt.Sprintf("Error when creating ios_tests.zip: %v", err))
		return err
	}

	print.StepCompleted("iOS builds zipped successfully.")

	return nil
}
