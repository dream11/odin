package docker

import (
	"context"
	"time"
	"io"
	"errors"
	"bufio"
	"encoding/json"
	
	"github.com/brownhash/golog"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

func BuildImage(dockerfile, dockerfileDir string, tags []string, buildArgs map[string]*string) error {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	tar, err := archive.TarWithOptions(dockerfileDir, &archive.TarOptions{})
	if err != nil {
		return err
	}

	// For details on build options
	// https://pkg.go.dev/github.com/docker/docker@v20.10.9+incompatible/api/types#ImageBuildOptions
	opts := types.ImageBuildOptions{
		Dockerfile: dockerfile,
		Tags:       tags,
		Remove:     true,
		BuildArgs:  buildArgs,
	}
	res, err := dockerClient.ImageBuild(ctx, tar, opts)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	err = print(res.Body)
	if err != nil {
		golog.Error(err)
	}

	return nil
}

type ErrorLine struct {
	Error       string      `json:"error"`
	ErrorDetail ErrorDetail `json:"errorDetail"`
}

type ErrorDetail struct {
	Message string `json:"message"`
}

func print(rd io.Reader) error {
	var lastLine string

	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		lastLine = scanner.Text()
		golog.Println(scanner.Text())
	}

	errLine := &ErrorLine{}
	json.Unmarshal([]byte(lastLine), errLine)
	if errLine.Error != "" {
		return errors.New(errLine.Error)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}