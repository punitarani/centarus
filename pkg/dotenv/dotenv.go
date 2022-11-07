package dotenv

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// Load loads variables from *.env file into the environment.
func Load(envFile string) error {
	// Check if file exists
	if _, err := os.Stat(envFile); err != nil {
		return err
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

// CreateDotEnv create
func CreateDotEnv(fp string, envVars map[string]string) (string, error) {
	// Convert fp to absolute path
	if fp, err := filepath.Abs(fp); err != nil {
		return fp, err
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
	// Split line by delimiter "="
	parts := strings.Split(line, "=")
	if len(parts) != 2 {
		return "", "", errors.New("invalid line")
	}

	// Strip spaces, tabs and delimiters
	key := strings.TrimSpace(parts[0])
	value := strings.TrimSpace(parts[1])

	return key, value, nil
}
