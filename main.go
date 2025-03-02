package main

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BitcaskHandle *string

type DataStructure struct {
	tstamp  timestamp.Timestamp
	key     string
	value   any
	deleted bool
}

type BitcaskStore struct {
	handler      string
	mu           sync.Mutex
	data         []DataStructure
	last_opened  timestamp.Timestamp
	last_updated timestamp.Timestamp
}

var indexes = make(map[string]int)

var store BitcaskStore = BitcaskStore{
	handler: "dmnjkshoqwhorfhn2417801237@0y2941hqhw-$@480&(@)&$)HY*!(?)",
}

func main() {
	store.open("user_db")
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	return path.Dir(b)
}

func (s *BitcaskStore) open(dir_name string) BitcaskHandle {
	root_dir := rootDir()
	full_path := root_dir + "/" + dir_name
	_createDir(full_path)
	os.Create(filepath.Join(filepath.Base(full_path), "bitcask.data"))
	os.Create(filepath.Join(filepath.Base(full_path), "bitcask.metadata"))
	return &store.handler
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
