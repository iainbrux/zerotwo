package tarkov

import (
	"encoding/json"
	"os"
	"time"
)

type tarkovCache struct {
	Timestamp time.Time `json:"timestamp"`
	Items []QueryItem `json:"items"`
}

func (cache *tarkovCache) update(key string, value []QueryItem) {
	cacheItem := tarkovCache{
		Items:     value,
		Timestamp: time.Now(),
	}

	// Read the existing data
	file, err := os.ReadFile("tarkovCache.json")
	data := make(map[string]tarkovCache)
	if err == nil {
		_ = json.Unmarshal(file, &data)
	}

	// Update the specific key with the new value
	data[key] = cacheItem

	// Write the updated data back to the file
	file, _ = json.MarshalIndent(data, "", " ")
	_ = os.WriteFile("tarkovCache.json", file, 0644)
}

func (cache *tarkovCache) read(key string) (tarkovCache, bool) {
	// Read the existing data
	file, err := os.ReadFile("tarkovCache.json")
	if err != nil {
		return tarkovCache{}, false
	}

	data := make(map[string]tarkovCache)
	_ = json.Unmarshal(file, &data)

	// Check if the key exists in the map
	cacheItem, ok := data[key]
	if !ok {
		return tarkovCache{}, false
	}

	return cacheItem, true
}