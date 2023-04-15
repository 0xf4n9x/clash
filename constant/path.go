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
func (p *path) Resolve(path string) string {
	// Clean path
	cleanedPath := filepath.Clean(path)

	// Splicing path with p.HomeDir()
	joinedPath := filepath.Join(p.HomeDir(), cleanedPath)

	// Get the absolute path of the joinedPath
	absPath, err := filepath.Abs(joinedPath)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return "NoAbsPath"
	}

	// Check if the absPath is still in the p.HomeDir()
	if !strings.HasPrefix(absPath, p.HomeDir()) {
		fmt.Println("Path traversal attempt detected!")
		return "PT"
	}

	return absPath
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
