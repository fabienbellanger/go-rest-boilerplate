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

runApi:
	$(GORUN) main.go web
serve: install runApi

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

deps:
	$(GOGET) -u github.com/spf13/cobra/cobra
	$(GOGET) -u github.com/labstack/echo/...
	$(GOGET) -u github.com/go-sql-driver/mysql
	$(GOGET) -u github.com/fatih/color
	$(GOGET) -u github.com/dgrijalva/jwt-go
