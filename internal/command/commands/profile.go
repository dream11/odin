package commands

import (
	"fmt"
	"io/ioutil"
	"path"

	"gopkg.in/yaml.v3"

	"github.com/brownhash/golog"
	"github.com/dream11/d11-cli/pkg/dir"
	"github.com/dream11/d11-cli/pkg/shell"
)

type Chart struct {
	Name    string    `yaml:"name"`
}

type Profile struct {
	Deploy     bool
	Destroy    bool
}

func (n *Profile) Run(args []string) int {
	action := "" // initiate empty action
	if n.Deploy {
		action = "install"
	} else if n.Destroy {
		action = "uninstall"
	}

	if action == "" {
		if len(args) != 1 {
			golog.Error(fmt.Errorf("`profile` requires exactly one argument `profile name`, %d were given.", len(args)))
		}

		golog.Success("Listing all envs") // TODO: convert this log to debug type
		return 0
	}

	if len(args) != 3 {
		golog.Error(fmt.Errorf("`profile %s` requires exactly three arguments `profile name, version, env name`, %d were given.", action, len(args)))
	}

	profilePath := "/Users/harshitsharma/Documents/Dream11/poc/helm/Services"
	services, err := dir.SubDirs(profilePath)
	if err != nil {
		golog.Error(err)
	}

	for _, service := range services {
		servicePath := path.Join(profilePath, service)
		components, err := dir.SubDirs(servicePath)
		if err != nil {
			golog.Error(err)
		}

		for _, component := range components {
			componentPath := path.Join(servicePath, component)
			isDir, err := dir.IsDir(componentPath)
			if err != nil {
				golog.Error(err)
			}

			if isDir {
				golog.Warn(fmt.Sprintf("Running profile %s for %s", action, component))
				golog.Debug(fmt.Sprintf("Running profile %s in %s", action, componentPath))
				chart, err := parseHelmChart(path.Join(componentPath, "Chart.yaml"))
				if err != nil {
					golog.Error(err)
				}

				addRepoCommand := "helm repo add d11-helm-charts https://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@raw.githubusercontent.com/dream11/d11-helm-charts/feat/redis-operator/"
				status := shell.Exec(addRepoCommand)
				if status > 0 {
					return status
				}

				repoUpdateCommand := "helm repo update"
				status = shell.Exec(repoUpdateCommand)
				if status > 0 {
					return status
				}

				actionCommand := fmt.Sprintf("cd %s && helm upgrade --%s %s d11-helm-charts/%s -f %s -f %s -n %s", componentPath, action, component, chart.Name, path.Join(componentPath, "values.yaml"), path.Join(componentPath, "values-stag.yaml"), args[2])
				status = shell.Exec(actionCommand)
				if status > 0 {
					return status
				} 
			}
		}
	}

	golog.Success(fmt.Sprintf("Profile/%s@%s %sed in %s", args[0], args[1], action, args[2]))
	return 0
}

func (n *Profile) Help() string {
	if n.Deploy {
		return "use `profile deploy <profile-name> <version> <env-name>` to deploy the provided profile in the provided env"
	} else if n.Destroy {
		return "use `profile destroy <profile-name> <version> <env-name>` to destroy the provided profile in the provided env"
	}

	return "use `profile <name>` to list the created versions for the mentioned profile"
}

func (n *Profile) Synopsis() string {
	if n.Deploy {
		return "deploy the profile"
	} else if n.Destroy {
		return "destroy the deployed profile"
	}
	
	return "list profile versions"
}


// parse helm chart for chart properties
func parseHelmChart(filePath string) (Chart, error) {
	var chart Chart

	yFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return chart, err
	}

	err = yaml.Unmarshal(yFile, &chart)

	return chart, err
}
