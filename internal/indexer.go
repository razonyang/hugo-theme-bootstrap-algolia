package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Indexer interface {
	FetchRecords() ([]Record, error)
}

type FileIndexer struct {
	filename string
}

func NewFileIndexer(filename string) Indexer {
	return &FileIndexer{
		filename: filename,
	}
}

func (fi *FileIndexer) FetchRecords() (records []Record, err error) {
	f, err := os.Open(fi.filename)
	if err != nil {
		return nil, fmt.Errorf("failed to parse index file: %s.\n", err.Error())
	}

	defer f.Close()

	data, _ := ioutil.ReadAll(f)

	err = json.Unmarshal([]byte(data), &records)

	return
}

type RemoteIndexer struct {
	url string
}

func NewRemoteIndexer(url string) Indexer {
	return &RemoteIndexer{
		url: url,
	}
}

func (ri *RemoteIndexer) FetchRecords() (records []Record, err error) {
	return
}
