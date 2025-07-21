package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"parse/config"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var file string

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Exibe o conteúdo do server do arquivo YAML ou JSON",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := loadConfigFile(file)
		if err != nil {
			fmt.Printf("Erro ao ler arquivo: %v\n", err)
			os.Exit(1)
		}
		printServerConfig(cfg)
	},
}

func loadConfigFile(path string) (config.Config, error) {
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

func printServerConfig(cfg config.Config) {
	for i, server := range cfg.Servers {
		if !validateServer(server) {
			fmt.Printf("Servidor #%d com campos obrigatórios ausentes\n", i)
		} else {
			showServer(server)
		}
	}
}

func validateServer(server config.ServerConfig) bool {
	return server.Name != "" && server.Host != "" && server.Port != 0
}

func showServer(server config.ServerConfig) {
	fmt.Println("Segue os campos do servidor")
	fmt.Println("Nome :", server.Name)
	fmt.Println("Host :", server.Host)
	fmt.Println("Port :", server.Port)
	fmt.Println("Replicas :", server.Replicas)
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&file, "file", "f", "", "Arquivo de configuração (YAML ou JSON)")
	serverCmd.MarkFlagRequired("file")
}
