package infraestructure

import (
	"encoding/json"
	"os"
)

type LocalStorage struct {
	filepath string
}

func NewLocalStorage(filepath string) *LocalStorage {
	return &LocalStorage{filepath: filepath}
}

func (s *LocalStorage) SaveData(data interface{}) error {
	file, err := os.Create(s.filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(data)
}

func (s *LocalStorage) LoadData(v interface{})error{
	file, err := os.Open(s.filepath)
    if err != nil {
        return err
    }
    defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(v)
}