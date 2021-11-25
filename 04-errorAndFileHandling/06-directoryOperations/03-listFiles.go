package _010406_directoryOperations_

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func ListFiles() {
	fmt.Println("\nRunning `ListFiles`...")

	// func ReadDir(name string) ([]DirEntry, error)
	ls, err := os.ReadDir(filepath.Join(chapterDir, "test"))
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range ls {
		fmt.Println(f.Name(), f.IsDir())
	}
}
