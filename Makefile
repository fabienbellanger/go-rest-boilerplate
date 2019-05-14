# Go parameters
GOCMD=go
GOINSTALL=$(GOCMD) install
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=goRestBoilerplate
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build

install:
	$(GOINSTALL) ./...

update:
	$(GOGET) -u

updateAll:
	$(GOGET) -u all

runApi:
	$(GORUN) main.go serve
serve: install runApi

runApiRace:
	$(GORUN) -race main.go serve
serveRace: install runApiRace

runDbInit:
	$(GORUN) main.go db --init
dbInit: install runDbInit

runDbDump:
	$(GORUN) main.go db --dump
dbDump: install runDbDump

runLogsRotation:
	$(GORUN) main.go log
log: install runLogsRotation

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v

test: 
	$(GOTEST) -cover ./...

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run-prod:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
