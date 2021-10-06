package docker

import (
	"context"
	"time"
	
	"github.com/brownhash/golog"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

func BuildImage(dockerfilePath string, tags []string, buildArgs map[string]*string) error {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	tar, err := archive.TarWithOptions("d11-cli-build/", &archive.TarOptions{})
	if err != nil {
		return err
	}

	// For details on build options
	// https://pkg.go.dev/github.com/docker/docker@v20.10.9+incompatible/api/types#ImageBuildOptions
	opts := types.ImageBuildOptions{
		Dockerfile: dockerfilePath,
		Tags:       tags,
		Remove:     true,
		BuildArgs:  buildArgs,
	}
	res, err := dockerClient.ImageBuild(ctx, tar, opts)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	golog.Println(res.Body)

	return nil
}