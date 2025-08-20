## Steps to runs
1. Run docker containers
```
docker compose up -d
```

2. Run producer
```
go run producer.go
```

3. Run worker
```
go run worker.go
```

4. Test by sending data to the api endpoint
```
curl -Method POST -Uri "http://127.0.0.1:3000/api/v1/comments" -ContentType "application/json" -Body '{"text":"message 1"}'
```
