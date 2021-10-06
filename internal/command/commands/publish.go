package commands

import (
	"path"
	"fmt"

	"github.com/dream11/d11-cli/internal/artifact"
	"github.com/dream11/d11-cli/internal/artifact/javaMaven"
	"github.com/dream11/d11-cli/pkg/shell"
	"github.com/dream11/d11-cli/pkg/file"
	"github.com/brownhash/golog"
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
	if len(args) != 1 {
		golog.Error(fmt.Errorf("`publish` requires exactly one argument `component dir path`, %d were given.", len(args)))
	}

	componentDir := args[0]
	artifactFilePath := path.Join(componentDir, artifactFileName)
	golog.Debug(fmt.Sprintf("Reading %s file from: %s", artifactFileName, artifactFilePath))

	artifact, err := artifact.ParseFile(artifactFilePath)
	if err != nil {
		golog.Error(err)
	}
	
	propertyFilePath := path.Join(componentDir, artifact.PropertyFile)
	golog.Debug(fmt.Sprintf("Reading %s file from: %s", artifact.PropertyFile, propertyFilePath))

	var artifactName string = ""
	var artifactDir string = ""
	var DockerFile string = ""

	if artifact.Flavour.Name == "javaMaven" {
		properties, err := javaMaven.ParseFile(propertyFilePath)
		if err != nil {
			golog.Error(err)
		}

		artifactName = fmt.Sprintf("%s-%s.jar", properties.Name, properties.Version)
		artifactDir = path.Join(componentDir, properties.Properties.ArtifactPath)

		DockerFile = javaMaven.DockerFile()
	} else {
		golog.Error(fmt.Sprintf("Unknown flavour `%s`", artifact.Flavour.Name))
	}

	// create container image for the component
	dockerfilePath := path.Join(artifactDir, "Dockerfile")
	golog.Warn("Creating Dockerfile")
	golog.Debug(fmt.Sprintf("Location: %s", artifactDir))
	err = file.Write(dockerfilePath, DockerFile, 0644)
	if err != nil {
		golog.Error(err)
	}
	golog.Success("Dockerfile generated successfully")

	golog.Warn(fmt.Sprintf("Creating %s", artifactName))
	golog.Debug(fmt.Sprintf("Location: %s", artifactDir))

	artifactPath := path.Join(artifactDir, artifactName)

	golog.Info("Running Pre steps")
	for i:=0; i<len(artifact.Steps.Pre); i++ {
		exitCode := shell.Exec(fmt.Sprintf("cd %s && %s", componentDir, artifact.Steps.Pre[i]))
		if exitCode > 0 {
			return 1
		}
	}

	golog.Info("Running Build steps")
	for i:=0; i<len(artifact.Steps.Build); i++ {
		exitCode := shell.Exec(fmt.Sprintf("cd %s && %s", componentDir, artifact.Steps.Build[i]))
		if exitCode > 0 {
			return exitCode
		}
	}

	golog.Success("Artifact Generated!")
	golog.Debug(fmt.Sprintf("Artifact Path: %s", artifactPath))

	return 0
}

// Help should return an explanatory string, 
// that can explain the command
func (t *Publish) Help() string {
	return "use `publish <component dir path>` to publish the required artifacts for the component"
}

// Synopsis should return a breif helper text for the command
func (t *Publish) Synopsis() string {
	return "publish artifacts"
}