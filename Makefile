branch = $(shell git rev-parse --abbrev-ref HEAD)
tag = $(shell git describe --abbrev=0 --tags)
rev = $(shell git rev-parse --short HEAD)

pkgs = $(shell go list ./... | grep -v /vendor/)
project_src = $(shell find . -name "*.go" | grep -v /vendor/ | grep -v _test.go)

bin_prefix = "vendor/.bin"

.PHONY: build install \
	install-tools lint-only test-only test \
	update-deps clean man release

cmd/cati/cati: $(project_src)
	go build -race -o $@ \
		-ldflags "-X github.com/roshi619/cati/internal/command.Version=$(branch)-$(rev)" \
		github.com/roshi619/cati/cmd/cati
build: cmd/cati/cati
install:
	go install \
		-ldflags "-X github.com/roshi619/cati/internal/command.Version=$(branch)-$(rev)" \
		github.com/roshi619/cati/cmd/cati

vendor/.bin/dep: $(shell find vendor/github.com/golang/dep -name "*.go")
	mkdir -p $(@D)
	go build -o $@ ./vendor/github.com/golang/dep/cmd/dep
vendor/.bin/golint: $(shell find vendor/github.com/golang/lint -name "*.go")
	mkdir -p $(@D)
	go build -o $@ ./vendor/github.com/golang/lint/golint
vendor/.bin/megacheck: $(shell find vendor/honnef.co/go/tools -name "*.go")
	mkdir -p $(@D)
	go build -o $@ ./vendor/honnef.co/go/tools/cmd/megacheck
install-tools: vendor/.bin/dep vendor/.bin/golint vendor/.bin/megacheck
lint-only: vendor/.bin/golint vendor/.bin/megacheck
	$(bin_prefix)/golint -set_exit_status $(pkgs)
	$(bin_prefix)/megacheck $(pkgs)
	go vet $(pkgs)
test-only:
	go test -v -cover -race $(pkgs)
test: lint-only test-only

update-deps: vendor/.bin/dep
	$(bin_prefix)/dep ensure
	$(bin_prefix)/dep ensure -update
	$(bin_prefix)/dep prune
clean:
	go clean
	rm -f cmd/cati/cati
	rm -rf vendor/.bin
	rm -rf release/
	git clean -x -f -d
	git remote prune origin

docs/man/cati.1: docs/man/cati.1.md
	pandoc -s -t man $< -o $@
docs/man/cati.yaml.5: docs/man/cati.yaml.5.md
	pandoc -s -t man $< -o $@
man: docs/man/cati.1 docs/man/cati.yaml.5

release/cati$(tag).linux-amd64.tar.gz: $(shell find . -name "*.go")
	mkdir -p $(@D)
	GOOS=linux GOARCH=amd64 \
		go build \
		-ldflags "-s -w -X github.com/roshi619/cati/internal/command.Version=$(tag)" \
		github.com/roshi619/cati/cmd/cati
	tar -czf release/cati$(tag).linux-amd64.tar.gz cati
	rm -f cati
release/cati$(tag).darwin-amd64.tar.gz: $(shell find . -name "*.go")
	mkdir -p $(@D)
	GOOS=darwin GOARCH=amd64 \
		go build \
		-ldflags "-s -w -X github.com/roshi619/cati/internal/command.Version=$(tag)" \
		github.com/roshi619/cati/cmd/cati
	tar -czf release/cati$(tag).darwin-amd64.tar.gz cati
	rm -f cati
release/cati$(tag).windows-amd64.tar.gz: $(shell find . -name "*.go")
	mkdir -p $(@D)
	GOOS=windows GOARCH=amd64 \
		go build \
		-ldflags "-s -w -X github.com/roshi619/cati/internal/command.Version=$(tag)" \
		github.com/roshi619/cati/cmd/cati
	tar -czf release/cati$(tag).windows-amd64.tar.gz cati.exe
	rm -f cati.exe
release: release/cati$(tag).linux-amd64.tar.gz release/cati$(tag).darwin-amd64.tar.gz \
	release/cati$(tag).windows-amd64.tar.gz
