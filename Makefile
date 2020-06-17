GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GODOC=$(GOCMD)doc
COMPOSE=docker-compose
EXEC=$(COMPOSE) exec
BUILD=$(COMPOSE) build
UP=$(COMPOSE) up -d
LOGS=$(COMPOSE) logs
STOP=$(COMPOSE) stop
RM=$(COMPOSE) rm
DOWN=$(COMPOSE) down
GOLANG=$(EXEC) go
DB=$(EXEC) mysql
VUE=$(EXEC) vue
NGINX=$(EXEC) nginx

all: docker/up

docker/build: ## docker build
	$(BUILD)

docker/up: ## docker up
	$(UP)

docker/logs: ## docker logs
	$(LOGS)

docker/stop: ## docker stop
	$(STOP)

docker/rm: ## docker clean
	$(RM)

docker/down: ## docker down
	$(DOWN)

go/bash: ## go container bash
	$(GOLANG) bash

db/bash: ## db(MySQL) container bash
	$(DB) bash

vue/bash: ## vue container bash
	$(VUE) bash

nginx/bash: ## NGINX container bash
	$(NGINX) bash

doc: ## godoc http:6060
	$(GODOC) -http=:6060

help: ## Display this help screen
	@grep -E '^[a-zA-Z/_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'