start:
	docker-compose -f ./docker/compose.yaml up --build -d

stop: 
	docker-compose -f ./docker/compose.yaml down 

test:
	go test ./...

mocks:
	go generate ./...
