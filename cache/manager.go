package cache

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/justinneff/kumiho/config"
	"github.com/justinneff/kumiho/entities"
	"github.com/justinneff/kumiho/utils"
)

func getCacheFilename(hash string) (string, error) {
	cacheDir, err := config.GetCacheDir()
	if err != nil {
		return "", err
	}

	return path.Join(cacheDir, fmt.Sprintf("%s.json", hash)), nil
}

func Clear() error {
	cacheDir, err := config.GetCacheDir()
	if err != nil {
		return err
	}

	return os.RemoveAll(cacheDir)
}

func ReadDatabaseObject(hash string) (*entities.DatabaseObject, error) {
	cacheFilename, err := getCacheFilename(hash)
	if err != nil {
		return nil, err
	}

	// Ignore errors when reading from the cache. If the file is not found in the
	// cache then a new object will be created
	data, _ := os.ReadFile(cacheFilename)
	if len(data) == 0 {
		return nil, nil
	}

	obj := &entities.DatabaseObject{}
	err = json.Unmarshal(data, obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func WriteDatabaseObject(obj entities.DatabaseObject) error {
	cacheFilename, err := getCacheFilename(obj.Hash)
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	return utils.WriteOutFile(cacheFilename, string(jsonData))
}
