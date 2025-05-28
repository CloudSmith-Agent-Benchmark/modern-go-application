package database

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig_Validate(t *testing.T) {
	t.Parallel()
	tests := map[string]Config{
		"database host is required": {
			Port: 3306,
			User: "root",
			Pass: "",
			Name: "database",
		},
		"database port is required": {
			Host: "localhost",
			User: "root",
			Pass: "",
			Name: "database",
		},
		"database user is required": {
			Host: "localhost",
			Port: 3306,
			Pass: "",
			Name: "database",
		},
		"database name is required": {
			Host: "localhost",
			Port: 3306,
			User: "root",
			Pass: "",
		},
	}

	for name, test := range tests {
		name, test := name, test

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			err := test.Validate()

			require.EqualError(t, err, name)
		})
	}
}

func TestConfig_DSN(t *testing.T) {
	t.Parallel()
	config := Config{
		Host: "host",
		Port: 3306,
		User: "root",
		Pass: "",
		Name: "database",
		Params: map[string]string{
			"parseTime": "true",
		},
	}

	dsn := config.DSN()

	require.Equal(t, "root:@tcp(host:3306)/database?parseTime=true", dsn)
}
