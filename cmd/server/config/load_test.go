package config

import (
	"testing"
)

func TestLoadConfiguration(t *testing.T) {
	tests := []struct {
		name       string
		configFile string
		wantErr    bool
	}{
		{
			name:       "Test with a valid configuration file",
			configFile: "testdata/good_test_config.yaml",
			wantErr:    false,
		},
		{
			name:       "Test with an invalid configuration file",
			configFile: "testdata/bad_test_config.yaml",
			wantErr:    true,
		},
		{
			name:       "Test with a non-existent configuration file",
			configFile: "testdata/non_existent_config.yaml",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := LoadConfiguration(tt.configFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfiguration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
