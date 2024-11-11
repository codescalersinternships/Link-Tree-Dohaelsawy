# :link: Link Tree ~ Your Links, Your Story
platform that allows users to share social media profiles, or other important links defined in a single place

## Technology Stack:
- Go language
- Postgres DB 
- Vue framework
- Cypress
- Docker & docker-compose
- Kubernetes
- Helm
- AWS S3 

## Project structure:
```
.
├── backend
├── frontend
├── deployments
├── docker-compose.yaml
├── makefile
└── README.md
```
## Get Started:
- 1. get source code
  ```golang
  go get "github.com/codescalersinternships/Link-Tree-Dohaelsawy"
  ```
- 2. set up environment variables, make sure to create `.env` file in the root.
  ```golang
      DB_HOST=*****
      DB_USER=*****
      DB_PASSWORD=*****
      DB_NAME=*****
      DB_PORT=*****
      DB_TEST_HOST=*****
      DB_TEST_USER=*****
      DB_TEST_PASSWORD=*****
      DB_TEST_NAME=*****
      DB_TEST_PORT=*****
      PORT=*****
      JWT_SECRET=*****
      TOKEN_HOUR_LIFESPAN=*****
      BASE_URL=*****
      LINK_TREE_URL=*****
      AWS_REGION=*****
      AWS_ACCESS_KEY_ID=*****
      AWS_SECRET_ACCESS_KEY=*****
      ALLOW_ORIGIN=*****
  ```
- 3. build your local environment using docker-compose
  ```
    docker-compose up
  ```
> [!NOTE]
> For more information about each section please check Readme file in each subdirectory
