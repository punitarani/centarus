package dotenv

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// Load loads variables from *.env file into the environment.
//
// param fp: .env file path. fp can be relative filename or full absolute path.
// param check: if true, it will check if the file exists.
//
// return: error if any.
//
// Key-value pairs are of the form: KEY=value
// Comment lines start with '#'
// Empty lines are ignored
func Load(envFile string, check bool) error {
	// Check if file exists
	if _, err := os.Stat(envFile); err != nil {
		if check {
			return err
		}
		return nil
	}

	// Open file
	file, err := os.Open(envFile)
	if err != nil {
		return err
	}

	// Close file on function exit
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Strip spaces, tabs and delimiters
		line = strings.TrimSpace(line)

		// Skip empty lines and comments
		if len(line) > 0 && line[0] != '#' {
			// Parse line
			if key, value, err := parseLine(line); err == nil {
				// Check if environment variable is already set
				_, ok := os.LookupEnv(key)

				// Write environment variable if it is not set
				if !ok {
					// Set environment variable
					if err := os.Setenv(key, value); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

// Create creates a new *.env file and writes environment variables to it.
//
// param fp: .env file path. fp can be relative filename or full absolute path.
// param envVars: environment variables to write to the file.
//
// return: file path and error if any.
//
// Writes key-value pairs as: KEY=value
func Create(fp string, envVars map[string]string) (string, error) {
	// Validate file path
	if fp == "" {
		return "", errors.New("file path is empty: " + fp)
	}

	// Convert fp to absolute path
	if fp, err := filepath.Abs(fp); err != nil {
		return fp, err
	}

	// Validate .env file extension
	if filepath.Ext(fp) != ".env" {
		return "", errors.New("invalid .env file extension: " + fp)
	}

	// Check and create file
	if _, err := os.Stat(fp); err != nil {
		// Create .env file
		if _, err := os.Create(fp); err != nil {
			return fp, err
		}
	}

	// Open file
	file, err := os.OpenFile(fp, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return fp, err
	}

	// Close file on function return
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	// Write environment variables to the file
	for key, value := range envVars {
		if _, err = file.WriteString(key + "=" + value + "\n"); err != nil {
			return fp, err
		}
	}

	return fp, nil
}

// parseLine parses line and returns key and value.
func parseLine(line string) (string, string, error) {
	if strings.TrimSpace(line) == "" {
		return "", "", errors.New("empty line")
	}

	// Find the index of the "=" delimiter
	delimiterIndex := strings.Index(line, "=")
	if delimiterIndex == -1 {
		return "", "", errors.New("invalid line")
	}

	// Split the line into key and value parts
	key := strings.TrimSpace(line[:delimiterIndex])
	value := strings.TrimSpace(line[delimiterIndex+1:])

	// Validate key and value
	if key == "" {
		return "", "", errors.New("missing key")
	}
	if value == "" {
		return "", "", errors.New("missing value")
	}

	return key, value, nil
}
