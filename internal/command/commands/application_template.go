package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/dream11/odin/pkg/file"
	"gopkg.in/yaml.v3"
)

// ApplicationTemplate : Sample command declaration
type ApplicationTemplate command

type ApplicationSpec struct {
	Version string `yaml:"version"`
}

// Run implements the actual functionality of the command
func (a *ApplicationTemplate) Run(args []string) int {
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)

	serviceName := flagSet.String("name", "", "Name of the service")
	err := flagSet.Parse(args)
	if err != nil {
		a.Logger.Error("Unable to parse flags! " + err.Error())
		return 1
	}

	if a.Generate {
		emptyParameters := emptyParameters(map[string]string{"--name": *serviceName})

		if len(emptyParameters) == 0 {
			a.Logger.Info("Generating application template")
			var path = ".odin/" + *serviceName
			err := os.MkdirAll(path, 0766)
			if err != nil {
				a.Logger.Error("Unable to create directoy: " + path + "\n" + err.Error())
			}

			applicationSpec := ApplicationSpec{Version: "1.0.0-SNAPSHOT"}

			applicationSpecContent, err := yaml.Marshal(&applicationSpec)
			if err != nil {
				a.Logger.Error("Unable to parse content to yml " + err.Error())
			}

			// create and write data to files
			err = file.Write(path+"/start.sh", "", 0755)
			if err != nil {
				a.Logger.Error("Unable to create file `start.sh`." + err.Error())
			}
			err = file.Write(path+"/build.sh", "", 0755)
			if err != nil {
				a.Logger.Error("Unable to create file `build.sh`." + err.Error())
			}
			err = file.Write(path+"/pre-deploy.sh", "", 0755)
			if err != nil {
				a.Logger.Error("Unable to create file `pre-deploy.sh`." + err.Error())
			}
			err = file.Write(path+"/application-spec.yml", string(applicationSpecContent), 0755)
			if err != nil {
				a.Logger.Error("Unable to write file `application-spec.yml`." + err.Error())
			}
			a.Logger.Info("Template generated successfully.")
			return 0
		}
		a.Logger.Error(fmt.Sprintf("%s cannot be blank", emptyParameters))
		return 1
	}

	a.Logger.Error("Not a valid command")
	return 127
}

func (a *ApplicationTemplate) Help() string {
	if a.Generate {
		return commandHelper("generate", "application-template", "", []Options{
			{Flag: "--name", Description: "name of the service"},
		})
	}

	return defaultHelper()
}

func (a *ApplicationTemplate) Synopsis() string {
	if a.Generate {
		return "generate directory structure with files"
	}

	return defaultHelper()
}
