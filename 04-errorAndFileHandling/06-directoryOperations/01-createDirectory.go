package _010406_directoryOperations_

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func CreateDirectory() {
	fmt.Println("\nRunning `CreateDirectory`...")
	subdir := fmt.Sprintf("%s/subdir", chapterDir)

	d, err := os.Stat(subdir)
	fmt.Printf("error returned from `os.Stat()` is `%s`.\n", err)

	if err == nil {
		fmt.Println(d.Mode(), d.IsDir())
		log.Fatal("file/directory name already exists")
	}

	if errors.Is(err, os.ErrNotExist) {
		// MkdirAll() will do nothing and return nil if the directory already
		// exists. That's why we need to do the checks above.
		err := os.Mkdir(subdir, 0777)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(subdir, "is created")
	}
}
