# Go parameters
GOCMD=go
GOINSTALL=$(GOCMD) install
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOTOOL=$(GOCMD) tool
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
	$(GORUN) -race main.go api
api: install runApi

runWebSocket:
	$(GORUN) -race main.go websocket
ws: install runWebSocket

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
	$(GORUN) main.go make-migration
make-migration: install runMakeMigration

runMigrate:
	$(GORUN) main.go migrate --force
migrate: install runMigrate

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v

test: 
	$(GOTEST) -cover ./...

testCoverCount: 
	$(GOTEST) -covermode=count -coverprofile=cover-count.out ./...

testCoverAtomic: 
	$(GOTEST) -covermode=atomic -coverprofile=cover-atomic.out ./...

coverCount:
	$(GOTOOL) cover -func=cover-count.out

coverAtomic:
	$(GOTOOL) cover -func=cover-atomic.out

htmlCoverCount:
	$(GOTOOL) cover -html=cover-count.out

htmlCoverAtomic:
	$(GOTOOL) cover -html=cover-atomic.out

runCoverCount: testCoverCount coverCount
runCoverAtomic: testCoverAtomic coverAtomic
viewCoverCount: testCoverCount htmlCoverCount
viewCoverCount: testCoverCount htmlCoverAtomic

bench: 
	$(GOTEST) -bench=. ./...

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run-prod:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)
