package main

import (
	"sync"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BitcaskHandle *string

type DataStructure struct {
	tstamp timestamp.Timestamp
	key    string
	value  any
}

type BitcaskStore struct {
	handler      string
	mu           sync.Mutex
	data         []DataStructure
	last_opened  timestamp.Timestamp
	last_updated timestamp.Timestamp
}

type KeyOffsets map[string]uint

var store BitcaskStore = BitcaskStore{
	handler: "dmnjkshoqwhorfhn2417801237@0y2941hqhw-$@480&(@)&$)HY*!(?)",
}

func main() {

}

func (s *BitcaskStore) open(dir_name string) BitcaskHandle {
	return &store.handler
}
func (s *BitcaskStore) openWopts(dir_name string, opts map[string]interface{}) BitcaskHandle {
	return nil
}
func (s *BitcaskStore) get(key string) any {
	store.mu.Lock()

	store.mu.Unlock()
	return nil
}
func (s *BitcaskStore) put(handler BitcaskHandle, key string, value any) {
	store.mu.Lock()
	s.data = append(s.data, DataStructure{
		key:    key,
		value:  value,
		tstamp: *timestamppb.Now(),
	})
	store.mu.Unlock()
}
func (s *BitcaskStore) delete(handler BitcaskHandle, key string, value any) {
	store.mu.Lock()
	s.data = append(s.data, DataStructure{
		key:    key,
		value:  value,
		tstamp: *timestamppb.Now(),
	})
	store.mu.Unlock()
}
