GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test

GOLINTCMD=golangci-lint
GOLINT=$(GOLINTCMD) run

TARGET_PKG=./cmd/linebot
OUTFILE=linebot

.PHONY: all
all: test build

.PHONY: build
build:
	$(GOBUILD) -o $(OUTFILE) $(TARGET_PKG)

dev:
	$(GORUN) $(TARGET_PKG)/main.go

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: lint
lint:
	$(GOLINT) --fast

.PHONY: clean
clean:
	$(RM) $(OUTFILE)
