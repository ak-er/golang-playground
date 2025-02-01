package main

import "fmt"

// embedding-interface -> Interface can be embedded within other interfaces to create more complex interfaces.

// Reader interface
type Reader interface {
	Read() string
}

// Writer interface
type Writer interface {
	Write(content string)
}

// Embed Interface
type ReadWriter interface {
	Reader
	Writer
}

// Struct for File -> implement interface
type File struct {
	Content string
}

func (f *File) Read() string {
	return f.Content
}

func (f *File) Write(content string) {
	f.Content = content
}

func EmbeddingInterface() {
	file := File{Content: "Initial Content"}
	var rw ReadWriter = &file
	fmt.Println(rw.Read())
	rw.Write("Updated Content")
	fmt.Println(rw.Read())
}
