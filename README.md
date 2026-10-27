# sacoli

API REST simples para gerenciar uma lista de compras, feita em Go com Gin e Swagger.

> Os dados ficam em memória: ao reiniciar o servidor, a lista é zerada.

## Como rodar

```bash
git clone https://github.com/carlosfdsn/sacoli.git
cd sacoli
go run ./cmd/server
```

Servidor em `http://localhost:8080`. Documentação Swagger em `http://localhost:8080/swagger/index.html`.

Para regenerar o Swagger:

```bash
swag init -g cmd/server/main.go --parseInternal
```

## Endpoints

| Método | Rota                 | Descrição                    |
|--------|----------------------|------------------------------|
| GET    | `/health`            | Status da API                |
| GET    | `/items`             | Lista os itens               |
| GET    | `/items/{id}`        | Busca um item                |
| POST   | `/items`             | Cria um item                 |
| PUT    | `/items/{id}`        | Atualiza um item             |
| DELETE | `/items/{id}`        | Remove um item               |
| PATCH  | `/items/{id}/bought` | Marca um item como comprado  |

Exemplo de corpo para criar um item:

```json
{ "name": "Arroz", "quantity": 2, "price": 25.9 }
```

## Estrutura

```
cmd/server/       # ponto de entrada e rotas
internal/handlers # handlers HTTP
internal/models   # modelo Item
docs/             # Swagger gerado
```