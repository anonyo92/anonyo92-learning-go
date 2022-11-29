package iopkg

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var testFilePath string = "./resources/file-reading-test.txt"

func readThroughFile() {
	// Start by Opening a file to obtain an os.File value.
	f, err := os.Open(testFilePath) // returns os.File
	check(err)
	defer f.Close()

	/* Read some bytes from the beginning of the file. */
	// Allow up to 5 to be read but also note how many actually were read.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// We can also seek to a known location in the file and read from there.
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 4)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %v\n", n2, o2, string(b2[:n2]))

	// reads like the ones above can be more robustly implemented with io.ReadAtLeast()
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 4)
	n3, err := io.ReadAtLeast(f, b3, 3)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// There is no built-in rewind, but Seek(0, 0) accomplishes this
	_, err = f.Seek(0, 0)
	check(err)

	/* The bufio package implements a buffered reader that may be useful
	 * both for its efficiency with many small reads
	 * and because of the additional reading methods it provides. */
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
}

func RunFileRead() {
	// slurping a file’s entire contents into memory.
	dat, err := os.ReadFile(testFilePath)
	check(err)
	fmt.Print(string(dat))

	/* We often want more control over how and what parts of a file are read.
	 * For this, we might want to read through the file, instead all contents in one go.*/
	readThroughFile()
}
