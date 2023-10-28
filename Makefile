# Variables declaration ---------------------------------------------------------------------------------------------- #
BUILD_TIMESTAMP := `date +"%a %d %b %Y %H:%M:%S"`
BUILD_VERSION   := `git describe --tags --abbrev=0`
FLAGS           := "-X 'main.buildTimestamp=$(BUILD_TIMESTAMP)' -X 'main.buildVersion=$(BUILD_VERSION)'"

# Go ----------------------------------------------------------------------------------------------------------------- #
.PHONY: go/run
go/run:
	go run -ldflags=$(FLAGS) -v cmd/main.go
