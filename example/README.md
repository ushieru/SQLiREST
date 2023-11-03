# SQLiREST

### .env
```ini
SQL_FILE="init.sql"
```

### init.sql
```sql
CREATE TABLE IF NOT EXISTS "authors" (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS "books" (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL CHECK(length(name) > 3),
  author_id INTEGER NOT NULL,
  FOREIGN KEY(author_id) REFERENCES authors(id)
);

INSERT INTO authors(name) VALUES
  ("Octavio Paz"),
  ("Juan Rulfo"),
  ("Rosario Castellanos");

INSERT INTO books(name, author_id) VALUES
  ("El laberinto de la soledad", 1),
  ("Piedra de sol", 1),
  ("Pedro Páramo", 2),
  ("El llano en llamas", 2),
  ("Balún Canán", 3),
  ("Poesía no eres tú", 3);
```

### Rest

```
GET http://localhost:8080/authors
```
> [{"id":1,"name":"Octavio Paz"},{"id":2,"name":"Juan Rulfo"},{"id":3,"name":"Rosario Castellanos"}]

```
POST http://localhost:8080/authors
{ "name": "Valeria Luiselli" }
```
> HTTP/1.1 201 Created

```
GET http://localhost:8080/authors?name=eq.Valeria%20Luiselli
```
> [{"id":4,"name":"Valeria Luiselli"}]

```
GET http://localhost:8080/books
```
> [{"author_id":1,"id":1,"name":"El laberinto de la soledad"},{"author_id":1,"id":2,"name":"Piedra de sol"},{"author_id":2,"id":3,"name":"Pedro Páramo"},{"author_id":2,"id":4,"name":"El llano en llamas"},{"author_id":3,"id":5,"name":"Balún Canán"},{"author_id":3,"id":6,"name":"Poesía no eres tú"}]

```
GET http://localhost:8080/books?select=id,name
```
> [{"id":1,"name":"El laberinto de la soledad"},{"id":2,"name":"Piedra de sol"},{"id":3,"name":"Pedro Páramo"},{"id":4,"name":"El llano en llamas"},{"id":5,"name":"Balún Canán"},{"id":6,"name":"Poesía no eres tú"}]

```
GET http://localhost:8080/books?select=name,id&id=gt.4
```
> [{"id":5,"name":"Balún Canán"},{"id":6,"name":"Poesía no eres tú"}]

```
PATCH http://localhost:8080/books?id=eq.5
{"name": "Oficio de tinieblas"}
```
> HTTP/1.1 200 OK

```
DELETE http://localhost:8080/books?id=eq.5
```
> HTTP/1.1 200 OK

## Lua

Puedes agregar extenciones con lua para hacer consulas mas complejas mientras crece el proyecto

```bash
-- WORKDIR
  |- SQLiREST
  |- extentions
    |- test.lua
```

```lua
-- test.lua
return callRawQuery("select sqlite_version()")
```
```lua
-- inner_join_test.lua
return callRawQuery [[
    SELECT books.id as book_id, books.name as book_name, books.author_id, authors.name as author_name
    FROM books
    INNER JOIN authors
    ON books.author_id = authors.id
]]
```

`callRawQuery` Es una funcion que go inyecta a las extenciones de lua para poder hacer peticiones SQL en crudo.

Llamalo con una peticion GET al nombre de tu archivo.

> Ten cuidado con los nombres de tus extenciones. No uses palabras reservadas ni nombres muy complejos para el Query encode.

```
GET http://localhost:8080/test
```
> [{"sqlite_version()":"3.41.2"}]


```
GET http://localhost:8080/inner_join_test
```
> [{"author_id":1,"author_name":"Octavio Paz","book_id":1,"book_name":"El laberinto de la soledad"},{"author_id":1,"author_name":"Octavio Paz","book_id":2,"book_name":"Piedra de sol"},{"author_id":2,"author_name":"Juan Rulfo","book_id":3,"book_name":"Pedro Páramo"},{"author_id":2,"author_name":"Juan Rulfo","book_id":4,"book_name":"El llano en llamas"},{"author_id":3,"author_name":"Rosario Castellanos","book_id":5,"book_name":"Balún Canán"},{"author_id":3,"author_name":"Rosario Castellanos","book_id":6,"book_name":"Poesía no eres tú"}]
