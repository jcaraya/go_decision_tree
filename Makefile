PROJECTNAME := $(shell basename "$(PWD)")

# Go related variables.
GOBASE=$(shell pwd)
GOPATH=$(GOBASE)/vendor:$(GOBASE)
GOBIN=$(GOBASE)/bin
GOMAIN=$(GOBASE)/cmd

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

# Execute the the application binary.
run: go-run

# Checks if there is any missing dependency
# and then proceeds to build the go binary.
build: go-get go-build

# init initializaes a new go modeule for the project
init: go-mod

# Clean the files created by the build rule.
clean: go-clean
	@rm -rf $(GOBIN)

# Go Related rules
go-mod:
	@echo "  >  Initializing module..."
	@go mod init

go-get:
	@echo "  >  Checking missing dependencies..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get $(GOMAIN)

go-build:
	@echo "  >  Building binary..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(PROJECTNAME) $(GOMAIN)

go-install:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go install $(GOMAIN)

go-run:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run $(GOMAIN)

go-clean:
	@echo "  >  Cleaning build cache"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean

.PHONY : clean