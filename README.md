# EnvContractValidation

> Estudo prático sobre validação de contrato entre `.env` e `.env.example` em Go, baseado nos padrões **Fail Fast**, **12-Factor App** e **Configuration as Contract**.

---

## 🧠 O Problema

Em projetos com múltiplos desenvolvedores, é comum que variáveis de ambiente sejam adicionadas ao `.env` local sem que o `.env.example` seja atualizado. O resultado:

- Outro desenvolvedor clona o projeto, sobe a aplicação e nada funciona
- O erro aparece em runtime, longe do ponto real do problema
- Não existe nenhuma trava que obrigue manter o `.env.example` atualizado

---

## ✅ A Solução

A aplicação **não sobe** se as chaves definidas no `.env.example` não estiverem presentes nas variáveis de ambiente carregadas. O `.env.example` passa a ser o **contrato oficial de configuração** do projeto.

```
❌ Configuração inválida:
variáveis de ambiente obrigatórias não definidas:
  - JWT_SECRET
  - DATABASE_URL

Verifique o arquivo .env e atualize com base no .env.example
```

---

## 📐 Padrões Aplicados

| Padrão | Descrição |
|---|---|
| **Fail Fast** | A aplicação detecta o problema na inicialização e para imediatamente com uma mensagem clara |
| **12-Factor App — Factor III** | Configuração armazenada em variáveis de ambiente, documentada e versionada via `.env.example` |
| **Configuration as Contract** | O `.env.example` define o schema esperado; qualquer desvio é um erro em tempo de boot |
| **Living Documentation** | O `.env.example` sempre reflete o estado real do projeto, pois sem ele a aplicação não funciona |

---

## 🗂️ Estrutura do Projeto

```
EnvContractValidation/
├── .env                  # Variáveis reais — NÃO commitar (está no .gitignore)
├── .env.example          # Contrato de configuração — SEMPRE commitar
├── .gitignore
├── go.mod
├── main.go
└── config/
    ├── config.go         # Carrega e expõe a configuração
    └── env_validator.go  # Lógica de validação do contrato
```

---

## ⚙️ Como Funciona

### 1. `.env.example` define o contrato (sem valores reais)

```bash
# Banco de dados
DATABASE_URL=

# Servidor
PORT=

# Autenticação
JWT_SECRET=
```

### 2. `env_validator.go` lê as chaves do `.env.example` e valida

```go
func ValidateEnvContract(examplePath string) error {
    exampleKeys, err := parseEnvKeys(examplePath)
    // ...
    var missing []string
    for _, key := range exampleKeys {
        if os.Getenv(key) == "" {
            missing = append(missing, key)
        }
    }
    if len(missing) > 0 {
        return fmt.Errorf("variáveis obrigatórias não definidas:\n  - %s", strings.Join(missing, "\n  - "))
    }
    return nil
}
```

### 3. `config.go` carrega o `.env` e chama a validação no boot

```go
func Load() *Config {
    godotenv.Load()

    if err := ValidateEnvContract(".env.example"); err != nil {
        log.Fatalf("❌ Configuração inválida:\n%v", err)
    }

    return &Config{
        DatabaseURL: os.Getenv("DATABASE_URL"),
        Port:        os.Getenv("PORT"),
        JWTSecret:   os.Getenv("JWT_SECRET"),
    }
}
```

---

## 🚀 Como Rodar

```bash
# Clone o repositório
git clone https://github.com/seu-usuario/EnvContractValidation.git
cd EnvContractValidation

# Copie o exemplo e preencha com seus valores
cp .env.example .env

# Instale as dependências
go mod tidy

# Rode a aplicação
go run main.go
```

Se o `.env` estiver incompleto, a aplicação **não sobe** e mostra exatamente o que está faltando.

---

## 📋 Regras do Time

| Arquivo | Commitado? | Contém valores reais? |
|---|---|---|
| `.env.example` | ✅ Sempre | ❌ Nunca |
| `.env` | ❌ No `.gitignore` | ✅ Sim |

### Fluxo obrigatório ao criar uma nova variável de ambiente

```
1. Adicionar no .env         → com o valor real (local)
2. Adicionar no .env.example → sem valor (só a chave)
3. Commitar o .env.example
```

Se o passo 3 for esquecido, a aplicação vai parar de subir para qualquer outro desenvolvedor que tentar rodar o projeto — o que é exatamente a trava que queremos.

---

## 🔄 Alternativas com Biblioteca

Caso prefira usar uma biblioteca que já implemente o padrão **Configuration as Contract** nativamente:

| Biblioteca | Como declara o contrato |
|---|---|
| [`cleanenv`](https://github.com/ilyakaznacheev/cleanenv) | Struct com tags `env-required:"true"` |
| [`envconfig`](https://github.com/kelseyhightower/envconfig) | Struct com tags `required:"true"` |
| [`viper`](https://github.com/spf13/viper) | Combinado com validação manual ou `validator` |

Neste repositório a implementação é **manual e sem dependências extras**, para fins didáticos.

---

## 📚 Referências

- [The Twelve-Factor App — Factor III: Config](https://12factor.net/config)
- [Fail Fast — Martin Fowler](https://martingfowler.com/ieeeSoftware/failFast.pdf)
- [cleanenv — Go](https://github.com/ilyakaznacheev/cleanenv)
- [godotenv — Go](https://github.com/joho/godotenv)

---

## 📄 Licença

MIT
