package main

import (
	"fmt"
	"os"
	"path"
	"log"

	"github.com/dream11/d11-cli/internal/artifact"
	"github.com/dream11/d11-cli/internal/artifact/javaMaven"
	"github.com/dream11/d11-cli/pkg/shell"
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

	var artifactName string = ""
	if artifact.Flavour.Name == "java-maven" {
		artifactName = fmt.Sprintf("%s-%s.jar", properties.Name, properties.Version)
	}

	artifactDir := path.Join(componentDir, properties.Properties.ArtifactPath)
	log.Println("Creating", artifactName, "at", artifactDir)

	artifactPath := path.Join(artifactDir, artifactName)

	log.Println("Running Pre steps")
	for i:=0; i<len(artifact.Steps.Pre); i++ {
		exitCode := shell.Exec(fmt.Sprintf("cd %s && %s", componentDir, artifact.Steps.Pre[i]))
		if exitCode > 0 {
			os.Exit(exitCode)
		}
	}

	log.Println("Running Build steps")
	for i:=0; i<len(artifact.Steps.Build); i++ {
		exitCode := shell.Exec(fmt.Sprintf("cd %s && %s", componentDir, artifact.Steps.Build[i]))
		if exitCode > 0 {
			os.Exit(exitCode)
		}
	}

	log.Println("Artifact generated:", artifactPath)
	

	// shell.Exec(fmt.Sprintf("cd %s && %s", componentDir, artifact.Steps.Build[0]))
}