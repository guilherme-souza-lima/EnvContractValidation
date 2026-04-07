# Cenário 3 — Chave a mais no .env ❌

## Situação

> **Este é o caso que causou dor de cabeça na vida real.**

Um desenvolvedor adicionou `REDIS_URL` e `INTERNAL_API_KEY` no próprio `.env`  
para implementar uma feature, mas **esqueceu de atualizar o `.env.example`**.

O projeto funcionava perfeitamente na máquina dele.  
Quando outro desenvolvedor clonou e tentou subir — a aplicação travou em runtime  
com erros misteriosos, sem nenhuma mensagem clara do que estava faltando.

**Sem a validação de contrato, esse erro só aparecia longe do ponto real do problema.**  
Com a validação, a aplicação recusa o boot imediatamente e aponta exatamente o que está errado.

## Arquivos

**.env.example** (desatualizado — falta documentar as novas chaves):
```
DATABASE_URL=
PORT=
JWT_SECRET=
```

**.env** (tem chaves que não estão no contrato):
```
DATABASE_URL=postgres://user:pass@localhost:5432/mydb
PORT=8080
JWT_SECRET=supersecretkey
REDIS_URL=redis://localhost:6379
INTERNAL_API_KEY=abc123
```

## Output esperado

```
❌ Configuração inválida:
contrato de configuração violado:

chaves presentes no .env mas AUSENTES no .env.example:
  - REDIS_URL
  - INTERNAL_API_KEY

  ⚠️  Atualize o .env.example — ele é a fonte da verdade!
```

## Como corrigir

Atualizar o `.env.example` com as novas chaves (sem valores reais):

```
DATABASE_URL=
PORT=
JWT_SECRET=
REDIS_URL=
INTERNAL_API_KEY=
```

E commitar o `.env.example` para que todos os desenvolvedores saibam da mudança.
