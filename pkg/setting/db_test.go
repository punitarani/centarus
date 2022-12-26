package setting

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBuildDSN(t *testing.T) {
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

				Params: map[DbCfgParams]string{
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
		got := BuildDSN(&tt.cfg)
		if got != tt.want {
			t.Errorf("BuildDSN() = %v, want %v", got, tt.want)
		}
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
				CheckDbParams(t, fieldVal.(map[DbCfgParams]string))
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

// CheckDbParams checks that the DbCfgParams struct was populated correctly.
func CheckDbParams(t *testing.T, params map[DbCfgParams]string) {
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
