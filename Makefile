.PHONY: help
## help: show this helpful help message. Default target
help: Makefile
	@echo
	@echo " Available targets"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo

.PHONY: fmt
## fmt: Performs a terraform fmt
fmt:
	terraform fmt -diff -check .

.PHONY: e2etest
## e2etest: Runs the terratest end to end test for this module. See test/TESTING.md
e2etest:
	@cd test/e2e && go test