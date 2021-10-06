package commands

import (
	"path"
	"fmt"
	"strings"

	"github.com/dream11/d11-cli/internal/artifact"
	"github.com/dream11/d11-cli/internal/artifact/javaMaven"
	"github.com/dream11/d11-cli/pkg/shell"
	"github.com/dream11/d11-cli/pkg/file"
	"github.com/dream11/d11-cli/pkg/docker"
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

	// Directory in which all required files reside
	componentDir := args[0]
	// File to refer for artifact properties
	artifactFilePath := path.Join(componentDir, artifactFileName)
	golog.Debug(fmt.Sprintf("Reading %s file from: %s", artifactFileName, artifactFilePath))

	artifact, err := artifact.ParseFile(artifactFilePath)
	if err != nil {
		golog.Error(err)
	}
	
	// File path of component property file, stated by artifact property file
	propertyFilePath := path.Join(componentDir, artifact.PropertyFile)
	golog.Debug(fmt.Sprintf("Reading %s file from: %s", artifact.PropertyFile, propertyFilePath))

	// component name
	var artifactName string = ""
	// directory to generate artifacts
	var artifactDir string = ""
	// file path of artifact build
	var artifactPath string = ""
	// versioned tag to attach with artifact
	var artifactTag string = ""
	// Dockerfile content
	var DockerFile string = ""
	// Build arguments to pass with docker image build
	var dockerBuildArgs map[string]*string = map[string]*string{}

	// perform javaMaven specific actions
	if artifact.Flavour.Name == "javaMaven" {
		properties, err := javaMaven.ParseFile(propertyFilePath)
		if err != nil {
			golog.Error(err)
		}

		artifactName = fmt.Sprintf("%s-%s.jar", properties.Name, properties.Version)
		artifactDir = path.Join(componentDir, properties.Properties.ArtifactPath)
		artifactPath = path.Join(artifactDir, artifactName)
		artifactTag = properties.Name + ":" + properties.Version
		jarPath := path.Join(properties.Properties.ArtifactPath, artifactName)
		
		runSteps := strings.Join(artifact.Steps.Run, " && ")
		
		// fetch dockerfile for javaMaven type artifact
		DockerFile = javaMaven.DockerFile()
		// generate docker build arguments
		dockerBuildArgs = map[string]*string{
			"MAVEN_VERSION": &artifact.Flavour.Version.Maven,
			"JAVA_VERSION": &artifact.Flavour.Version.Java,
			"JAR_PATH": &jarPath,
			"JAR_NAME": &artifactName,
			"PORT": &artifact.Port,
			"RUN_COMMAND": &runSteps,
		}
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

	golog.Success("Dockerfile Generated!")
	golog.Debug(fmt.Sprintf("Location: %s", dockerfilePath))

	golog.Warn(fmt.Sprintf("Creating %s", artifactName))
	golog.Debug(fmt.Sprintf("Location: %s", artifactDir))

	// run pre steps
	// this should include dependecy installations, structure generation etc..
	golog.Info("Running Pre steps")
	for i:=0; i<len(artifact.Steps.Pre); i++ {
		exitCode := shell.Exec(fmt.Sprintf("cd %s && %s", componentDir, artifact.Steps.Pre[i]))
		if exitCode > 0 {
			return exitCode
		}
	}

	// run buils steps
	// this should include the steps required to build the base artifact
	golog.Info("Running Build steps")
	for i:=0; i<len(artifact.Steps.Build); i++ {
		exitCode := shell.Exec(fmt.Sprintf("cd %s && %s", componentDir, artifact.Steps.Build[i]))
		if exitCode > 0 {
			return exitCode
		}
	}

	// Build image accepts the following parameters 
	// 1. Name of dockerfile
	// 2. directory from where the dockerfile will be picked
	//    a. This can be altered based on the files to copy (artifact dir / component dir)
	// 3. list of tags
	// 4. map of build arguments
	err = docker.BuildImage("Dockerfile", artifactDir, []string{artifactTag}, dockerBuildArgs)
	if err != nil {
		golog.Error(err)
	}

	golog.Success("Artifacts Generated!")
	golog.Debug(fmt.Sprintf("Artifact Path: %s", artifactPath))

	golog.Warn(fmt.Sprintf("Publishing: %s", artifactTag))
	// add publish steps

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