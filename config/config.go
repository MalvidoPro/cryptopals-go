package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

var (
	// ProjectRoot stores the absolute path to the repository root directory.
	ProjectRoot string

	// DataDir stores the absolute path to the input data directory.
	DataDir string

	// ResultsDir stores the absolute path to the output results directory.
	ResultsDir string
)

func init() {
	// runtime.Caller(0) returns the absolute path of this source file at runtime/build time
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		panic("config: failed to determine configuration file path")
	}

	// Navigate up two directory levels: from 'internal/config/config.go' to the repository root
	configDir := filepath.Dir(currentFile) // .../cryptopals-go/config
	ProjectRoot := filepath.Dir(configDir) // .../cryptopals-go

	DataDir = filepath.Join(ProjectRoot, "testdata")
	ResultsDir = filepath.Join(ProjectRoot, "results")
}

// DataPath returns the absolute path to a file located inside the /data directory.
func DataPath(filename string) string {
	return filepath.Join(DataDir, filename)
}

// ResultPath returns the absolute path for an output file within /results.
// It automatically creates the target directory and any intermediate parent folders if they do not exist.
func ResultPath(filename string) (string, error) {
	fullPath := filepath.Join(ResultsDir, filename)
	dir := filepath.Dir(fullPath)

	// Ensure directory exists with standard read/write/execute permissions (0755)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("config: failed to create results directory %s: %w", dir, err)
	}

	return fullPath, nil
}

// WriteResult writes byte data directly to a file inside /results,
// automatically handling directory creation and file persistence.
func WriteResult(filename string, data []byte) error {
	path, err := ResultPath(filename)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
