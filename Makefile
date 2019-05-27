PROG          := $(shell basename ${PWD})
VERSION       := $(shell git describe --tags HEAD)
VERSION_SHORT := $(shell git describe --abbrev=0 --tags HEAD)
BUILD         := $(shell date '+%Y%m%d@%T')
LIB           := gzipInfo

D_VENDOR	:= $(PWD)/vendor
F_AUTOGEN_LIB	:= pkg/gzipInfo/version_autogen.go
F_AUTOGEN_MAIN	:= ./version_autogen.go
P_GENVERSION	:= $(D_VENDOR)/futurequest.net/FQgolibs/tools/genVersion.go
PKG_FQVERSION	:= futurequest.net/FQgolibs/FQversion

all: $(PROG)

.PHONY: $(PROG)
$(PROG): vgen
	@go build -ldflags="-s -w"

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
vgen:
	@test -f $(F_AUTOGEN_MAIN) && /bin/rm $(F_AUTOGEN_MAIN) 2>/dev/null || true
	@P_GENVERSION=$(P_GENVERSION) D_VENDOR=$(D_VENDOR) PROG=$(PROG) VERSION=$(VERSION) BUILD=$(BUILD) IMPFQVERSION=$(PKG_FQVERSION) go generate $(dir $(F_AUTOGEN_MAIN))version.go
	@test -f $(F_AUTOGEN_LIB) && /bin/rm $(F_AUTOGEN_LIB) 2>/dev/null || true
	@P_GENVERSION=$(P_GENVERSION) D_VENDOR=$(D_VENDOR) PROG=$(PROG) VERSION=$(VERSION) BUILD=$(BUILD) IMPFQVERSION=$(PKG_FQVERSION) go generate $(dir $(F_AUTOGEN_LIB))version.go
