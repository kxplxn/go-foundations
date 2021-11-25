package _010404_basicFileOperations_

import (
	"fmt"
	"log"
	"os"
)

func cleanup() {
	_, err := os.Create(fmt.Sprintf("%s/del.txt", chapterDir))
	if err != nil {
		log.Fatal(err)
	}
}

func RemoveFile() {
	defer cleanup()

	fmt.Println("\nRunning `RemoveFile`...")
	err := os.Remove(fmt.Sprintf("%s/del.txt", chapterDir))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file successfuly removed")
}
