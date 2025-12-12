package export_artifacts_utils

import (
	"os"
	"testing"
)

func TestCreateFolder_New(t *testing.T) {
	dir := t.TempDir()
	folder := dir + "/newfolder"
	if err := CreateFolder(folder); err != nil {
		t.Fatalf("CreateFolder failed: %v", err)
	}
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		t.Errorf("folder was not created")
	}
}

func TestCreateFolder_Existing(t *testing.T) {
	dir := t.TempDir()
	if err := CreateFolder(dir); err != nil {
		t.Fatalf("CreateFolder failed on existing: %v", err)
	}
}
