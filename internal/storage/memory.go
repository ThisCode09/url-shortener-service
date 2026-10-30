package storage

import (
	"encoding/json"
	"os"
	"sync"
)

type JSONStore struct {
	mu       sync.RWMutex
	data     map[string]string
	filePath string
}

func NewJSONStore(filepath string) (*JSONStore, error) {
	store := &JSONStore{
		data:     make(map[string]string),
		filePath: filepath,
	}
	file, err := os.ReadFile(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return store, nil
		}
		return nil, err
	}
	if len(file) > 0 {
		err = json.Unmarshal(file, &store.data)
		if err != nil {
			return nil, err
		}
	}
	return store, nil
}

func (s *JSONStore) Set(code, url string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[code] = url

	bytes, err := json.MarshalIndent(s.data, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filePath, bytes, 0644)

}

func(s *JSONStore) Get(code string) (string,bool){
	s.mu.RLock()
	defer s.mu.RUnlock()

	url,exists:=s.data[code]
	return url,exists
}
