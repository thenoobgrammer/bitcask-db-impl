package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sync"

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

type BitcaskStore struct {
	mu    sync.Mutex
	files []BitcaskHandle
}

var indexes = make(map[string]int)

var store BitcaskStore = BitcaskStore{
	handler: "dmnjkshoqwhorfhn2417801237@0y2941hqhw-$@480&(@)&$)HY*!(?)",
}

func main() {
	handler := store.open("user_db")
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	return path.Dir(b)
}

func (s *BitcaskStore) open(dir_name string) BitcaskHandle {
	_createDir(filepath.Base(dir_name))
	_, err := os.Create(filepath.Join(filepath.Base(dir_name), "bitcask.data"))
	if err != nil {
		fmt.Printf("Data file failed to create: %s", err.Error())
	}
	_, err = os.Create(filepath.Join(filepath.Base(dir_name), "bitcask.metadata"))
	if err != nil {
		fmt.Printf("Metadata file failed to create: %s", err.Error())
	}
	return &dir_name
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
