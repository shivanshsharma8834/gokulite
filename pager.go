package main

import "os"

type Pager struct {
	file        *os.File
	file_length uint32
	pages       [][]byte
	pagesLoaded int
}
