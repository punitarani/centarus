package dotenv

import (
	"github.com/punitarani/centarus/pkg/random"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	// Generate random environment variables
	envVars := genRandomVars()

	// Check empty file path error
	if _, err := Create("", envVars); err == nil {
		t.Error("expected error for empty file path")
	}

	// Check invalid file extension error
	if _, err := Create("invalid.ENV", envVars); err == nil {
		t.Error("expected error for invalid file extension")
	}

	// Create temp .env file
	file, err := os.CreateTemp(os.TempDir(), "test.*.env")
	if err != nil {
		return
	}

	// Create .env file
	fp, err := Create(file.Name(), envVars)
	if err != nil {
		t.Error(err)
	}

	// Delete .env file on function return
	defer deleteDotEnv(file)

	// Validate fp
	if fp != file.Name() {
		t.Error("invalid file path")
	}
}

func TestLoad(t *testing.T) {
	// Generate random environment variables
	envVars := genRandomVars()

	// Test 'check' param functionality
	if err := Load("missing.env", false); err != nil {
		t.Error("expected no error for missing file with check=false")
	}
	if err := Load("missing.env", true); err == nil {
		t.Error("expected error for missing file with check=true")
	}

	// Create temp .env file
	file, err := os.CreateTemp(os.TempDir(), "test.*.env")
	if err != nil {
		return
	}

	// Create .env file
	fp, err := Create(file.Name(), envVars)
	if err != nil {
		t.Error(err)
	}

	// Delete .env file on function return
	defer deleteDotEnv(file)

	// Load environment variables from the .env file
	err = Load(fp, true)
	if err != nil {
		t.Error(err)
	}

	// Check if environment variables are loaded
	for key, value := range envVars {
		if os.Getenv(key) != value {
			t.Errorf("Environment variable %s is not loaded", key)
		}
	}
}

// genRandomVars generates random environment variables of varying lengths.
//
// The function returns a map of the variables.
func genRandomVars() map[string]string {
	envVars := map[string]string{}

	// Generate 3 random environment variables
	// 	Length of key 	=  4 x index
	// 	Length of value	= 16 x index
	for i := 0; i < 3; i++ {
		key, _ := random.String(4*(i+1), random.AlphaUpper)
		value, _ := random.String(16*(i+1), random.AlphaNumeric)
		envVars[key] = value
	}

	return envVars
}

// deleteDotEnv deletes the .env file.
func deleteDotEnv(file *os.File) {
	// Validate file input
	if _, err := os.Stat(file.Name()); err != nil {
		return
	}

	// Close file
	if err := file.Close(); err != nil {
		return
	}

	// Remove file
	if err := os.Remove(file.Name()); err != nil {
		return
	}
}
