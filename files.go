package payload

import (
	"fmt"
	"os"
)

// Load loads payload.RawMessage message from specified path.
// Returns loaded RawMessage data or load error.
func Load(path string) (data RawMessage, err error) {
	if data, err = os.ReadFile(path); err != nil {
		return nil, fmt.Errorf("%w: load file: %v: path %v", Error, err, path)
	}

	return data, nil
}

// MustLoad loads payload.RawMessage message from specified file path. Panics if load failed.
func MustLoad(path string) (data RawMessage) {
	var err error

	if data, err = Load(path); err != nil {
		panic(err)
	}

	return data
}

// Save stores RawMessage data into specified file path. Returns error if save failed.
func Save(data RawMessage, path string) (err error) {
	if err = os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("%v: file: write data: %w", Error, err)
	}

	return nil
}

// MustSave saves RawMessage data into specified file path. Panics if save failed.
func MustSave(data RawMessage, path string) {
	if err := Save(data, path); err != nil {
		panic(err)
	}
}
