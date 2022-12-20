package setting

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLoadConfigFile(t *testing.T) {
	cfg, err := LoadConfigFile("config.toml.template")
	if err != nil {
		t.Fatalf("unable to load config file: %v", err)
	}

	fmt.Println(cfg)
	// Check that the Config struct was populated correctly.
	cfgVal := reflect.ValueOf(cfg)
	for i := 0; i < cfgVal.NumField(); i++ {
		// Get the field name and value.
		fieldName := cfgVal.Type().Field(i).Name
		fieldVal := cfgVal.Field(i).Interface()

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
