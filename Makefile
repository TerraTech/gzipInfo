PROG          := $(shell basename ${PWD})
VERSION       := $(shell git describe --tags HEAD)
VERSION_SHORT := $(shell git describe --abbrev=0 --tags HEAD)
BUILD         := $(shell date '+%Y%m%d@%T')
LIB           := gzipInfo

D_VENDOR  := $(PWD)/vendor
F_AUTOGEN_LIB := pkg/gzipInfo/version_autogen.go
F_AUTOGEN_MAIN := ./version_autogen.go


all: $(PROG)

.PHONY: $(PROG)
$(PROG): vgen
	@make -s vendOn
	@go build -ldflags="-s -w"
	@make -s vendOff

.PHONY: clean
clean:
	@\rm $(PROG) $(F_AUTOGEN_MAIN) $(F_AUTOGEN_LIB) 2>/dev/null || true

.PHONY: fmt
fmt:
	@go fmt $(shell go list ./... | grep -v /vendor/)

.PHONY: release
release:
	@make -s vgen VERSION=$(VERSION_SHORT)
	@git add $(F_AUTOGEN_LIB)
	@git ci $(F_AUTOGEN_LIB) -m"release: $(VERSION_SHORT)"
	@git push
	@git tag --force $(VERSION_SHORT)
	@git push --force --tags

.PHONY: test
test:
	@go test

.PHONY: vgen
vgen: vendOn
	@test -f $(F_AUTOGEN_MAIN) && /bin/rm $(F_AUTOGEN_MAIN) 2>/dev/null || true
	@D_VENDOR=$(D_VENDOR) PROG=$(PROG) VERSION=$(VERSION) BUILD=$(BUILD) go generate $(dir $(F_AUTOGEN_MAIN))version.go
	@test -f $(F_AUTOGEN_LIB) && /bin/rm $(F_AUTOGEN_LIB) 2>/dev/null || true
	@D_VENDOR=$(D_VENDOR) PROG=$(PROG) VERSION=$(VERSION) BUILD=$(BUILD) LIB=$(LIB) go generate $(dir $(F_AUTOGEN_LIB))version.go
	@make -s vendOff

.PHONY: vendOff
vendOff:
	@echo vendOff
	ls -l vendor
	@\mv vendor/futurequest.net vendor/_futurequest.net 2>/dev/null || true
	ls -l vendor

.PHONY: vendOn
vendOn:
	@echo vendOn
	ls -l vendor
	@\mv vendor/_futurequest.net vendor/futurequest.net 2>/dev/null || true
	ls -l vendor
