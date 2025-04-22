package main

import (
	"context"
	"fmt"
	"math"
	"math/rand"

	"dagger/hello-dagger/internal/dagger"
)

type HelloDagger struct{}

// Publish the application container after building and testing it on-the-fly
func (m *HelloDagger) Publish(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
) (string, error) {
	_, err := m.Test(ctx, source)
	if err != nil {
		return "", err
	}
	return m.Build(source).
		Publish(ctx, fmt.Sprintf("ttl.sh/hello-dagger-%.0f", math.Floor(rand.Float64()*10000000))) //#nosec
}

// Build the application container
func (m *HelloDagger) Build(
	// +defaultPath="/"
	source *dagger.Directory,
) *dagger.Container {
	build := m.BuildEnv(source).
		WithExec([]string{"npm", "run", "build"}).
		Directory("./dist")
	return dag.Container().From("nginx:1.25-alpine").
		WithDirectory("/usr/share/nginx/html", build).
		WithExposedPort(80)
}

// Return the result of running unit tests
func (m *HelloDagger) Test(
	ctx context.Context,
	// +defaultPath="/"
	source *dagger.Directory,
) (string, error) {
	return m.BuildEnv(source).
		WithExec([]string{"npm", "run", "test:unit", "run"}).
		Stdout(ctx)
}

// Build a ready-to-use development environment
func (m *HelloDagger) BuildEnv(
	// +defaultPath="/"
	source *dagger.Directory,
) *dagger.Container {
	nodeCache := dag.CacheVolume("node")
	return dag.Container().
		From("node:21-slim").
		WithDirectory("/src", source).
		WithMountedCache("/root/.npm", nodeCache).
		WithWorkdir("/src").
		WithExec([]string{"npm", "install"})
}

// A coding agent for developing new features
func (m *HelloDagger) Develop(
	// Assignment to complete
	assignment string,
	// +defaultPath="/"
	source *dagger.Directory,
) *dagger.Directory {
	environment := dag.Env(dagger.EnvOpts{Privileged: true}).
		WithWorkspaceInput(
			"workspace",
			dag.Workspace(source),
			"the workspace with tools to edit and test code").
		WithWorkspaceOutput(
			"completed",
			"the workspace with the completed assignment")

	work := dag.LLM().
		WithEnv(environment).
		WithPrompt(`
			You are a develop on a Vue.js project.
			You will be given an assignment and the tools to complete the assignment.
			Do not stop until you have completed the assignment and the tests pass.
			Your assignment is:` + assignment)

	completed := work.
		Env().
		Output("completed").
		AsWorkspace()

	return completed.Source().WithoutDirectory("node_modules")
}
