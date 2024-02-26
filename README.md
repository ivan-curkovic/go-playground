# go-playground

Simple go backend with gin, gorm and postres

# Create new task using curl
`curl -H 'Content-Type: application/json' \
      -d '{ "description":"new task", "completed":false}' \
      -X POST \
      http://localhost:8080/tasks`


# build the docker image
`docker buildx build --sbom=true --provenance=true -t ivancurkovic046/go-playground:0.0.1 .`
