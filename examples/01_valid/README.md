# Cenário 1 — Contrato Válido ✅

## Situação

O `.env` possui exatamente as mesmas chaves que o `.env.example`.  
O contrato está respeitado e a aplicação sobe normalmente.

## Arquivos

**.env.example** (fonte da verdade):
```
DATABASE_URL=
PORT=
JWT_SECRET=
```

**.env** (valores reais):
```
DATABASE_URL=postgres://user:pass@localhost:5432/mydb
PORT=8080
JWT_SECRET=supersecretkey
```

## Output esperado

```
✅ Contrato de configuração válido.
Servidor iniciando na porta 8080
```
