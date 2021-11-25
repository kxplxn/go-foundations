package _010404_basicFileOperations_

import (
	"fmt"
	"log"
	"os"
)

// RenameFile renames (moves) file.
func RenameFile() {
	fmt.Println("\nRunning `RenameFile`...")

	oldPath := fmt.Sprintf("%s/file.txt", chapterDir)
	newPath := fmt.Sprintf("%s/new/new.txt", chapterDir)

	// func Rename(oldpath, newpath string) error
	if err := os.Rename(oldPath, newPath); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("`file.txt` is renamed as (or moved to) `./new/new.txt`")

	// cleanup
	if err := os.Rename(newPath, oldPath); err != nil {
		log.Fatal(err)
	}
}
