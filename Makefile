GO_FILES = $(shell find . -type f -name '*.go' -not -path "*/mocks/*" -not -path "*.pb.go" -not -path "./proto/*" | tr "\n" " ")
GO_FILES_GOLINES = $(shell find . -type f -name '*.go' -not -path "*/mocks/*" -not -path "*.pb.go" -not -path "*test.go" -not -path "./proto/*" | tr "\n" " ")

fmt: ### Code formatting
	golines --base-formatter=gofmt --ignore-generated -m 130 -w $(GO_FILES_GOLINES)
	gofumpt -w $(GO_FILES)
	gci write \
		--section Standard \
		--section Default \
		--section "Prefix($(GO_INTERNAL_PKG))" \
		--section "Prefix($(GO_SDK))" \
		$(GO_FILES) > /dev/null 2>&1

grpc_gen:
	./protocgen.sh