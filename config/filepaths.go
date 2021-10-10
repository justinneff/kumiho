package config

import (
	"os"
	"path"

	"github.com/spf13/viper"
)

func GetDatabaseDir() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return path.Join(cwd, viper.GetString("Dir")), nil
}

func GetObjectDir(objectType, schema string) (string, error) {
	dbDir, err := GetDatabaseDir()
	if err != nil {
		return "", err
	}

	return path.Join(dbDir, "objects", schema, objectType), nil
}

func GetScriptDir(scriptType string) (string, error) {
	dbDir, err := GetDatabaseDir()
	if err != nil {
		return "", err
	}

	return path.Join(dbDir, scriptType), nil
}
