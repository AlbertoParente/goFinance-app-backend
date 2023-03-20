-- Commands --
-- Check existing container:
docker ps

-- Check existing docker images:
docker images

-- Download a postgres docker container in version 14-alpine:
docker pull postgres:14-alpine

-- Creates the postgres image in relation to the previously created container:
docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine

-- Run script via cmd in postgres in the created instance:
docker exec -it postgres psql -U postgres

-- Exit the scripting environment:	
\q

-- Stop container by nome:
docker stop <container name>

-- Delete a container by name:
docker rm <container name>

-- Start container by name:
docker start <container name>

-- Install scoop on pc via power shell:
Set-ExecutionPolicy RemoteSigned -Scope CurrentUser
irm get.scoop.sh | iex

-- Install migrations on pc via power shell:
scoop install migrate

-- Create the database in postgres from the created migrations:
docker exec -it postgres createdb --username=postgres --owner=postgres go_finance

-- Delete the postgres database from the created migrations:
docker exec -it postgres dropdb --username=postgres --owner=postgres go_finance

-- Run migrations to create tables:
-- Note: Run via powershell as the vscode terminal may not recognize the "migrate" syntax
migrate -path db/migration -database "postgres://postgres:postgres@localhost:5432/go_finance?sslmodedisable" -verbose up

-- Download a SQLC docker container
docker pull kjconroy/sqlc

Update project dependencies and create go.mod file:
go mod tidy

-- Executa SQLC generate para gerar os cruds:
-- Note: This command runs better in CMD because in powershell it doesn't understand correctly. 
    -- In the CMD you have to be in the root folder of the project to be able to find the SQLC configuration file "sqlc.yaml".
docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

-- Install dependency pq
go get github.com/lib/pq

-- Install dependency pq
go get github.com/stretchr/testify

-- Add project sql import
go get -t github.com/albertoparente/go-finance-app/db/sqlc

-- Install dependency gin
go get -u github.com/gin-gonic/gin

-- Install dependency godotenv
go get github.com/joho/godotenv

-- Install dependency jwt
go get -u github.com/golang-jwt/jwt/v4
