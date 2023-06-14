.PHONY: init
init:
	@go mod tidy

.PHONY: run
run:
	@go run tsop.go