
provision:
	@echo "Provisioning GCounter Cluster"	
	bash scripts/provision.sh

gcounter-build:
	@echo "Building GCounter Docker Image"	
	docker build -t gcounter -f Dockerfile .

gcounter-run:
	@echo "Running Single GCounter Docker Container"
	docker run -p 8080:8080 -d gcounter

info:
	echo "GCounter Cluster Nodes"
	docker ps | grep 'gcounter'
	docker network ls | grep gcounter_network

clean:
	@echo "Cleaning GCounter Cluster"
	docker ps -a | awk '$$2 ~ /gcounter/ {print $$1}' | xargs -I {} docker rm -f {}
	docker network rm gcounter_network

build:
	@echo "Building GCounter Server"	
	go build -o bin/gcounter main.go

fmt:
	@echo "go fmt GCounter Server"	
	go fmt ./...

test:
	@echo "Testing GCounter"	
	go test -v --cover ./...