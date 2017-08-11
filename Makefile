GO_PROJECT_NAME := vaahan

define info_echo
  @tput setaf 4
  @echo $(1)
  @tput sgr0
endef

# Go rules
go_format:
	$(call info_echo,"\n....Formatting $(GO_PROJECT_NAME)'s go files....")
	gofmt -w $(GOPATH)/src/$(GO_PROJECT_NAME)

go_prep_install:
	$(call info_echo,"\n....Preparing installation environment for $(GO_PROJECT_NAME)....")
	mkdir -p $(GOPATH)/bin $(GOPATH)/pkg
	go get github.com/cespare/reflex

go_dep_install:
  $(call info_echo,"\n....Installing dependencies for $(GO_PROJECT_NAME)....")
  go get ./...

go_install:
  $(call info_echo,"\n....Compiling $(GO_PROJECT_NAME)....")
  go install $(GO_PROJECT_NAME)

go_test:
  $(call info_echo,"\n....Running tests for $(GO_PROJECT_NAME)....")
  go test ./src/$(GO_PROJECT_NAME)/...

go_run:
  $(call info_echo,"\n....Running $(GO_PROJECT_NAME)....")
  $(GOPATH)/bin/$(GO_PROJECT_NAME)


# Project rules
install:
	$(MAKE) go_prep_install
	$(MAKE) go_dep_install
	$(MAKE) go_install

run:
ifeq ($(CODE_ENV), dev)
	reflex -s -g 'src/$(GO_PROJECT_NAME)/**/*.go' make restart
else
	$(MAKE) go_run
endif

restart:
	@$(MAKE) go_format
	@$(MAKE) go_install
	@$(MAKE) go_test
	@$(MAKE) go_run

.PHONY: go_format go_prep_install go_dep_install go_install go_run install run restart
