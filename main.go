package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("read.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println(fileInfo.Name())    // return name of the file
	fmt.Println(fileInfo.Size())    // return file size.
	fmt.Println(fileInfo.Mode())    // permissions
	fmt.Println(fileInfo.ModTime()) // modified date-time

	// read file content
	buf := make([]byte, fileInfo.Size())
	bufLen, err := f.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(bufLen, string(buf))

	// another method to read file
	data, err := os.ReadFile("read.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	// Read Directory
	file, err := os.Open("../")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println(file.ReadDir(0)) // read directory

	// create a file
	wf, err := os.Create("write.txt")
	defer wf.Close()
	if err != nil {
		panic(err)
	}
	/*
		wf.WriteString("Hello Golang with code")
		wf.WriteString("\nThis is the append.")
	*/
	// another method to write file
	bytes := []byte("wow")
	_, err = wf.Write(bytes)
	if err != nil {
		panic(err)
	}

	// read and write another file (read one file and write to another file)
	source, err := os.Open("read.txt")
	defer source.Close()
	if err != nil {
		panic(err)
	}

	dest, err := os.Create("read_writer.txt")
	defer dest.Close()
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(source)
	writer := bufio.NewWriter(dest)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}
		err = writer.WriteByte(b)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	// delete a file
	err = os.Remove("random.txt")
	if err != nil {
		panic(err)
	}

}
