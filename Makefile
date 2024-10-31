.PHONY: setup
setup:
	npm install -D tailwindcss
	npm install -D @tailwindcss/typography
	npm i -D daisyui@latest
	npm install
	go install github.com/air-verse/air@latest
	go get entgo.io/ent/cmd/ent
	echo "Install kamal from https://kamal-deploy.org/docs/installation/"

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

# Run linting
.PHONY: lint
lint:
	golangci-lint run

# Run linting and fix
.PHONY: lint-fix
lint-fix:
	golangci-lint run --fix

# Run kamal init: Run once to initialize the config
.PHONY: kamal-init
kamal-init:
	kamal init

# Run kamal setup: Run once to setup the deploy server
.PHONY: kamal-setup
kamal-setup:
	kamal setup

# Run kamal deploy: Run subsequent times to deploy the app
.PHONY: kamal-deploy
kamal-deploy:
	kamal deploy