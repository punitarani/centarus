package dotenv

import (
	"github.com/punitarani/centarus/pkg/random"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// Generate random environment variables
	envVars := genRandomVars()

	// Create temp .env file
	file, err := os.CreateTemp(os.TempDir(), "test.*.env")
	if err != nil {
		return
	}

	// Create .env file
	fp, err := CreateDotEnv(file.Name(), envVars)
	if err != nil {
		t.Error(err)
	}

	// Load environment variables from the .env file
	err = Load(fp)
	if err != nil {
		t.Error(err)
	}

	// Check if environment variables are loaded
	for key, value := range envVars {
		if os.Getenv(key) != value {
			t.Errorf("Environment variable %s is not loaded", key)
		}
	}

	// Close file
	if err = file.Close(); err != nil {
		t.Error(err)
	}

	// Remove file
	if err = os.Remove(fp); err != nil {
		t.Error(err)
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
