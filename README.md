# Go-Project-Layout
gRPC + SQLC + Mongo + Redis + Kafka and ID pool

## Stack structures

* Postgres: SQL database
* MongoDB: Document database
* Redis: Cache
* Kafka: Message queue
* API: Go App serves gRPC

## Getting stated

* Requirement:
  * SQLC: <https://github.com/kyleconroy/sqlc>
  * Migrate: <https://github.com/golang-migrate/migrate>
  * Docker
* Clone

```bash
git clone https://github.com/golang-migrate/migrate
```

* Create config.yml file
  * build/gitToken: token for getting private go package (grpc) 


* Start API

```bash
run cmd/client/main.go dev
```

* Running command inside a container

```bash
docker exec -it explorer-api-1 ash
```

## Project structures


1. `/api` Document for api. OpenAPI or Swagger specifications, JSON Schema files, protocol definition files.
2. `/build` Script for build. Docker file for local, dev, production
3. `/cmd` The entry point for our application
4. `/config` Initialization of the general app configurations
5. `/internal` Internal logic of application. Internal contain module which has:

   
   1. `/handler` Handle request to module. For each module, it was named with user domain, such as: public_handler for end-user, ops_handler for operator,…
   2. `/pkg` Logic of module:
      * `service`
      * `model`
      * `const`
      * `validation`
6. `/pkg` logic can be imported into a different proj
7. `/script` Scripts for building, installing, analyzing, and conducting other operations on the project

## Module


1. pkg/db/postgres

   
   1. Migration SQL script

      `YYYYMMDDHHMMSS_name.up.sql` For migrate up

      `YYYYMMDDHHMMSS_name.down.sql` For migrate down
   2. Generate SQLC code

      ```bash
      go run cmd/client/main.go sqlc
      ```
   3. Migrate SQL Database

      ```bash
      go run cmd/client/main.go migrateup
      ```
2. pkg/db/mongo

   
   1. Create new collection/repo
      * Define model of Collection in `db/mongo/model/`
      * Define logic of repo in `db/mongo/repo/`
      * Define repo in Driver struct `db/mongo/driver/`
3. pkg/uuid
   * UUID format
     * Len: 20 numeric digit \~ 64 bit
     * 41 bit for timestamp
     * 6 bit for sharding id
     * 6 bit for type
     * 11 bit for counter
   * Load capacity: 2048 per nano second \~ 2M per second


