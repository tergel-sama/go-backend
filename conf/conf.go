package conf

import (
	"errors"
	"net/url"
	"os"

	"gopkg.in/yaml.v3"
)

type Database struct {
	Driver string `yaml:"driver"`
	Url    string `yaml:"url"`
	Parsed *url.URL
}

type Migrate struct {
	Dir string `yaml:"dir"`
}

type Config struct {
	Database Database `yaml:"database"`
	Migrate  Migrate  `yaml:"migrate"`
}

// Perform validation (i.e., data type, format) for config attributes.
func (c *Config) Validate() error {
	if len(c.Database.Url) <= 0 {
		return errors.New("unable to find database url configuration")
	}

	// Validate the postgres URL
	temp, err := url.Parse(c.Database.Url)
	if err != nil {
		return err
	}

	c.Database.Parsed = temp

	return nil
}

// Check if the file exists, and read the contents, unmarshal the yaml.
func CreateConfig(c *Config) error {
	path := "conf/conf.yml"

	// Check if the dir/file exists.
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}

	// Read the file as a byte slice.
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// Unmarshal the byte slice to a config struct.
	err = yaml.Unmarshal(data, c)
	if err != nil {
		return err
	}

	return nil
}
