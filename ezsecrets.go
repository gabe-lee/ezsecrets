package ezsecrets

import (
	"os"
	"strings"
)

func LoadSecrets(filename string, separator string) (map[string]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	strFile := string(file)
	pairs := strings.Fields(strFile)
	secrets := make(map[string]string, len(pairs))
	for _, pair := range pairs {
		parts := strings.Split(pair, separator)
		if len(parts) < 2 {
			continue
		}
		key := parts[0]
		val := parts[1]
		secrets[key] = val
	}
	return secrets, nil
}
