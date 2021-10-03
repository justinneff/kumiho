package publishing

import (
	"crypto/sha1"
	"encoding/hex"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/justinneff/kumiho/providers"
)

type DatabaseObject struct {
	Name         string   `json:"name"`
	Schema       string   `json:"schema"`
	SourceFile   string   `json:"sourceFile"`
	Hash         string   `json:"hash"`
	Dependencies []string `json:"dependencies"`
}

func GetDatabaseObjectPaths(dbDir string) ([]string, error) {
	objectsDir := filepath.Join(dbDir, "objects")

	var objectFiles []string

	err := filepath.WalkDir(objectsDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".sql" {
			objectFiles = append(objectFiles, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return objectFiles, nil
}

func CreateDatabaseObject(filename string, provider providers.Provider) (*DatabaseObject, error) {
	item := DatabaseObject{}
	item.SourceFile = filename

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	schema, name := provider.GetObjectSchemaAndName(data)
	item.Schema = schema
	item.Name = name

	// Compute checksum hash of the file contents. This is used to determine
	// if the file has been modified.
	h := sha1.New()
	h.Write(data)
	item.Hash = hex.EncodeToString(h.Sum(nil))

	return &item, nil
}
