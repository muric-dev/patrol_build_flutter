package export_artifacts_utils

import (
	"fmt"
	"os"
)

// CreateFolder ensures the given directory exists.
// Returns nil if the folder exists or is created, or an error if creation fails.
func CreateFolder(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("failed to create folder %s: %w", path, err)
	}
	return nil
}
