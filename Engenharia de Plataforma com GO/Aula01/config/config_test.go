package config

import (
	"testing"
)

func TestValidConfig(t *testing.T) {
	cfg := Config{
		Servers: []ServerConfig{
			{
				Name:     "web",
				Host:     "localhost",
				Port:     8080,
				Replicas: 2,
			},
		},
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     5432,
			User:     "admin",
			Password: "secret",
		},
	}
	if cfg.Servers[0].Name == "" {
		t.Error("Nome do servidor não pode estar vazio")
	}
	if cfg.Database.User != "admin" {
		t.Errorf("Esperado usuário 'admin', mas veio %s", cfg.Database.User)
	}
}
