# Dagger Demo Monorepo

This repo starts from `dagger/hello-dagger` and turns it into a small monorepo
that is useful for showing Dagger as native CI:

- `apps/web`: Vue + Vite frontend from the original starter
- `services/api`: Go HTTP API with unit tests and a runnable image
- `packages/contracts`: shared API contract placeholder
- `.dagger`: Go Dagger module with CI functions

## Demo Commands

```sh
# install Dagger if needed
curl -fsSL https://dl.dagger.io/dagger/install.sh | sh

# list the native CI surface
./bin/dagger functions

# run the whole graph locally
./bin/dagger call ci

# run it in Dagger Cloud and show the trace
./bin/dagger call ci --cloud

# scale independent work when demoing larger runs
./bin/dagger call ci --cloud --scale-out
```

Individual functions are also useful for tracing:

```sh
./bin/dagger call web-build export --path=./apps/web/dist
./bin/dagger call api-test
./bin/dagger call api-image export --path=./api.tar
```

## Local Backend

```sh
cd services/api
go test ./...
go run ./cmd/server
```

The API exposes `GET /healthz` and `GET /api/products`.
