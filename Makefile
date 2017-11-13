include .env

run: go-run

compile: go-compile

dist: go-dist

deps: go-deps

init: db-init db-up go-deps

kill: db-down db-kill

clean: go-clean

go-run: go-compile
	@echo '# running server'
	@./$$SELFESTEEM_BIN

go-compile:
	@echo '# compiling files'
	@go build -o $$SELFESTEEM_BIN .

go-dist:
	@echo '# compiling Go in a linux-based static binary'
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $$SELFESTEEM_BIN .

go-deps:
	@echo '# ensuring Go dependencies'
	@dep ensure

go-clean:
	@echo '# removing Go binaries'
	@rm -f $$SELFESTEEM_BIN

db-run:
	@echo '# running psql'
	@psql -U $$SELFESTEEM_DB_USER $$SELFESTEEM_DB_NAME

db-init:
	@echo '# initializing database'
	@createuser $$SELFESTEEM_DB_USER
	@createdb -O $$SELFESTEEM_DB_USER $$SELFESTEEM_DB_NAME

db-up:
	@echo '# building up database tables'
	@psql -U $$SELFESTEEM_DB_USER $$SELFESTEEM_DB_NAME -f database/up.sql

db-down:
	@echo '# dismantling database tables'
	@psql -U $$SELFESTEEM_DB_USER $$SELFESTEEM_DB_NAME -f database/down.sql

db-kill:
	@echo '# killing database'
	@dropdb -i $$SELFESTEEM_DB_NAME
	@dropuser -i $$SELFESTEEM_DB_USER
