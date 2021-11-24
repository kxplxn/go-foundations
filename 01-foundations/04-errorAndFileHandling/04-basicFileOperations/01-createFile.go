package _010404_basicFileOperations_

import (
	"fmt"
	"log"
	"os"
)

const chapterDir string = "01-foundations/04-errorAndFileHandling/04-basicFileOperations"

func CreateFile() {
	fmt.Println("\nRunning `CreateFile`...")
	// func Create(name string) (*File, error)
	f, err := os.Create(fmt.Sprintf("%s/create.txt", chapterDir))
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)
}
