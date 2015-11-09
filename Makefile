
GOPATH  := $(GOPATH):$(shell pwd)/../../../../


define HG_ERROR

FATAL: you need mercurial (hg) to download datadust dependencies.
       Check README.md for details


endef

define GIT_ERROR

FATAL: you need git to download datadust dependencies.
       Check README.md for details
endef

define BZR_ERROR

FATAL: you need bazaar (bzr) to download datadust dependencies.
       Check README.md for details
endef

.PHONY: all check-path get hg git bzr get-code test

all: check-path get test

build: check-path get _go_test _datadust

# It does not support GOPATH with multiple paths.
check-path:
ifndef GOPATH
	@echo "FATAL: you must declare GOPATH environment variable, for more"
	@echo "       details, please check README.md file and/or"
	@echo "       http://golang.org/cmd/go/#GOPATH_environment_variable"
	@exit 1
endif
	@exit 0

get: hg git bzr get-code godep

hg:
	$(if $(shell hg), , $(error $(HG_ERROR)))

git:
	$(if $(shell git), , $(error $(GIT_ERROR)))

bzr:
	$(if $(shell bzr), , $(error $(BZR_ERROR)))

get-code:
	go get $(GO_EXTRAFLAGS) -u -d -t -insecure ./...

godep:
	go get $(GO_EXTRAFLAGS) github.com/tools/godep
	godep restore ./...

_go_test:
#	go clean $(GO_EXTRAFLAGS) ./...
#	go test $(GO_EXTRAFLAGS) ./...

_datadust:
	rm -f datadust
	go build $(GO_EXTRAFLAGS) -a github.com/scalegray/datadust

_datadustr:
	sudo ./datadust start
#	rm -f datadust

test: _go_test _datadust _datadustr

_install_deadcode: git
	go get $(GO_EXTRAFLAGS) github.com/remyoudompheng/go-misc/deadcode

deadcode: _install_deadcode
	@go list ./... | sed -e 's;github.com/scalegray/datadust/;;' | xargs deadcode

deadc0de: deadcode
