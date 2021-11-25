package _010404_basicFileOperations_

import (
	"fmt"
	"log"
	"os"
)

func RemoveFile() {
	fmt.Println("\nRunning `RemoveFile`...")
	err := os.Remove(fmt.Sprintf("%s/del.txt", chapterDir))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file successfuly removed")
}
