# echo-server

A minimal HTTP server that reflects every request back as JSON. Useful for debugging HTTP clients, ingress configuration, and proxy behaviour.

## Response

Every request to any path returns a JSON body:

```json
{
  "method": "POST",
  "path": "/foo/bar?q=1",
  "headers": {
    "Content-Type": "application/json",
    "User-Agent": "curl/8.1.2"
  },
  "body": "{\"hello\":\"world\"}"
}
```

## Run locally

```sh
go run .
# or with a custom port
PORT=9090 go run .
```

```sh
curl -X POST http://localhost:8080/test -d '{"hello":"world"}'
```

## Docker

```sh
docker build -t echo-server .
docker run -p 8080:8080 echo-server
```

## Deploy to Kubernetes

The Helm chart is in `helm/echo-server`. Update `image.repository` in `values.yaml` to match your registry, then:

```sh
helm install echo-server ./helm/echo-server \
  --set image.repository=ghcr.io/<your-org>/echo-server \
  --set image.tag=<tag>
```

To expose it via an ingress:

```sh
helm upgrade echo-server ./helm/echo-server \
  --set ingress.enabled=true \
  --set ingress.hosts[0].host=echo.example.com
```

## CI

Pushing to `main` or a `v*` tag triggers `.github/workflows/docker.yml`, which builds and pushes the image to GHCR using the built-in `GITHUB_TOKEN` — no additional secrets required.
