package setting

import (
	"reflect"
	"strings"
	"testing"
)

func TestLoadConfigFile(t *testing.T) {
	cfg, err := LoadConfigFile("config.toml.template")
	if err != nil {
		// Handle "invalid Db.<db>" error
		if strings.Contains(err.Error(), "invalid Db.") {
			// ignore because config.toml.template is not a valid config file according to ValidateDbCfg()
		} else {
			t.Fatalf("unable to load config file: %v", err)
		}
	}

	// Check that the Config struct was populated correctly.
	cfgVal := reflect.ValueOf(cfg)
	for i := 0; i < cfgVal.NumField(); i++ {
		// Get the field name and value.
		fieldName := cfgVal.Type().Field(i).Name
		fieldVal := cfgVal.Field(i).Interface()

		switch fieldName {
		case "Db":
			CheckDBs(t, fieldVal.(map[string]DbCfg))

		default:
			// Check that the field value is not empty or 0.
			if fieldVal == "" || fieldVal == nil || fieldVal == 0 {
				t.Errorf("field %v is empty", fieldName)
			}

			// Check that the field value is equal to the expected value.
			if fieldVal != fieldName {
				t.Errorf("field %v is not equal to %v", fieldName, fieldName)
			}
		}
	}
}
