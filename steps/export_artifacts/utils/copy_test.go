package export_artifacts_utils

import (
	"os"
	"path/filepath"
	"testing"
)

type stubEnvExporter struct {
	t        *testing.T
	exported map[string]string
}

func (s *stubEnvExporter) Export(key, value string) error {
	if s.exported == nil {
		s.exported = make(map[string]string)
	}
	s.exported[key] = value
	s.t.Setenv(key, value)
	return nil
}

func setupEnvExporterStub(t *testing.T) *stubEnvExporter {
	stub := &stubEnvExporter{
		t:        t,
		exported: make(map[string]string),
	}
	SetEnvExporter(stub)
	t.Cleanup(func() {
		SetEnvExporter(nil)
	})
	return stub
}

func TestCopyFilesToFolder(t *testing.T) {
	stub := setupEnvExporterStub(t)
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
		dstPath := filepath.Join(dstDir, base)
		if _, err := os.Stat(dstPath); err != nil {
			t.Errorf("file %s not copied", base)
		}
		// Check env variable is set
		val := os.Getenv(envKeys[i])
		if val != dstPath {
			t.Errorf("env %s expected %s, got %s", envKeys[i], dstPath, val)
		}
		if exported, ok := stub.exported[envKeys[i]]; !ok || exported != dstPath {
			t.Errorf("exporter expected %s=%s, got %s", envKeys[i], dstPath, exported)
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

func TestCopyFilesToFolder_Directory(t *testing.T) {
	// GIVEN a source directory with a file
	stub := setupEnvExporterStub(t)
	srcDir := t.TempDir()
	nestedFile := filepath.Join(srcDir, "nested.txt")
	if err := os.WriteFile(nestedFile, []byte("data"), 0644); err != nil {
		t.Fatalf("failed to write nested file: %v", err)
	}
	dstDir := t.TempDir()

	// WHEN copying the directory
	envKey := "TEST_DIR_ENV"
	if err := CopyFilesToFolder([]string{srcDir}, dstDir, []string{envKey}); err != nil {
		t.Fatalf("CopyFilesToFolder failed: %v", err)
	}

	// THEN the directory and contents are copied and env exported
	copiedDir := filepath.Join(dstDir, filepath.Base(srcDir))
	if _, err := os.Stat(copiedDir); err != nil {
		t.Fatalf("expected directory copied: %v", err)
	}
	if _, err := os.Stat(filepath.Join(copiedDir, "nested.txt")); err != nil {
		t.Fatalf("expected nested file copied: %v", err)
	}
	if val := os.Getenv(envKey); val != copiedDir {
		t.Fatalf("expected env %s=%s, got %s", envKey, copiedDir, val)
	}
	if exported, ok := stub.exported[envKey]; !ok || exported != copiedDir {
		t.Fatalf("exporter expected %s=%s, got %s", envKey, copiedDir, exported)
	}
}

func TestCopyFile(t *testing.T) {
	// GIVEN a source file
	srcDir := t.TempDir()
	srcPath := filepath.Join(srcDir, "source.txt")
	if err := os.WriteFile(srcPath, []byte("payload"), 0644); err != nil {
		t.Fatalf("failed to write source file: %v", err)
	}
	dstDir := t.TempDir()
	dstPath := filepath.Join(dstDir, "source.txt")

	// WHEN copying the file
	if err := copyFile(srcPath, dstPath, 0644); err != nil {
		t.Fatalf("copyFile failed: %v", err)
	}

	// THEN contents are preserved
	data, err := os.ReadFile(dstPath)
	if err != nil {
		t.Fatalf("failed to read dest file: %v", err)
	}
	if string(data) != "payload" {
		t.Fatalf("expected payload, got %s", string(data))
	}
}

func TestCopyDir(t *testing.T) {
	// GIVEN a directory tree
	srcDir := t.TempDir()
	nestedDir := filepath.Join(srcDir, "nested")
	if err := os.MkdirAll(nestedDir, 0755); err != nil {
		t.Fatalf("failed to create nested dir: %v", err)
	}
	nestedFile := filepath.Join(nestedDir, "file.txt")
	if err := os.WriteFile(nestedFile, []byte("content"), 0644); err != nil {
		t.Fatalf("failed to write nested file: %v", err)
	}
	dstDir := t.TempDir()

	// WHEN copying the directory
	if err := copyDir(srcDir, filepath.Join(dstDir, "copied")); err != nil {
		t.Fatalf("copyDir failed: %v", err)
	}

	// THEN the nested file exists in the destination
	copiedFile := filepath.Join(dstDir, "copied", "nested", "file.txt")
	if _, err := os.Stat(copiedFile); err != nil {
		t.Fatalf("expected copied file, got error: %v", err)
	}
}
