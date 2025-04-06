.PHONY: setup
setup:
	npm install tailwindcss @tailwindcss/cli
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
	npx tailwindcss -i input.css -o static/styles.css --watch

# Run all tests
.PHONY: test
test:
	go test -count=1 -p 1 ./...

.PHONY: build
build: setup
	npx tailwindcss -i input.css -o static/styles.css
	go build -o pagoda cmd/web/main.go

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

.PHONY: image
image:
	docker build -t CONTAINER_REGISTRY_USERNAME/IMAGE_NAME .

.PHONY: shell
shell:
	docker run -v $(PWD):/app -it CONTAINER_REGISTRY_USERNAME/IMAGE_NAME bash

.PHONY: deploy
deploy:
	echo "Override the following environment variables"
	echo "  - KAMAL_REGISTRY_PASSWORD=<your-docker-hub-password>"
	echo "  - PAGODA_HTTP_HOSTNAME=0.0.0.0"
	echo "  - PAGODA_APP_ENVIRONMENT=prod"
	echo "  - PAGODA_APP_ENCRYPTIONKEY=<new-encryption-key>"
	echo "  - PAGODA_MAIL_HOSTNAME=<your-mail-hostname>"
	echo "  - PAGODA_MAIL_USER=<your-mail-user>"
	echo "  - PAGODA_MAIL_PASSWORD=<your-mail-password>"
	echo "  - PAGODA_MAIL_FROMADDRESS=<your-mail-from-address>"
	echo "Next, run these commands"
	echo "  - rvm use 3.4.2"
	echo "  - kamal deploy"

.PHONY: stripe-mock
stripe-mock:
	docker run --rm -it -p 12111-12112:12111-12112 stripe/stripe-mock:latest

.PHONY: redis-up
redis-up:
	docker-compose up

.PHONY: redis-down
redis-down:
	docker-compose down
