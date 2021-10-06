package commands

import (
	"log"
	"path"
	"fmt"

	"github.com/dream11/d11-cli/internal/artifact"
	"github.com/dream11/d11-cli/internal/artifact/javaMaven"
	"github.com/dream11/d11-cli/pkg/shell"
)

const (
	artifactFileName = "artifact.yml"
)

// --------------------------------------------------------
// Test Command
// --------------------------------------------------------
type Publish struct {}

// Run implements the actual functionality of the command
// and return exit codes based on success/failure of tasks performed
func (t *Publish) Run(args []string) int {
	componentDir := args[0]
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
			return 1
		}
	}

	log.Println("Running Build steps")
	for i:=0; i<len(artifact.Steps.Build); i++ {
		exitCode := shell.Exec(fmt.Sprintf("cd %s && %s", componentDir, artifact.Steps.Build[i]))
		if exitCode > 0 {
			return exitCode
		}
	}

	log.Println("Artifact generated:", artifactPath)

	return 0
}

// Help should return an explanatory string, 
// that can explain the command
func (t *Publish) Help() string {
	return "use `publish <version>` to publish the required artifacts for the component"
}

// Synopsis should return a breif helper text for the command
func (t *Publish) Synopsis() string {
	return "publish artifacts"
}