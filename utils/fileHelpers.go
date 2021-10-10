package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"path"
)

func ComputeHash(data []byte) string {
	h := sha1.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
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
