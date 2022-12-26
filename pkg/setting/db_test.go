package setting

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestValidateDbCfg(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		key  string
		val  string
	}{
		{
			name: "Test empty Driver",
			key:  "Driver",
			val:  "",
		},
		{
			name: "Test mysql Driver",
			key:  "Driver",
			val:  "mysql",
		},
		{
			name: "Test empty Username",
			key:  "Username",
			val:  "",
		},
		{
			name: "Test empty Password",
			key:  "Password",
			val:  "",
		},
		{
			name: "Test empty Host",
			key:  "Host",
			val:  "",
		},
		{
			name: "Test 0 Port",
			key:  "Port",
			val:  "0",
		},
		{
			name: "Test 65536 Port",
			key:  "Port",
			val:  "65536",
		},
		{
			name: "Test empty Name",
			key:  "Name",
			val:  "",
		},
	}

	// Valid DB config
	validCfg := DbCfg{
		Driver:   "postgresql",
		Username: "user",
		Password: "pass",
		Host:     "localhost",
		Port:     5432,
		Name:     "test",
	}
	validCfgVal := reflect.ValueOf(validCfg)

	// Check that the valid config is valid
	if err := ValidateDbCfg(&validCfg); err != nil {
		t.Errorf("ValidateDbCfg() = %v, want %v", err, nil)
	}

	// Test valid DB config
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Copy validCfg to cfg
			cfg := DbCfg{}
			cfgVal := reflect.ValueOf(&cfg).Elem()
			for i := 0; i < cfgVal.NumField(); i++ {
				cfgVal.Field(i).Set(validCfgVal.Field(i))
			}

			// Set invalid value
			// Convert to int if necessary
			if cfgVal.FieldByName(test.key).Kind() == reflect.Int {
				val, err := strconv.ParseInt(test.val, 10, 64)
				if err != nil {
					t.Errorf("strconv.Atoi() on %v returned error %v", test.val, err)
				}
				cfgVal.FieldByName(test.key).SetInt(val)
			} else {
				cfgVal.FieldByName(test.key).SetString(test.val)
			}

			// Validate DB config
			err := ValidateDbCfg(&cfg)
			if err == nil {
				t.Errorf("ValidateDbCfg() = nil, want error for %v=%v", test.key, test.val)
			}
		})
	}
}

func TestBuildDSN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		cfg  DbCfg
		want string
	}{
		{
			cfg: DbCfg{
				Driver:   "mysql",
				Username: "user1",
				Password: "pass1",
				Host:     "localhost",
				Port:     1234,
				Name:     "test1",
			},
			want: "mysql://user1:pass1@localhost:1234/test1",
		},
		{
			cfg: DbCfg{
				Driver:   "postgres",
				Username: "user2",
				Password: "pass2",
				Host:     "someHost@west.US-2",
				Port:     54321,
				Name:     "test2",

				Params: map[DbCfgParam]string{
					"SslMode":            "verify-full",
					"SslCert":            "cert1",
					"SslKey":             "key1",
					"SslRootCert":        "rootCert1",
					"SslCrl":             "crl1",
					"AppName":            "app1",
					"FallbackAppName":    "fallbackApp1",
					"ConnectTimeout":     "10",
					"Keepalives":         "true",
					"KeepalivesIdle":     "11",
					"KeepalivesInterval": "12",
					"KeepalivesCount":    "13",
				},
			},
			want: "postgres://user2:pass2@someHost@west.US-2:54321/test2?" +
				"AppName=app1&ConnectTimeout=10&FallbackAppName=fallbackApp1&Keepalives=true&" +
				"KeepalivesCount=13&KeepalivesIdle=11&KeepalivesInterval=12&SslCert=cert1&" +
				"SslCrl=crl1&SslKey=key1&SslMode=verify-full&SslRootCert=rootCert1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := BuildDSN(&tt.cfg)
			if got != tt.want {
				t.Errorf("BuildDSN() = %v, want %v", got, tt.want)
			}
		})
	}
}

// CheckDBs checks that the DbCfg map was populated correctly.
func CheckDBs(t *testing.T, dbs map[string]DbCfg) {
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
			if fieldName == "Params" && fieldVal != nil {
				CheckDbParams(t, fieldVal.(map[DbCfgParam]string))
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

// CheckDbParams checks that the DbCfgParam struct was populated correctly.
func CheckDbParams(t *testing.T, params map[DbCfgParam]string) {
	for k, v := range params {
		// Check that value is not empty
		if v == "" {
			t.Errorf("field %v is empty", k)
		}

		acceptableValues := []string{
			fmt.Sprintf("%v_%v", "Value", k),
			"false",
			"-1",
		}
		// Check that the field value is acceptable
		if !(v == acceptableValues[0] || v == acceptableValues[1] || v == acceptableValues[2]) {
			t.Errorf("field %v is acceptable %v", v, acceptableValues)
		}
	}
}
