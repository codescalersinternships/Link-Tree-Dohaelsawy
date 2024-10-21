delete-none-image-docker:
	docker rmi $(docker images --filter "dangling=true" -q --no-trunc)

docker-network:
	docker network create linktreenet

docker-postgres-container:
	docker run --name linktreeDB --network linktreenet -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=adminpassword -d postgres:latest

createdb:
	docker exec -it postgres createdb --username=admin --owner=root linktreeDB