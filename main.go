package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/golang/protobuf/ptypes/timestamp"
)

var (
	Mu sync.Mutex
)

type BitcaskHandle struct {
	File *os.File
	Path string
	// Index map[string]int64
}

type FileData struct {
	tstamp  timestamp.Timestamp
	key     string
	value   any
	deleted bool
}

func main() {
	open("user_db")
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	return path.Dir(b)
}

func open(dir_name string) *BitcaskHandle {
	Mu.Lock()
	_createDir(filepath.Base(dir_name))
	data_file, err := os.Create(filepath.Join(filepath.Base(dir_name), "bitcask.data"))
	if err != nil {
		fmt.Printf("Data file failed to create: %s", err.Error())
	}
	_, err = os.Create(filepath.Join(filepath.Base(dir_name), "bitcask.metadata"))
	if err != nil {
		fmt.Printf("Metadata file failed to create: %s", err.Error())
	}
	Mu.Unlock()
	return &BitcaskHandle{
		File: data_file,
		Path: dir_name,
	}
}

func _createDir(name string) {
	root_dir := rootDir()
	full_path := root_dir + "/" + name
	os.MkdirAll(full_path, os.ModePerm)
}
