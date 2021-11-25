package _010405_readingFiles_

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

//The `ReadFile` is very convenient. But, because it reads in the whole
//file, we should avoid using it on very large files.

func ReadingFilesIntoMemory() {
	fmt.Println("\nRunning `ReadingFilesIntoMemory`...")

	//func ReadFile(name string) (byte[], error)
	contents, err := os.ReadFile(fmt.Sprintf("%s/file.txt", chapterDir))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("type of `contents`:", reflect.TypeOf(contents))
	fmt.Println("value of `contents`:", contents)

	//type cast slice of uint8 values in "contents" to `string` type
	fmt.Println("`contents` to string:", string(contents))
}
