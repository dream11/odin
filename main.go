package main

import (
	"os"
	"log"
	"path"

	"github.com/dream11/d11-cli/internal/artifact"
)

const (
	artifactFileName = "artifact.yml"
)

func main() {
	componentDir := os.Args[1]
	artifactFilePath := path.Join(componentDir, artifactFileName)

	artifact, err := artifact.ParseFile(artifactFilePath)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(artifact)
	log.Println("Generating artifact at:", artifact.ArtifactPath, "using:", artifact.BuildFile)
}