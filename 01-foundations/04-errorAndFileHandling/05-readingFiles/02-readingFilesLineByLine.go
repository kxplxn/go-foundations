package _010405_readingFiles_

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadingFilesLineByLine() {
	fmt.Println("\nRunning `ReadingFilesLineByLine`...")

	f, err := os.Open(fmt.Sprintf("%s/file.txt", chapterDir))
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Package `bufio` implements buffered IO and it uses Reader/Scanner types.
	// These are important types in bufio...
	// 1. Reader (eg. f) implements buffering for a Reader object.
	// 2. Scanner implements an interface for reading data such as lines of text
	//    in a file with line breaks.

	// func NewScanner(r io.Reader) *Scanner
	s := bufio.NewScanner(f)

	// func (s *Scanner) Scan() bool
	// 1. Scan can be called to advance the Scanner token by token and make it
	// 	   available via a Text method.
	// 2. Scan returns false when scanning stops at the end of input or if an
	//    error occurs.
	// 3. Since Scan will return true as long as there's data to read, we can
	//    use it in a for loop like so:
	for s.Scan() {
		fmt.Println(s.Text())
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
