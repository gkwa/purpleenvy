package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

func ensureParentDirectoryExists(filename string) error {
	dir := filepath.Dir(filename)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create parent directory: %v", err)
	}
	return nil
}

func createTestConfig(filePath string) {
	err := ensureParentDirectoryExists(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello, Golang!")
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return
	}

	fmt.Println("File created successfully.")
}

func main() {
	log.SetFlags(log.Lshortfile)

	// XDG Base Directory paths.
	log.Println("Home data directory:", xdg.DataHome)
	log.Println("Data directories:", xdg.DataDirs)
	log.Println("Home config directory:", xdg.ConfigHome)
	log.Println("Config directories:", xdg.ConfigDirs)
	log.Println("Home state directory:", xdg.StateHome)
	log.Println("Cache directory:", xdg.CacheHome)
	log.Println("Runtime directory:", xdg.RuntimeDir)

	// Other common directories.
	log.Println("Home directory:", xdg.Home)
	log.Println("Application directories:", xdg.ApplicationDirs)
	log.Println("Font directories:", xdg.FontDirs)

	// Obtain a suitable location for application config files.
	// ConfigFile takes one parameter which must contain the name of the file,
	// but it can also contain a set of parent directories. If the directories
	// don't exist, they will be created relative to the base config directory.

	configRelPath := "appname/config.yaml"
	configFilePath, err := xdg.ConfigFile(configRelPath)
	if err != nil {
		log.Fatal(err)
	}
	createTestConfig(configFilePath)
	log.Println("Save the config file at:", configFilePath)

	// For other types of application files use:
	// xdg.DataFile()
	// xdg.StateFile()
	// xdg.CacheFile()
	// xdg.RuntimeFile()

	// Finding application config files.
	// SearchConfigFile takes one parameter which must contain the name of
	// the file, but it can also contain a set of parent directories relative
	// to the config search paths (xdg.ConfigHome and xdg.ConfigDirs).
	configFilePath, err = xdg.SearchConfigFile(configRelPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Config file was found at:", configFilePath)

	// For other types of application files use:
	// xdg.SearchDataFile()
	// xdg.SearchStateFile()
	// xdg.SearchCacheFile()
	// xdg.SearchRuntimeFile()
}
