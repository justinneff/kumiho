package entities

import (
	"fmt"
	"os"

	"github.com/justinneff/kumiho/utils"
)

type DatabaseObject struct {
	Name         string   `json:"name"`
	Schema       string   `json:"schema"`
	SourceFile   string   `json:"sourceFile"`
	Hash         string   `json:"hash"`
	Dependencies []string `json:"dependencies"`
	content      []byte
}

func (obj DatabaseObject) Content() []byte {
	return obj.content
}

func (obj DatabaseObject) FullName() string {
	if len(obj.Schema) > 0 {
		return fmt.Sprintf("%s.%s", obj.Schema, obj.Name)
	} else {
		return obj.Name
	}
}

func NewDatabaseObject(sourceFile string, provider Provider) (*DatabaseObject, error) {
	content, err := os.ReadFile(sourceFile)
	if err != nil {
		return nil, err
	}

	schema, name := provider.GetObjectSchemaAndName(content)
	hash := utils.ComputeHash(content)

	obj := &DatabaseObject{Name: name, Schema: schema, SourceFile: sourceFile, Hash: hash, content: content}
	return obj, nil
}
