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

	// Check that the Config struct was populated correctly.
	cfgVal := reflect.ValueOf(cfg)
	for i := 0; i < cfgVal.NumField(); i++ {
		// Get the field name and value.
		fieldName := cfgVal.Type().Field(i).Name
		fieldVal := cfgVal.Field(i).Interface()

		switch fieldName {
		case "DBs":
			testDBs(t, fieldVal.(map[string]DbCfg))

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

// testDBs checks that the DBs map was populated correctly.
func testDBs(t *testing.T, dbs map[string]DbCfg) {
	for dbName, db := range dbs {
		// Check that the DB name is not empty.
		if dbName == "" {
			t.Errorf("DB name is empty")
		}

		// Check that the DB struct was populated correctly.
		dbVal := reflect.ValueOf(db)
		for i := 0; i < dbVal.NumField(); i++ {
			// Get the field name and value.
			fieldName := dbVal.Type().Field(i).Name
			fieldVal := dbVal.Field(i).Interface()

			if fieldVal == "" || fieldVal == nil || fieldVal == 0 {
				// Check that the field value is not empty or 0.
				t.Errorf("field %v is empty", fieldName)
			} else if fieldName == "Port" {
				// Handle port field separately.
				if fieldVal != 1234 {
					t.Errorf("field %v is not equal to %v", fieldName, 1234)
				}
			} else if expectedVal := fmt.Sprintf("%v_%v", dbName, fieldName); fieldVal != expectedVal {
				// Check that the field value is equal to the expected value.
				// expectedVal = db_field
				t.Errorf("field %v is not equal to %v", fieldVal, expectedVal)
			}
		}
	}
}
