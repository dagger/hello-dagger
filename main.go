package main

import (
	"context"
)

type Example struct{}

const defaultNodeVersion = "16"

// The container image containing our production app, with built assets served by nginx.
func (m *Example) AppContainer(nodeVersion Optional[string]) *Container {
	return dag.Container().From("cgr.dev/chainguard/nginx:latest").
		WithDirectory("/usr/share/nginx/html", m.Build(nodeVersion)).
		WithExposedPort(8080)
}

// The app-container as a service for local testing
func (m *Example) Service(nodeVersion Optional[string]) *Service {
	return m.AppContainer(nodeVersion).AsService()
}

// A container w/ the built app and node toolchain for debugging
func (m *Example) Debug(nodeVersion Optional[string]) *Container {
	return m.buildBase(nodeVersion).Container().
		WithEntrypoint([]string{"sh"})
}

// The directory containing the built app
func (m *Example) Build(nodeVersion Optional[string]) *Directory {
	return m.buildBase(nodeVersion).Build().Container().Directory("./build")
}

// Run the app's tests
func (m *Example) Test(ctx context.Context, nodeVersion Optional[string]) (string, error) {
	return m.buildBase(nodeVersion).
		Run([]string{"test", "--", "--watchAll=false"}).
		Stderr(ctx)
}

// Publish the app-container (to ttl.sh)
func (m *Example) PublishContainer(ctx context.Context, nodeVersion Optional[string]) (string, error) {
	return dag.Ttlsh().Publish(ctx, m.AppContainer(nodeVersion))
}

func (m *Example) buildBase(nodeVersion Optional[string]) *Node {
	return dag.Node().
		WithVersion(nodeVersion.GetOr(defaultNodeVersion)).
		WithNpm().
		WithSource(dag.Host().Directory(".", HostDirectoryOpts{
			Exclude: []string{".git", "**/node_modules"},
		})).
		Install(nil)
}
