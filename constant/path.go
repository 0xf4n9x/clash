package constant

import (
	"fmt"
	"os"
	P "path"
	"path/filepath"
	"strings"
)

const Name = "clash"

// Path is used to get the configuration path
var Path = func() *path {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir, _ = os.Getwd()
	}

	homeDir = P.Join(homeDir, ".config", Name)
	return &path{homeDir: homeDir, configFile: "config.yaml"}
}()

type path struct {
	homeDir    string
	configFile string
}

// SetHomeDir is used to set the configuration path
func SetHomeDir(root string) {
	Path.homeDir = root
}

// SetConfig is used to set the configuration file
func SetConfig(file string) {
	Path.configFile = file
}

func (p *path) HomeDir() string {
	return p.homeDir
}

func (p *path) Config() string {
	return p.configFile
}

// Resolve return a absolute path or a relative path with homedir
func (p *path) Resolve(path string) (string, error) {
	if filepath.IsAbs(path) {
		return "", fmt.Errorf("Absolute paths are not allowed")
	}

	joinedPath := filepath.Join(p.HomeDir(), path)
	cleanedPath := filepath.Clean(joinedPath)
	absPath, err := filepath.Abs(cleanedPath)
	if err != nil {
		return "", fmt.Errorf("Error getting absolute path: %v", err)
	}

	if !strings.HasPrefix(absPath, p.HomeDir()) {
		return "", fmt.Errorf("Path traversal attempt detected!")
	}

	return absPath, nil
}

func (p *path) MMDB() string {
	return P.Join(p.homeDir, "Country.mmdb")
}

func (p *path) OldCache() string {
	return P.Join(p.homeDir, ".cache")
}

func (p *path) Cache() string {
	return P.Join(p.homeDir, "cache.db")
}
