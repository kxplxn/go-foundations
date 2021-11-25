package _010406_directoryOperations_

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func CreateNestedDirectories() {
	fmt.Println("\n\nRunning `CreateNestedDirectories`...")

	// ```func Join(elem ...string) string```
	// `Join` combines an arbitrary number of path elements into a single path
	// with os-specific seperators. It also handles relative paths (eg. "../test")
	p := filepath.Join(chapterDir, "test", "subdir1", "subdir2")

	err := os.MkdirAll(p, 0777)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p, "nested directory is created")
}
