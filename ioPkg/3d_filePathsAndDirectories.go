package ioPkg

/* The filepath package provides functions to parse and construct file paths in a way
 * that is portable between operating systems; dir/file on Linux vs. dir\file on Windows, for example. */

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func RunFilePaths() {
	/* Join should be used to construct paths in a portable way. 
	 * It takes any number of arguments and constructs a hierarchical path from them.*/
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	/* In addition to providing portability, 
	 * Join will also normalize paths by removing superfluous separators and directory changes. */
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))
	
	/* Dir and Base can be used to split a path to the directory and the file. 
	 * Alternatively, Split will return both in the same call. */
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))
	
	// We can check whether a path is absolute.
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))
	
	// Some file names have extensions following a dot. We can split the extension out of such names with Ext.
	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext)
	
	// To find the file’s name with the extension removed, use strings.TrimSuffix.
	fmt.Println(strings.TrimSuffix(filename, ext))

	/* Rel finds a relative path between a base and a target. 
	 * It returns an error if the target cannot be made relative to base. */
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}

/* Go has several useful functions for working with directories in the file system. */
func RunDirectories() {
	// Create a new sub-directory in the current working directory.
	err := os.Mkdir("subdir", 0755)
	check(err)

	/* When creating temporary directories, it’s good practice to defer their removal. 
	 * os.RemoveAll will delete a whole directory tree (similarly to rm -rf). */
	defer os.RemoveAll("subdir")
	
	// Helper function to create a new empty file.
	createEmptyFile := func(name string) {
        d := []byte("")
        check(os.WriteFile(name, d, 0644))
	}
	
	createEmptyFile("subdir/file1")

	/* We can create a hierarchy of directories, including parents with MkdirAll. 
	 * This is similar to the command-line mkdir -p. */
    err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)
	
	createEmptyFile("subdir/parent/file2")
    createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")
	
	// ReadDir lists directory contents, returning a slice of os.DirEntry objects.
    c, err := os.ReadDir("subdir/parent")
    check(err)
    fmt.Println("Listing subdir/parent")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
	}
	
	// Chdir lets us change the current working directory, similarly to cd.
    err = os.Chdir("subdir/parent/child")
	check(err)
	
	// Now we’ll see the contents of subdir/parent/child when listing the current directory.
    c, err = os.ReadDir(".")
    check(err)
    fmt.Println("Listing subdir/parent/child")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
	}
	
	// cd back to where we started.
    err = os.Chdir("../../..")
	check(err)
	
	/* We can also visit a directory recursively, including all its sub-directories. 
	 * Walk accepts a callback function to handle every file or directory visited. */
    fmt.Println("Visiting subdir")
    err = filepath.Walk("subdir", visit)
}

// visit is called for every file or directory found recursively by filepath.Walk.
func visit(path string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(" ", path, info.IsDir())
    return nil
}