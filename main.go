package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const MAX_FILE_SIZE = 1024 * 1024

type Entry struct {
	Key    string
	Value  string
	Offset int64
	Size   int64
}

type Bitcask struct {
	ActiveFile *os.File
	Path       string
	Index      map[string]Entry
}

type Store struct {
	bitcask Bitcask
}

func main() {
	open("./bitcask")
}

func open(path string) *Store {
	bitcask := &Bitcask{
		Index: make(map[string]Entry),
	}

	rebuildIndex(path, bitcask)

	return &Store{bitcask: *bitcask}
}

func rebuildIndex(path string, bitcask *Bitcask) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if bitcask.ActiveFile != nil {
			return nil
		}

		if filepath.Ext(info.Name()) == ".data" {
			if info.Size() < MAX_FILE_SIZE {
				file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
				if err != nil {
					return err
				}
				bitcask.ActiveFile = file

				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					line := scanner.Bytes()

					key, value, err := parseKeyValue(line)
					if err != nil {
						return err
					}

					offset, err := file.Seek(0, io.SeekCurrent)
					if err != nil {
						return err
					}

					bitcask.Index[key] = Entry{Key: key, Value: value, Offset: offset, Size: info.Size()}
				}
				return nil
			}

		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func parseKeyValue(line []byte) (string, string, error) {
	parts := bytes.SplitN(line, []byte(","), 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid line format: %s", string(line))
	}
	return string(parts[0]), string(parts[1]), nil
}
