package export_artifacts_utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	print "patrol_install/utils/print"
)

// closeWithLog closes a file and logs an error if closing fails.
func closeWithLog(f io.Closer, name string) {
	if err := f.Close(); err != nil {
		print.Error(fmt.Sprintf("Error closing %s: %v", name, err))
	}
}

// CopyFilesToFolder copies each file in srcFiles to destFolder, preserving the base filename.
// Returns an error if any copy fails. This is generic and can be used for APKs, IPAs, etc.
func CopyFilesToFolder(srcFiles []string, destFolder string) error {
	for _, srcFile := range srcFiles {
		dst := filepath.Join(destFolder, filepath.Base(srcFile))
		src, err := os.Open(srcFile)
		if err != nil {
			print.Error(fmt.Sprintf("Error opening %s: %v", srcFile, err))
			return err
		}

		dstFile, err := os.Create(dst)
		if err != nil {
			print.Error(fmt.Sprintf("Error creating %s: %v", dst, err))
			closeWithLog(src, srcFile)
			return err
		}

		if _, err := io.Copy(dstFile, src); err != nil {
			print.Error(fmt.Sprintf("Error copying %s to %s: %v", srcFile, dst, err))
			closeWithLog(src, srcFile)
			closeWithLog(dstFile, dst)
			return err
		}
		print.Success(fmt.Sprintf("Copied %s to %s", srcFile, dst))
		// Close files explicitly to avoid too many open files in a loop
		closeWithLog(src, srcFile)
		closeWithLog(dstFile, dst)
	}
	return nil
}
