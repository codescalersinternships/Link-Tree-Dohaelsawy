delete-none-image-docker:
	docker rmi $(docker images --filter "dangling=true" -q --no-trunc)

docker-network:
	docker network create linktreenet

docker-postgres-container:
	docker run --name linktreeDB --network linktreenet -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=adminpassword -d postgres:latest

createdb:
	docker exec -it postgres createdb --username=admin --owner=root linktreeDB

k8s-db-connect:
	kubectl exec -it postgres-5bcd4d868b-6g5hb -- psql -h localhost -U admin --password -p 5432 linktreedb
format:
	gofmt -w .
test:
	go test -v ./...
lint:
	golangci-lint run ./...
docker-build-img-front:
	cd frontend && docker build . -f Dockerfile.vue.prod -t dohaelsawi/linktree-ui