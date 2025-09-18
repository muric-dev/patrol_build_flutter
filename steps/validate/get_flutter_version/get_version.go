package get_flutter_version

import (
	"fmt"
	"strings"

	v "github.com/Masterminds/semver/v3"

	"patrol_install/commands"
	regex "patrol_install/constants"
	"patrol_install/utils/exec"
)

var flutterVersion = commands.FlutterVersion

func GetVersion() (*v.Version, error) {
	output, err := exec.Command(flutterVersion)
	if err != nil {
		return nil, err
	}

	// Use regex to extract the version
	re := regex.Version("Flutter")
	match := re.FindStringSubmatch(output)
	if len(match) > 1 {
		parsedVersion, err := v.NewVersion(cleanVersion(match[1]))
		if err != nil {
			return nil, fmt.Errorf("invalid semantic version: %w", err)
		}
		return parsedVersion, nil
	}

	return nil, fmt.Errorf("could not find version in output")
}

func cleanVersion(version string) string {
	version = strings.TrimSpace(version)
	version = strings.TrimPrefix(version, "v")
	return version
}
