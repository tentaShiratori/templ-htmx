
gen:
	@go tool templ generate
.PHONY: gen

fmt:
	@go fmt ./...
	@go tool templ fmt .
.PHONY: fmt

css:
	@pnpm run build-css-prod
.PHONY: css

css-watch:
	@pnpm run build-css
.PHONY: css-watch

install:
	@pnpm install
.PHONY: install

build: css gen
	@go build -o tmp/main.exe cmd/main.go
.PHONY: build

run:
	@go tool air
.PHONY: run

dev: css-watch
.PHONY: dev