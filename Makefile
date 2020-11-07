
# Makefile for soil

VERSION ?= 0.0.0

test:

test-prerequisites:

install-tools:

### TEST ####################################################################

test-soil:
	ginkgo -r
test-soil-watch:
	ginkgo watch
test: test-soil
.PHONY: test-soil
.PHONY: test

clean:
	rm -r bin/* dist/*
