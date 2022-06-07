NAME := app
BUILD_CMD ?= CGO_ENABLED=0 go build -o bin/${NAME} -ldflags '-v -w -s' ./cmd/${NAME}
DEV_CONFIG_PATH := ./configs/dev.yml
STAGE_CONFIG_PATH := ./configs/stage.yml
CONFIG_TEMPLATE_PATH := ./configs/template.yml

# Docker
DOCKER_APP_FILENAME ?= deployments/docker/Dockerfile
DOCKER_COMPOSE_FILE ?= deployments/docker-compose/docker-compose.yml
VERSION ?= v0.1.2

# sed
SECRET_KEY ?= "very-secret-key"
CONFIG_PATH ?= ./configs/new.yml

define sedi
    sed --version >/dev/null 2>&1 && sed -- $(1) > ${CONFIG_PATH} || sed "" $(1) > ${CONFIG_PATH}
endef

.PHONY: run
run: entgen css
	go run cmd/$(NAME)/main.go ${DEV_CONFIG_PATH}

.PHONY: stage
stage: entgen css
	go run cmd/$(NAME)/main.go ${STAGE_CONFIG_PATH}

.PHONY: up
up: 
	docker-compose -f ${DOCKER_COMPOSE_FILE} up --build --remove-orphans --always-recreate-deps --renew-anon-volumes

.PHONY: test
test: entgen
	go test ./... -cover

.PHONY: build
build:
	echo "building"
	${BUILD_CMD}
	echo "build done"

.PHONY: lint
lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.45.2
	./bin/golangci-lint run -v

.PHONY: entgen
entgen:
	cd ./internal/repository/entgo && go generate ./ent

.PHONY: jup
jup: 
	cd scripts && pipenv run jupyter notebook

.PHONY: css
css:
	cd web && npm run build

.PHONY: db
db:
	cd deployments/dev && docker-compose up -d --force-recreate --build --remove-orphans --always-recreate-deps --renew-anon-volumes

.PHONY: substitute_config_vars
substitute_config_vars:
	$(call sedi," \
		s|{{db_password}}|${DB_PASSWORD}|g;         \
		s|{{db_name}}|${DB_NAME}|g;                 \
		s|{{db_host}}|${DB_HOST}|g;                 \
		s|{{db_port}}|${DB_PORT}|g;                 \
		s|{{db_user}}|${DB_USER}|g;                 \
		s|{{secret_key}}|${SECRET_KEY}|g;           \
		s|{{telegram_to}}|${TELEGRAM_TO}|g;         \
		s|{{telegram_token}}|${TELEGRAM_TOKEN}|g;         \
		s|{{google_auth_client}}|${GOOGLE_AUTH_CLIENT}|g;   \
		s|{{google_auth_secret}}|${GOOGLE_AUTH_SECRET}|g;   \
		s|{{google_auth_callback}}|${GOOGLE_AUTH_CALLBACK}|g;   \
		s|{{github_auth_client}}|${GITHUB_AUTH_CLIENT}|g;   \
		s|{{github_auth_secret}}|${GITHUB_AUTH_SECRET}|g;   \
		s|{{github_auth_callback}}|${GITHUB_AUTH_CALLBACK}|g;   \
		s|{{twitter_auth_client}}|${TWITTER_AUTH_CLIENT}|g;   \
		s|{{twitter_auth_secret}}|${TWITTER_AUTH_SECRET}|g;   \
		s|{{twitter_auth_callback}}|${TWITTER_AUTH_CALLBACK}|g;   \
		s|{{version}}|${VERSION}|g;   \
		" ${CONFIG_TEMPLATE_PATH})
	cat ${CONFIG_PATH}