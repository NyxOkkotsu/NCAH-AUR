package security

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type TrustStore struct {
	TrustedHashes map[string]bool `json:"trusted_hashes"`
}

func getConfigPath() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir = os.Getenv("HOME")
	}
	return filepath.Join(configDir, "ncah", "trusted.json")
}

func CalculateHash(content string) string {
	hash := sha256.Sum256([]byte(content))
	return fmt.Sprintf("%x", hash)
}

func IsHashTrusted(hash string) bool {
	filePath := getConfigPath()
	file, err := os.Open(filePath)
	if err != nil {
		return false
	}
	defer file.Close()

	var store TrustStore
	if err := json.NewDecoder(file).Decode(&store); err != nil {
		return false
	}
	return store.TrustedHashes[hash]
}

func SaveTrustedHash(hash string) error {
	filePath := getConfigPath()
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return err
	}

	store := TrustStore{TrustedHashes: make(map[string]bool)}
	if file, err := os.Open(filePath); err == nil {
		_ = json.NewDecoder(file).Decode(&store)
		file.Close()
	}

	store.TrustedHashes[hash] = true
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(store)
}