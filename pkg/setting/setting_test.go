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

// testDBs checks that the DbCfg map was populated correctly.
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

			// Handle the Params field.
			if fieldName == "Params" {
				testDbParams(t, fieldVal.(DbCfgParams))
			} else if fieldVal == "" || fieldVal == nil || fieldVal == 0 {
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

// testDbParams checks that the DbCfgParams struct was populated correctly.
func testDbParams(t *testing.T, params DbCfgParams) {
	// Check that the Params struct was populated correctly.
	paramsVal := reflect.ValueOf(params)
	for i := 0; i < paramsVal.NumField(); i++ {
		// Get the field name and value.
		fieldName := paramsVal.Type().Field(i).Name
		fieldVal := paramsVal.Field(i).Interface()
		fieldType := paramsVal.Type().Field(i).Type

		if fieldVal == "" || fieldVal == nil || fieldVal == 0 {
			// Ignore if empty.
		} else if fieldType == reflect.TypeOf(false) {
			// Handle bool fields separately.
			if fieldVal != false {
				t.Errorf("field %v is not equal to %v", fieldName, false)
			}
		} else if fieldType == reflect.TypeOf(1) {
			// Handle int fields separately.
			if fieldVal != 1 {
				t.Errorf("field %v is not equal to %v", fieldName, 1)
			}
		} else if expectedVal := fmt.Sprintf("param_%v", fieldName); fieldVal != expectedVal {
			// Check that the field value is equal to the expected value.
			t.Errorf("field %v is not equal to %v", fieldVal, expectedVal)
		}
	}
}
