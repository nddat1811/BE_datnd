BINARY = engine
configDB = config/database/dbconfig.yml
env = dynamic
dbUser = ${MYSQL_USERNAME}
dbPass = ${MYSQL_PASSWORD}
dbName = ${MYSQL_NAME}

start:
	gow run .

start-swag:
	swag init && gow run .

swag:
	swag init

migration_new:
	sql-migrate new -config=$(configDB) -env=$(env) $(name)

migration_up: $(configDB)
	- sql-migrate up -config=$(configDB) -env=$(env) --dryrun
	- sql-migrate up -config=$(configDB) -env=$(env)

migration_down: $(configDB)
	- sql-migrate down -config=$(configDB) -env=$(env) --dryrun -limit=0
	- sql-migrate down -config=$(configDB) -env=$(env) -limit=0

clean_database:
	- docker compose exec -T db mysql -u$(dbUser) -p$(dbPass) -e "DROP DATABASE IF EXISTS $(dbName)"
	- docker compose exec -T db mysql -u$(dbUser) -p$(dbPass) -e "CREATE DATABASE IF NOT EXISTS $(dbName)"

seed:
	- make clean_database
	- make migration_up
	- docker compose exec -T db mysql -u$(dbUser) -p$(dbPass) $(dbName) < config/database/seeds/seed.sql

test:
	- go test ./...

test_show_covers:
	- go test -coverprofile=cover.out ./...
	- go tool cover -html=cover.out

.PHONY: start start-swag swag
