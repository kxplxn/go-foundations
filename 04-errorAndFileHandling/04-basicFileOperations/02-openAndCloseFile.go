package _010404_basicFileOperations_

import (
	"fmt"
	"log"
	"os"
)

func closer(f *os.File) error {
	// func (f *File) Close() error
	f.Close()
	fmt.Println(f.Name(), "successfuly closed")
	return nil
}

func OpenAndCloseFile() {
	fmt.Println("\nRunning `OpenAndCloseFile`...")
	// func Open(name string) (*File, error)
	f, err := os.Open(fmt.Sprintf("%s/file.txt", chapterDir))
	defer closer(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f.Name(), "successfuly opened")
}
