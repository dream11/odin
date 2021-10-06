package main

import (
	"fmt"
	"os"
	"path"
	"log"

	"github.com/dream11/d11-cli/internal/artifact"
	"github.com/dream11/d11-cli/internal/artifact/javaMaven"
	// "github.com/dream11/d11-cli/pkg/shell"
)

const (
	artifactFileName = "artifact.yml"
)

func main() {
	componentDir := os.Args[1]
	artifactFilePath := path.Join(componentDir, artifactFileName)
	log.Println("Reading", artifactFilePath)

	artifact, err := artifact.ParseFile(artifactFilePath)
	if err != nil {
		log.Fatal(err)
	}
	
	propertyFilePath := path.Join(componentDir, artifact.PropertyFile)
	log.Println("Reading", propertyFilePath)

	properties, err := javaMaven.ParseFile(propertyFilePath)
	if err != nil {
		log.Fatal(err)
	}

	artifactPath := path.Join(componentDir, properties.Properties.ArtifactPath)
	log.Println("An artifact will be generated at:", artifactPath)

	artifactName := fmt.Sprintf("%s-%s", properties.Name, properties.Version)
	log.Println("Artifact name:", artifactName)

	// shell.Exec(fmt.Sprintf("cd %s && %s", componentDir, artifact.Steps.Build[0]))
}