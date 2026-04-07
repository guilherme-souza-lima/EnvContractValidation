package main

import (
	"EnvContractValidation/config"
	"fmt"
)

func main() {
	cfg := config.Load() // 💥 trava aqui se .env estiver incompleto

	fmt.Printf("Servidor iniciando na porta %s\n", cfg.Port)
}
