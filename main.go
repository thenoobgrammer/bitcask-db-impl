package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BitcaskHandle struct {
	File  *os.File
	Path  string
	Index map[string]int64
}

type DataStructure struct {
	tstamp  timestamp.Timestamp
	key     string
	value   any
	deleted bool
}

// type BitcaskStore struct {
// 	Mu    sync.Mutex
// 	DataFile BitcaskHandle
// 	MetadataFile BitcaskHandle
// }

func main() {
	handler := open("user_db")
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	return path.Dir(b)
}

func open(dir_name string) BitcaskHandle {
	_createDir(filepath.Base(dir_name))
	data_file, err := os.Create(filepath.Join(filepath.Base(dir_name), "bitcask.data"))
	if err != nil {
		fmt.Printf("Data file failed to create: %s", err.Error())
	}
	_, err = os.Create(filepath.Join(filepath.Base(dir_name), "bitcask.metadata"))
	if err != nil {
		fmt.Printf("Metadata file failed to create: %s", err.Error())
	}
	return &BitcaskHandle{
		File: data_file,
		Path: dir_name,
		Index: ,
	}
}
func (s *BitcaskStore) openWopts(dir_name string, opts map[string]interface{}) BitcaskHandle {
	return nil
}
func (s *BitcaskStore) get(key string) any {
	store.mu.Lock()
	idx := indexes[key]
	store.mu.Unlock()
	return s.data[idx]
}
func (s *BitcaskStore) put(handler BitcaskHandle, key string, value any) {
	store.mu.Lock()
	s.data = append(s.data, DataStructure{
		key:    key,
		value:  value,
		tstamp: *timestamppb.Now(),
	})
	indexes[key] = len(s.data) - 1
	store.mu.Unlock()
}
func (s *BitcaskStore) delete(handler BitcaskHandle, key string, value any) {
	store.mu.Lock()
	s.data = append(s.data, DataStructure{
		key:     key,
		value:   value,
		tstamp:  *timestamppb.Now(),
		deleted: true,
	})
	delete(indexes, key)
	store.mu.Unlock()
}

func _createDir(name string) {
	root_dir := rootDir()
	full_path := root_dir + "/" + name
	os.MkdirAll(full_path, os.ModePerm)
}
