GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOLIST=$(GOCMD) list
GOTOOL=$(GOCMD) tool

STATIK=statik

CMD_DIR=./cmd
BIN_DIR=./bin

COVERAGE_OUT=cover.out
COVERAGE_HTML=cover.html

BIN_NAME=manifest-updater

.PHONY: build
build:
	CGO_ENABLED=0 $(GOBUILD) -o $(BIN_NAME) -v ./

.PHONY: test
test:
	$(GOTEST) -coverprofile=$(COVERAGE_OUT) $$($(GOLIST) ./...)
	$(GOTOOL) cover -html=$(COVERAGE_OUT) -o $(COVERAGE_HTML)

.PHONY: clean
clean:
	rm -f $(COVERAGE_OUT)
	rm -f $(COVERAGE_HTML)
	rm -f $(BIN_NAME)/*
