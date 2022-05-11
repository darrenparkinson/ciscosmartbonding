# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit:
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	# requires staticcheck.io from go install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

## coverage: create coverage report
.PHONY: coverage
coverage:
	@echo 'Creating coverage report...'
	go test -coverprofile=./coverage/profile.out ./...
	go tool cover -func=./coverage/profile.out -o ./coverage/coverage.txt
	go tool cover -html=./coverage/profile.out -o ./coverage/coverage.html
