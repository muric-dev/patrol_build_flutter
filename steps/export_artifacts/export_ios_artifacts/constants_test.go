package export_ios_artifacts

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

func TestIOSOutputKeysMatchStepYml(t *testing.T) {
	// GIVEN step.yml contents
	stepYmlPath := filepath.Join("..", "..", "..", "step.yml")
	contents, err := os.ReadFile(stepYmlPath)
	if err != nil {
		t.Fatalf("read step.yml: %v", err)
	}

	// WHEN we search for the iOS output keys
	outputKeys := []string{
		IOSAppUnderTestPathEnvKey,
		IOSTestInstrumentationEnvKey,
		IOSRunnerFilePathEnvKey,
		IOSBuildExportsZipPathEnvKey,
	}

	// THEN each key exists in step.yml outputs
	for _, key := range outputKeys {
		pattern := fmt.Sprintf(`(?m)^\s*-\s+%s:`, regexp.QuoteMeta(key))
		if !regexp.MustCompile(pattern).Match(contents) {
			t.Fatalf("expected output key %s in step.yml", key)
		}
	}
}
