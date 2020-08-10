PROJECTNAME := $(shell basename "$(PWD)")

# Go related variables.
GOBASE=$(shell pwd)
GOPATH=$(GOBASE)/vendor:$(GOBASE)
GOBIN=$(GOBASE)/bin
GOCMD=$(GOBASE)/cmd
GOPKG := $(shell find $(GOBASE)/pkg -maxdepth 1 -mindepth 1)

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

# Execute the the application binary.
run: go-run

# Builds the go binary.
build: go-build

# Compiles the library code stored in pkg.
build-pkg: go-build-pkg

# Initializaes a new go module for the project.
init: go-mod

# Clean the files created by the build rule.
clean: go-clean
	@rm -rf $(GOBIN)
	@rm -rf $(GOBASE)/vendor

# Go Related rules
go-mod:
	@echo "  >  Initializing module..."
	@go mod init

go-get:
	@echo "  >  Checking missing dependencies..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get $(GOCMD)

go-build:
	@echo "  >  Building binary..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(PROJECTNAME) $(GOCMD)

go-build-pkg:
	@echo "  >  Building pkg..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build $(LDFLAGS) $(GOPKG)

go-install:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go install $(GOCMD)

go-run:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run $(GOCMD)

go-clean:
	@echo "  >  Cleaning build cache"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean

.PHONY : clean