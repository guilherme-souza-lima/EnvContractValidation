package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ValidateEnvContract(examplePath string) error {
	exampleKeys, err := parseEnvFileKeys(examplePath)
	if err != nil {
		return fmt.Errorf("não foi possível ler o contrato %s: %w", examplePath, err)
	}

	// .env é sempre .env — não tem motivo para ser configurável
	envKeys, err := parseEnvFileKeys(".env")
	if err != nil {
		return fmt.Errorf("arquivo .env não encontrado — copie o .env.example e preencha os valores: %w", err)
	}

	exampleSet := toSet(exampleKeys)
	envSet := toSet(envKeys)

	// chaves no .env.example que não estão no .env
	var missing []string
	for key := range exampleSet {
		if _, found := envSet[key]; !found {
			missing = append(missing, key)
		}
	}

	// chaves no .env que não estão no .env.example → o caso que causou dor de cabeça
	var extra []string
	for key := range envSet {
		if _, found := exampleSet[key]; !found {
			extra = append(extra, key)
		}
	}

	var errs []string

	if len(missing) > 0 {
		errs = append(errs, fmt.Sprintf(
			"chaves presentes no .env.example mas AUSENTES no .env:\n  - %s",
			strings.Join(missing, "\n  - "),
		))
	}

	if len(extra) > 0 {
		errs = append(errs, fmt.Sprintf(
			"chaves presentes no .env mas AUSENTES no .env.example:\n  - %s\n\n  ⚠️  Atualize o .env.example — ele é a fonte da verdade!",
			strings.Join(extra, "\n  - "),
		))
	}

	if len(errs) > 0 {
		return fmt.Errorf("contrato de configuração violado:\n\n%s", strings.Join(errs, "\n\n"))
	}

	return nil
}

func parseEnvFileKeys(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var keys []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if key := strings.TrimSpace(parts[0]); key != "" {
			keys = append(keys, key)
		}
	}
	return keys, scanner.Err()
}

func toSet(keys []string) map[string]struct{} {
	set := make(map[string]struct{}, len(keys))
	for _, k := range keys {
		set[k] = struct{}{}
	}
	return set
}
