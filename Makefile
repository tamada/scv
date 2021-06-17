GO := go
NAME := scv
VERSION := 1.0.0
DIST := $(NAME)-$(VERSION)

all: test build

setup:
	git submodule update --init

test: setup
	$(GO) test -covermode=count -coverprofile=coverage.out $$(go list ./...)

define __create_dist
	mkdir -p dist/$(1)_$(2)/$(DIST)
	GOOS=$1 GOARCH=$2 go build -o dist/$(1)_$(2)/$(DIST)/$(NAME)$(3) main.go args.go printer.go input.go
	cp -r README.md LICENSE completions dist/$(1)_$(2)/$(DIST)
	tar cfz dist/$(DIST)_$(1)_$(2).tar.gz -C dist/$(1)_$(2) $(DIST)
	echo "Done $(1)_$(2)"
endef

dist: all
	@$(call __create_dist,darwin,amd64,)
	@$(call __create_dist,darwin,arm64,)
	@$(call __create_dist,windows,amd64,.exe)
	@$(call __create_dist,linux,amd64,)

build: main.go args.go printer.go input.go vector
	go build -o $(NAME) -v main.go args.go printer.go input.go

clean:
	@rm -f scv *~
