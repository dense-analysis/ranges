package main

import (
	"os"
	"testing"
)

func TestCacheB(t *testing.T) {
	t.Parallel()

	dirName, err := os.MkdirTemp("", "")

	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dirName)

	// Run the code to generate files and ensure it doesn't blow up.
	createFiles(filesToCreate{
		tupleFilename: dirName + "/tuple.go",
		zipFilename:   dirName + "/zip.go",
	})
}
