# go-playground

Simple go backend with gin, gorm and postgreSQL

# Create new task using curl
```bash
curl -H 'Content-Type: application/json' \
      -d '{ "description":"new task", "completed":false}' \
      -X POST \
      http://localhost:8080/tasks
```


# build the docker image
```bash
docker buildx build --sbom=true --provenance=true -t ivancurkovic046/go-playground:0.0.1 .
```
