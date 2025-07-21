package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"parse/config"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var filePath string

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Faz o parse de um arquivo de configuração YAML ou JSON",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := loadConfig(filePath)
		if err != nil {
			fmt.Printf("Erro ao ler arquivo: %v\n", err)
			os.Exit(1)
		}
		validateConfig(cfg)
		printConfig(cfg)
	},
}

func loadConfig(path string) (config.Config, error) {
	var cfg config.Config
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	if yaml.Unmarshal(data, &cfg) == nil || json.Unmarshal(data, &cfg) == nil {
		return cfg, nil
	}
	return cfg, fmt.Errorf("formato inválido. Esperado YAML ou JSON válido")
}

func validateConfig(cfg config.Config) {
	for i, server := range cfg.Servers {
		if server.Name == "" || server.Host == "" || server.Port == 0 {
			fmt.Printf("Servidor #%d com campos obrigatórios ausentes\n", i)
		}
	}

	db := cfg.Database
	if db.Host == "" || db.Port == 0 || db.User == "" {
		fmt.Println("Banco de dados com campos obrigatórios ausentes")
	}
}

func printConfig(cfg config.Config) {
	fmt.Printf("Configuração carregada com sucesso:\n%+v\n", cfg)
}

func init() {
	rootCmd.AddCommand(parseCmd)
	parseCmd.Flags().StringVarP(&filePath, "file", "f", "", "Arquivo de configuração (YAML ou JSON)")
	parseCmd.MarkFlagRequired("file")
}
