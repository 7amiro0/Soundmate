up:
	docker-compose -f ./deployments/docker-compose.yaml up --build

down:
	docker-compose -f ./deployments/docker-compose.yaml down