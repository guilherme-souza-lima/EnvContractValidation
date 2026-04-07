.PHONY: run build validate tidy clean help

## run: sobe a aplicação (valida o contrato antes de iniciar)
run:
	go run main.go

## build: compila o binário
build:
	go build -o bin/app main.go

## validate: valida o contrato entre .env e .env.example sem subir a aplicação
validate:
	go run cmd/validate/main.go

## tidy: baixa e organiza as dependências
tidy:
	go mod tidy

## clean: remove binários gerados
clean:
	rm -rf bin/

## help: lista os comandos disponíveis
help:
	@echo ""
	@echo "EnvContractValidation — comandos disponíveis:"
	@echo ""
	@grep -E '^## ' Makefile | sed 's/## /  make /'
	@echo ""
