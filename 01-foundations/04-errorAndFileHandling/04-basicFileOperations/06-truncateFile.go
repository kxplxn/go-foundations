package _010404_basicFileOperations_

import (
	"fmt"
	"log"
	"os"
)

func TruncateFile() {
	fmt.Println("\n\nRunning `TruncateFile`...")

	// Truncate myFile.txt to the size of 10 bytes.
	err := os.Truncate(fmt.Sprintf("%s/myFile.txt", chapterDir), 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("`myFile.txt` truncated to 10 bytes")
}
