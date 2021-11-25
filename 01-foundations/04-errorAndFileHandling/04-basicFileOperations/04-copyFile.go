package _010404_basicFileOperations_

import (
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CopyFile() {
	fmt.Println("\nRunning `CopyFile`...")

	src, err := os.Open(fmt.Sprintf("%s/src.txt", chapterDir))
	handleError(err)

	// func OpenFile(name string, flag int, perm FileMode) (*File, error)
	dst, err := os.OpenFile(
		fmt.Sprintf("%s/dst.txt", chapterDir),
		os.O_RDWR|os.O_CREATE,
		0755,
	)
	handleError(err)

	// func Copy(dst Writer, src Reader) (written int64, err error)
	w, err := io.Copy(dst, src)
	handleError(err)

	fmt.Println(reflect.TypeOf(w))
	fmt.Println(w)
}
