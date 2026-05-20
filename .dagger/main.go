// Demo CI functions for a small frontend + Go backend monorepo.

package main

import (
	"context"
	"dagger/demo-monorepo/internal/dagger"
)

type DemoMonorepo struct{}

// Run the full CI graph for the monorepo.
// +cache="never"
func (m *DemoMonorepo) Ci(ctx context.Context) (string, error) {
	if _, err := m.WebBuild(ctx); err != nil {
		return "", err
	}
	if _, err := m.ApiTest(ctx); err != nil {
		return "", err
	}
	if _, err := m.ApiImage().Sync(ctx); err != nil {
		return "", err
	}

	return "web build, api tests, and api image build passed", nil
}

// Build the Vue frontend and return its dist directory.
// +cache="never"
func (m *DemoMonorepo) WebBuild(ctx context.Context) (*dagger.Directory, error) {
	return m.web().
		WithExec([]string{"npm", "run", "build"}).
		Directory("dist").
		Sync(ctx)
}

// Run the Go backend test suite.
// +cache="never"
func (m *DemoMonorepo) ApiTest(ctx context.Context) (string, error) {
	return m.goEnv().
		WithExec([]string{"go", "test", "./..."}).
		Stdout(ctx)
}

// Build a runnable backend image.
// +cache="never"
func (m *DemoMonorepo) ApiImage() *dagger.Container {
	binary := m.goEnv().
		WithEnvVariable("CGO_ENABLED", "0").
		WithEnvVariable("GOOS", "linux").
		WithEnvVariable("GOARCH", "amd64").
		WithExec([]string{"go", "build", "-o", "/out/api", "./cmd/server"}).
		File("/out/api")

	return dag.Container().
		From("gcr.io/distroless/static-debian12:nonroot").
		WithFile("/usr/local/bin/api", binary).
		WithUser("nonroot:nonroot").
		WithExposedPort(8080).
		WithEntrypoint([]string{"/usr/local/bin/api"})
}

func (m *DemoMonorepo) web() *dagger.Container {
	return dag.Container().
		From("node:22-alpine").
		WithMountedCache("/root/.npm", dag.CacheVolume("demo-web-npm")).
		WithDirectory("/src", m.source().Directory("apps/web"), dagger.ContainerWithDirectoryOpts{
			Exclude: []string{"node_modules", "dist"},
		}).
		WithWorkdir("/src").
		WithExec([]string{"npm", "ci"})
}

func (m *DemoMonorepo) goEnv() *dagger.Container {
	return dag.Container().
		From("golang:1.24-alpine").
		WithMountedCache("/go/pkg/mod", dag.CacheVolume("demo-api-go-mod")).
		WithMountedCache("/root/.cache/go-build", dag.CacheVolume("demo-api-go-build")).
		WithDirectory("/src", m.source().Directory("services/api")).
		WithWorkdir("/src")
}

func (m *DemoMonorepo) source() *dagger.Directory {
	return dag.CurrentWorkspace().Directory(".", dagger.WorkspaceDirectoryOpts{
		Exclude: []string{".git", "bin", "apps/web/node_modules", "apps/web/dist"},
	})
}
