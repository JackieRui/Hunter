package config

import (
	"encoding/json"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	c := LoadConfig("./primary.yml")
	s, _ := json.Marshal(c)
	t.Log(string(s))
}
