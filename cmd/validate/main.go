package main

import (
	"EnvContractValidation/config"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Entrypoint exclusivo para validar o contrato sem subir a aplicação.
// Usado via: make validate
func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("⚠️  Arquivo .env não encontrado — usando variáveis do sistema")
	}

	if err := config.ValidateEnvContract(".env.example"); err != nil {
		fmt.Fprintf(os.Stderr, "❌ Contrato inválido:\n%v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Contrato de configuração válido.")
}
