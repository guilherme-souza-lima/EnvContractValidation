# Cenário 2 — Chave ausente no .env ❌

## Situação

O `.env.example` define `JWT_SECRET` mas o desenvolvedor esqueceu de adicionar  
essa chave no `.env` local. A aplicação **não sobe**.

## Arquivos

**.env.example** (fonte da verdade):
```
DATABASE_URL=
PORT=
JWT_SECRET=
```

**.env** (incompleto):
```
DATABASE_URL=postgres://user:pass@localhost:5432/mydb
PORT=8080
```

## Output esperado

```
❌ Configuração inválida:
contrato de configuração violado:

chaves presentes no .env.example mas AUSENTES no .env:
  - JWT_SECRET
```

## Como corrigir

Adicionar a chave faltante no `.env`:

```
JWT_SECRET=supersecretkey
```
