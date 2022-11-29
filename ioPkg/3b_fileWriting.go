package ioPkg

import (
    "bufio"
    "fmt"
    "os"
)

func writeThroughFile() {
	// For more granular writes, open a file for writing,
	var testFilePath string = "./resources/file-writing-test2.txt"
	f, err := os.Create(testFilePath)
    check(err)
	defer f.Close()
	
	// and write byte-slices to the file
	d2 := []byte{115, 111, 109, 101, 10} // ascii for "some\n"
    n2, err := f.Write(d2)
    check(err)
	fmt.Printf("wrote %d bytes\n", n2)
	
	// A WriteString is also available.
	n3, err := f.WriteString("writes\n")
    check(err)
	fmt.Printf("wrote %d bytes\n", n3)
	
	// Issue a Sync to flush writes to stable storage.
	f.Sync()

	// bufio provides buffered writers
	w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n4)

	// Use Flush to ensure all buffered operations have been applied to the underlying writer.
    w.Flush()
}

func RunFileWrite() {
	var testFilePath string = "./resources/file-writing-test.txt"

	// Dump a string (or just bytes) into a file.
	d1 := []byte("hello\ngo\n")
    err := os.WriteFile(testFilePath, d1, 0644)
    check(err)
	
	writeThroughFile()
}