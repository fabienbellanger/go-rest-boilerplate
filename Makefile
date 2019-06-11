# Go parameters
GOCMD=go
GOINSTALL=$(GOCMD) install
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=goRestBoilerplate
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build

install:
	$(GOINSTALL) ./...

update:
	$(GOGET) -u && $(GOMOD) tidy

updateAll:
	$(GOGET) -u all && $(GOMOD) tidy

runApi:
	$(GORUN) main.go api
api: install runApi

runWebSocket:
	$(GORUN) -race main.go websocket
ws: install runWebSocket

runApiRace:
	$(GORUN) -race main.go api
apiRace: install runApiRace

runDbInit:
	$(GORUN) main.go db --init
dbInit: install runDbInit

runDbDump:
	$(GORUN) main.go db --dump
dbDump: install runDbDump

runLogsRotation:
	$(GORUN) main.go log
log: install runLogsRotation

runMakeMigration:
	$(GORUN) main.go make:migration
make-migration: install runMakeMigration

runMigrate:
	$(GORUN) main.go migrate --force
migrate: install runMigrate

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v

test: 
	$(GOTEST) -cover ./...
	
bench: 
	$(GOTEST) -bench=. ./...

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run-prod:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
