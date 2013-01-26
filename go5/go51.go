package gocfg

import (
	"errors"
	"strconv"
)

type Config struct {
	data map[string]map[string]string
}

// Create a new, empty configuration.
func NewConfig() *Config {
	return &Config{data : make(map[string]map[string]string)}
}

// Add a new section to the configuration, if it's not already added.
func (c *Config) AddSection(section string) {
	if _, ok := c.data[section]; !ok {
		c.data[section] = make(map[string]string)
	}
}

// Add a new key,value pair to the given section.
// Returns an error if the section doesn't exist.
func (c *Config) Add(section, key, value string) error {
	if _, ok := c.data[section]; !ok {
		return errors.New("adding to non-existent section")
	}
	c.data[section][key] = value
	return nil
}

// Get a string value from a specific section.
// Returns an error if the section or key doesn't
// exist in the configuration.
func (c *Config) GetString(section, key string) (str string, err error) {
	return c.getValue(section, key)
}

// Get a bool value from a specific section.
// Returns an error if the section or key doesn't
// exist in the configuration.
func (c *Config) GetBool(section, key string) (b bool, err error) {
	if val, err := c.getValue(section, key); err == nil {
		return strconv.ParseBool(val)
	}
	return false, err
}

// Get an integer value from a specific section.
// Returns an error if the section or key doesn't
// exist in the configuration.
func (c *Config) GetInt(section, key string) (i int64, err error) {
	if val, err := c.getValue(section, key); err == nil {
		return strconv.ParseInt(val, 0, 64)
	}
	return 0, err
}

// Get an floating point value from a specific section.
// Returns an error if the section or key doesn't
// exist in the configuration.
func (c *Config) GetFloat(section, key string) (i float64, err error) {
	if val, err := c.getValue(section, key); err == nil {
		return strconv.ParseFloat(val, 64)
	}
	return 0, err
}

func (c *Config) getValue(section, key string) (str string, err error) {
	if sec, ok := c.data[section]; ok {
		if value, ok := sec[key]; ok {
			return value, nil
		}
		return "", errors.New("key does not exist")
	}
	return "", errors.New("section does not exist")
}
