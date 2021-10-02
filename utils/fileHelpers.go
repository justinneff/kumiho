package utils

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
)

func GetOutDir(scriptType, schema string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	if len(schema) > 0 {
		return path.Join(cwd, viper.GetString("Dir"), schema, scriptType), nil
	} else {
		return path.Join(cwd, viper.GetString("Dir"), scriptType), nil
	}
}

func WriteOutFile(filename, content string) error {
	if fileinfo, _ := os.Stat(filename); fileinfo != nil {
		return fmt.Errorf("file already exists at %s", filename)
	}

	baseDir := path.Dir(filename)

	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return err
	}

	if err := os.WriteFile(filename, []byte(content), 0777); err != nil {
		return err
	}

	return nil
}
