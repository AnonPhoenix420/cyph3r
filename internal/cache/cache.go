package cache

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type CacheEntry struct {
	Timestamp time.Time       `json:"timestamp"`
	Data      json.RawMessage `json:"data"`
}

type ResponseCache struct {
	filePath string
	storage  map[string]CacheEntry
}

func NewResponseCache() (*ResponseCache, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	dir := filepath.Join(home, ".config", "cyph3r")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	filePath := filepath.Join(dir, "cache.json")
	rc := &ResponseCache{
		filePath: filePath,
		storage:  make(map[string]CacheEntry),
	}

	rc.load()
	return rc, nil
}

func (rc *ResponseCache) load() {
	file, err := os.ReadFile(rc.filePath)
	if err != nil {
		return
	}
	_ = json.Unmarshal(file, &rc.storage)
}

func (rc *ResponseCache) save() error {
	data, err := json.MarshalIndent(rc.storage, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(rc.filePath, data, 0644)
}

func (rc *ResponseCache) Get(key string) ([]byte, bool) {
	entry, exists := rc.storage[key]
	if !exists {
		return nil, false
	}

	if time.Since(entry.Timestamp) > 24*time.Hour {
		delete(rc.storage, key)
		_ = rc.save()
		return nil, false
	}

	return entry.Data, true
}

func (rc *ResponseCache) Set(key string) error {
	raw, err := json.Marshal(data)
	if err != nil {
		return err
	}

	rc.storage[key] = CacheEntry{
		Timestamp: time.Now(),
		Data:      raw,
	}
	return rc.save()
}
