NAME=users
DEV_COMPOSE_FLAGS=-f deployments/docker-compose.yml -p dev

env_up: env_down
	docker-compose $(DEV_COMPOSE_FLAGS) pull
	docker-compose $(DEV_COMPOSE_FLAGS) build ${BUILD_ARGS} $(NAME)
	docker-compose $(DEV_COMPOSE_FLAGS) up -d $(NAME)

env_down:
	docker-compose $(DEV_COMPOSE_FLAGS) down --volumes