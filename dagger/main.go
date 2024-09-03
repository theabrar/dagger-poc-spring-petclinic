// A generated module for DaggerPoc functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/dagger-poc/internal/dagger"
	"fmt"
)

type DaggerPoc struct{}


func (m *DaggerPoc) Build(ctx context.Context, source *dagger.Directory) *dagger.File {
	return dag.
			Java().
			WithJdk("17").
			WithMaven("3.9.5").
			WithProject(source.WithoutDirectory("dagger")).
			Maven([]string{"package", "-DskipTests"}).
			File("target/spring-petclinic-3.3.0-SNAPSHOT.jar")
}

func (m *DaggerPoc) Publish(ctx context.Context, source *dagger.Directory, version string, registryAddress string, registryUsername string, registryPassword *dagger.Secret, imageName string) (string, error) {
	return dag.Container(dagger.ContainerOpts{Platform: "linux/amd64"}).
			From("eclipse-temurin:17-alpine").
			WithLabel("org.opencontainers.image.title", "dagger-poc").
			WithLabel("org.opencontainers.image.version", version).
			WithFile("spring-petclinic-3.3.0-SNAPSHOT.jar", m.Build(ctx, source)).
			WithEntrypoint([]string{"java", "-jar", "/app/spring-petclinic-3.3.0-SNAPSHOT.jar"}).
			WithRegistryAuth(registryAddress, registryUsername, registryPassword).
			Publish(ctx, fmt.Sprintf("%s/%s:%s", registryAddress, registryUsername, imageName))
}


// Returns a container that echoes whatever string argument is provided
// func (m *DaggerPoc) ContainerEcho(stringArg string) *dagger.Container {
// 	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
// }

// // Returns lines that match a pattern in the files of the provided Directory
// func (m *DaggerPoc) GrepDir(ctx context.Context, directoryArg *dagger.Directory, pattern string) (string, error) {
// 	return dag.Container().
// 		From("alpine:latest").
// 		WithMountedDirectory("/mnt", directoryArg).
// 		WithWorkdir("/mnt").
// 		WithExec([]string{"grep", "-R", pattern, "."}).
// 		Stdout(ctx)
// }
