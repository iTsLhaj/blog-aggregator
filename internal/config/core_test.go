package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestConfig(t *testing.T) {
	SetupTestingEnvironment(t, "nobody")

	username := "test-user"
	cfg, _ := Read()
	cfg.SetUser(username)

	cfg, _ = Read()
	if cfg.Username != username {
		t.Errorf("username mismatch. expected: %s, got: %s", username, cfg.Username)
	}
}

func TestRead(t *testing.T) {
	testConfig := SetupTestingEnvironment(t, "test_user")

	config, err := Read()
	if err != nil {
		t.Fatalf("Read() failed: %v", err)
	}

	if config.DbUrl != testConfig.DbUrl {
		t.Errorf("DBUrl = %v, want %v", config.DbUrl, testConfig.DbUrl)
	}
	if config.Username != testConfig.Username {
		t.Errorf("User = %v, want %v", config.Username, testConfig.Username)
	}
}

func TestReadMissingFile(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	_, err := Read()
	if err == nil {
		t.Error("Expected error for missing config file")
	}
}

func SetupTestingEnvironment(t testing.TB, username string) Config {
	t.Helper()

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, ".gatorconfig.json")

	testConfig := Config{
		DbUrl:    "postgres://test",
		Username: username,
	}

	data, err := json.Marshal(testConfig)
	if err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		t.Fatal(err)
	}

	t.Setenv("HOME", tmpDir)
	return testConfig
}
