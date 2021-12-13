# Neo4j 4.4 impersonation demo

This is a simple application that demonstrates how to configure impersonation.

## Database setup

Start a [Neo4j](https://neo4j.com) server.

On the system DB, run:

```cypher
-- create home DBs & users & roles ("system" DB)
CREATE DATABASE joesDb
CREATE DATABASE janesDb
CREATE USER joe SET PASSWORD 'joespass' SET HOME DATABASE joesDb
CREATE USER jane SET PASSWORD 'janespass' SET HOME DATABASE janesDb
CREATE ROLE impersonated
GRANT ALL GRAPH PRIVILEGES ON HOME GRAPH TO impersonated
GRANT ROLE impersonated TO joe,jane
CREATE ROLE impersonator
GRANT IMPERSONATE (joe, jane) ON DBMS TO impersonator
GRANT ROLE impersonator TO neo4j
-- init data ("joesDb" DB)
CREATE (:FavouriteMovie {title: 'Alien vs. Predator vs. CVE-2021-44228'})
-- init data ("janeDb" DB)
CREATE (:FavouriteMovie {title: 'Roundhay Garden Scene'})
```

## Run

`neo4j` impersonating `joe` in autocommit transaction:

```go
go run ./cmd/autocommit
```

`neo4j` impersonating `jane` in autocommit transaction:

```go
go run ./cmd/tx_func
```