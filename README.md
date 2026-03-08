# Wallace

## Setting up

```sh
# with gnu make
make up
# with docker
docker compose up --build -d
```

## Running app

```sh
# with gnu make
make container && go run ./cmd/api/main.go
# with docker
docker exec -it wallace-app sh -c "go run ./cmd/api/main.go"
```

## Testing locally

```sh
curl -i http://0.0.0.0:8000
curl -i http://0.0.0.0:8000/example/messages
curl -i http://0.0.0.0:8000/healthz
```
