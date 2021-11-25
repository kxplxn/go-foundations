package _010404_basicFileOperations_

import (
	"fmt"
	"log"
	"os"
)

/*
	`Stat` function returns a `FileInfo` that describes a file.

	type FileInfo {
		Name() string			// base name of the file
		Size() int64			// length in bytes for regular files — system-dependent for others
		Mode() FileMode			// file mode bits
		ModTime() time.Time		// modification time
		IsDir() bool			// abbreviation for Mode().IsDir()
		Sys interface			// underlying data source (can return nil)
	}
*/

func FileInfo() {
	fmt.Println("\nRunning `FileInfo`...")
	// func Stat(name string) (FileInfo, error)
	f, err := os.Stat(fmt.Sprintf("%s/myFile.txt", chapterDir))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("file name:", f.Name())
	log.Println("file size:", f.Size())
	log.Println("file permissions:", f.Mode())
}
