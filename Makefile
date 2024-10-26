.PHONY: setup
setup:
	npm install -D tailwindcss
	npm install -D @tailwindcss/typography
	npm i -D daisyui@latest
	npm install
	go install github.com/air-verse/air@latest
	go get entgo.io/ent/cmd/ent

# Generate Ent code
.PHONY: ent-gen
ent-gen:
	go generate ./ent

# Create a new Ent entity
.PHONY: ent-new
ent-new:
	go run entgo.io/ent/cmd/ent new $(name)

.PHONY: dev
dev:
	clear
	air -c .air.toml

.PHONY: dev-css
dev-css:
	npx tailwindcss -i tailwind-styles.css -o static/styles.css --postcss --watch

# Run all tests
.PHONY: test
test:
	go test -count=1 -p 1 ./...

# Check for direct dependency updates
.PHONY: check-updates
check-updates:
	go list -u -m -f '{{if not .Indirect}}{{.}}{{end}}' all | grep "\["
