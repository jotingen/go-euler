GOPATH   = $(HOME)/go

GOSRC = $(wildcard *.go)

.PHONY: $(GOSRC)
all: $(GOSRC)
$(GOSRC):
	go build -o $(patsubst %.go, bin/%, $@) $@ 
