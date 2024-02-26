

Create new task using curl
```curl -H 'Content-Type: application/json' \
      -d '{ "description":"new task", "completed":false}' \
      -X POST \
      http://localhost:8080/tasks`
