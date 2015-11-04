package conf

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

const ConfigurationFilename = "configuration.bolt"

var loadedDatabase = struct {
	*sync.Mutex
	m map[string]*bolt.DB
}{}

var (
	loadedConfig  Config
	defaultConfig = Config{
		Database: {
			Location: "./db",
		},
	}
)

type Config struct {
	*sync.Mutex
	Database struct {
		Location string
	}
}

// Copy returns a copy of the configuration
func (c *Config) Copy() Config {
	c.Lock()
	defer c.Unlock()

	a := c
	a.Mutex = new(sync.Mutex)
	return a
}

// LoadConfiguration calls OpenConfiguration and replaces the global-scoped
// config with it if no errors were returned.
func LoadConfiguration(filename string) error {
	c, err := OpenConfiguration(filename)
	if err == nil {
		loadedConfig.Lock()
		loadedConfig = c
		loadedConfig.Unlock()
	}
	return err
}

// OpenConfiguration opens a configuration file and returns it.
func OpenConfiguration(filename string) (Config, error) {
	var c Config = defaultConfig.Copy()

	f, err := os.Open(filename)
	if err != nil {
		return c, err
	}

	if err = json.NewDecoder(f).Decode(&c); err != nil {
		return c, fmt.Errorf("invalid configuration file: %s", err)
	}

	return c
}

func DatabaseFunc(name string) func() *bolt.DB {
	return func() *bolt.DB {
		loadedDatabases.Lock()
		defer loadedDatabases.Unlock()
		if db, ok := loadedDatabases.m[name]; ok {
			return db
		}

		filename := filepath.Join(loadedConfig.Database.Location, name)
		db, err := bolt.Open(filename, 0755, nil)
		if err != nil {
			panic(err)
		}

		loadedDatabases.m[name] = db
		return db
	}
}
