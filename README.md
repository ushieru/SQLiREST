# SQLiREST

SQLiREST es un servidor web que convierte su base de datos SQLite3 directamente en una API REST.
> Inspirado en PostgREST.

## Run
```bash
./sqlirest
```

## Config

### .env
```ini
DATABASE_URL="db.sqlite3" # default -> :memori:
SQL_FILE="init.sql"
PORT=8081 # default -> 8080
```

## Operadores
| REST | Operador |
| ---- | -------- |
| eq   | =        |
| gt   | >        |
| gte  | >=       |
| lt   | <        |
| lte  | <=       |
| neq  | <>       |
| like | like     |

## Ejemplos

[Books](./example/README.md)
