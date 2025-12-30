package export_artifacts_utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFilesToFolder(t *testing.T) {
	srcDir := t.TempDir()
	dstDir := t.TempDir()
	file1 := filepath.Join(srcDir, "a.txt")
	file2 := filepath.Join(srcDir, "b.txt")
	if err := os.WriteFile(file1, []byte("foo"), 0644); err != nil {
		t.Fatalf("failed to write file1: %v", err)
	}
	if err := os.WriteFile(file2, []byte("bar"), 0644); err != nil {
		t.Fatalf("failed to write file2: %v", err)
	}
	files := []string{file1, file2}
	envKeys := []string{"TEST_ENV_A", "TEST_ENV_B"}
	if err := CopyFilesToFolder(files, dstDir, envKeys); err != nil {
		t.Fatalf("CopyFilesToFolder failed: %v", err)
	}
	for i, f := range files {
		base := filepath.Base(f)
		if _, err := os.Stat(filepath.Join(dstDir, base)); err != nil {
			t.Errorf("file %s not copied", base)
		}
		// Check env variable is set
		val := os.Getenv(envKeys[i])
		if val == "" {
			t.Errorf("env %s not set", envKeys[i])
		}
	}
}

func TestCopyFilesToFolder_Error(t *testing.T) {
	dstDir := t.TempDir()
	files := []string{"/nonexistent/file.txt"}
	envKeys := []string{"DUMMY_ENV"}
	if err := CopyFilesToFolder(files, dstDir, envKeys); err == nil {
		t.Error("expected error for missing file")
	}
}

func TestCopyFilesToFolder_LengthMismatch(t *testing.T) {
	dstDir := t.TempDir()
	files := []string{"/tmp/a.txt", "/tmp/b.txt"}
	envKeys := []string{"ENV_A"} // Mismatch
	if err := CopyFilesToFolder(files, dstDir, envKeys); err == nil {
		t.Error("expected error for length mismatch")
	}
}
