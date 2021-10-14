package commands

import (
	"fmt"
	"io/ioutil"
	"path"

	"gopkg.in/yaml.v3"

	"github.com/brownhash/golog"
	"github.com/dream11/d11-cli/api/component"
	"github.com/dream11/d11-cli/api/profile"
	"github.com/dream11/d11-cli/api/service"
	"github.com/dream11/d11-cli/pkg/dir"
	"github.com/dream11/d11-cli/pkg/shell"
	"github.com/dream11/d11-cli/d11cli"
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

	//-------------------------------------------------------------------------
	// API IMPLEMENTATION
	//-------------------------------------------------------------------------
	golog.Println(fmt.Sprintf("Fetching profile: %s@%s", args[0], args[1]))
	// fetch profile from playground using profile name and version
	// now unmarshal the profile into api/profile.Profile
	profile := profile.Profile{
		Name: "p1",
		Version: "1.0.0",
		Services: []profile.Service{
			{
				Name: "fantasy-tour",
				Version: "1.1.1",
			},
		},
	}
	// Throw error if profile not successfuly unmarshaled
	golog.Success(fmt.Sprintf("Profile %s@%s fetched successfully!", profile.Name, profile.Version))

	// on parsing the above retreived profile
	// we now have a list of service version
	for _, service := range profile.Services {
		golog.Println(fmt.Sprintf("Fetching service: %s@%s", service.Name, service.Version))
		// now fetch service details from playground
		// now unmarshal each service into api/service.Service
		// Throw error if services not successfuly unmarshaled
		golog.Success(fmt.Sprintf("Service %s@%s fetched successfully!", service.Name, service.Version))
	}
	
	services := service.Services{
		{
			Name: "fantasy-tour",
			Version: "1.1.1",
			Components: []service.Component{
				{
					Name: "bb-rds",
					Version: "1.1.1",
				},
				{
					Name: "fantasy-tour",
					Version: "1.1.1",
				},
				{
					Name: "fantasy-tour-aerospike",
					Version: "1.1.1",
				},
				{
					Name: "fantasy-tour-admin",
					Version: "1.1.1",
				},
				{
					Name: "fantasy-tour-admin-rds",
					Version: "1.1.1",
				},
				{
					Name: "fantasy-tour-admin-redis",
					Version: "1.1.1",
				},
			},
		},
	}

	// on parsing the above retreived services
	// we now have a list of component version
	for _, service := range services {
		for _, component := range service.Components {
			golog.Println(fmt.Sprintf("Fetching component: %s@%s", component.Name, component.Version))
			// now fetch component details from playground
			// now unmarshal each component to api/component.Component
			// Throw error if components not successfuly unmarshaled
			golog.Success(fmt.Sprintf("Component %s@%s fetched successfully!", component.Name, component.Version))
		}
	}
	
	components := component.Components{
		{
			Name: "bb-rds",
			Version: "1.1.1",
			Type: "datastore",
			Artifact: component.Artifact{
				Url: "",
				Version: "",
				Type: "",
			},
		},
		{
			Name: "fantasy-tour",
			Version: "1.1.1",
			Type: "application",
			Artifact: component.Artifact{
				Url: "",
				Version: "",
				Type: "",
			},
		},
		{
			Name: "fantasy-tour-aerospike",
			Version: "1.1.1",
			Type: "datastore",
			Artifact: component.Artifact{
				Url: "",
				Version: "",
				Type: "",
			},
		},
		{
			Name: "fantasy-tour-admin",
			Version: "1.1.1",
			Type: "application",
			Artifact: component.Artifact{
				Url: "",
				Version: "",
				Type: "",
			},
		},
		{
			Name: "fantasy-tour-admin-rds",
			Version: "1.1.1",
			Type: "datastore",
			Artifact: component.Artifact{
				Url: "",
				Version: "",
				Type: "",
			},
		},
		{
			Name: "fantasy-tour-admin-redis",
			Version: "1.1.1",
			Type: "datastore",
			Artifact: component.Artifact{
				Url: "",
				Version: "",
				Type: "",
			},
		},
	}

	//-------------------------------------------------------------------------

	//-------------------------------------------------------------------------
	// FILE GENERATION & DEPLOY
	//-------------------------------------------------------------------------
	workDir := d11cli.WorkDir.Location
	// generate required files on the required path
	/*
	workdir
	|__ profileName
		|__ profileVersion
			|__ serviceName
				|__ serviceVersion
					|__ componentName
						|__ componentVersion
							|__ (helm files)
	*/
	profileDir := path.Join(workDir, profile.Name, profile.Version)
	golog.Warn(fmt.Sprintf("Generating files for %s@%s", profile.Name, profile.Version))
	golog.Debug(fmt.Sprintf("Location: %s", profileDir))

	profileExists, err := dir.Exists(profileDir)
	if err != nil {
		golog.Error(err)
	}

	if profileExists {
		golog.Warn(fmt.Sprintf("Running profile %s on %s@%s", action, profile.Name, profile.Version))
		golog.Debug(fmt.Sprintf("Location: %s", profileDir))

		for _, service := range profile.Services {
			serviceDir := path.Join(profileDir, service.Name, service.Version)
			serviceExists, err := dir.Exists(serviceDir)
			if err != nil {
				golog.Error(err)
			}

			if serviceExists {
				golog.Warn(fmt.Sprintf("Running profile %s on %s@%s/%s@%s", action, profile.Name, profile.Version, service.Name, service.Version))
				serviceDetails, err := services.GetService(service.Name)
				if err != nil {
					golog.Error(err)
				}

				for _, component := range serviceDetails.Components {
					componentDir := path.Join(serviceDir, component.Name, component.Version)
					componentExists, err := dir.Exists(componentDir)
					if err != nil {
						golog.Error(err)
					}

					if componentExists {
						golog.Warn(fmt.Sprintf("Running profile %s on %s@%s/%s@%s/%s@%s", action, profile.Name, profile.Version, service.Name, service.Version, component.Name, component.Version))
						componentDetails, err := components.GetComponent(component.Name)
						if err != nil {
							golog.Error(err)
						}

						chart, err := parseHelmChart(path.Join(componentDir, "Chart.yaml"))
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
						
						var actionCommand string
						if n.Deploy {
							actionCommand = fmt.Sprintf("cd %s && helm upgrade --%s %s d11-helm-charts/%s -f %s -f %s -n %s", componentDir, action, componentDetails.Name, chart.Name, path.Join(componentDir, "values.yaml"), path.Join(componentDir, "values-stag.yaml"), args[2])
						} else if n.Destroy {
							actionCommand = fmt.Sprintf("cd %s && helm %s %s -n %s", componentDir, action, componentDetails.Name, args[2])
						}
						
						status = shell.Exec(actionCommand)
						if status > 0 {
							return status
						}

						golog.Success(fmt.Sprintf("%sed %s@%s", action, componentDetails.Name, componentDetails.Version))
					} else {
						golog.Error(fmt.Sprintf("Error while reading component, does not exists. %s", componentDir))
					}
				}
				
			} else {
				golog.Error(fmt.Sprintf("Error while reading service, does not exists. %s", serviceDir))
			}
		}
	} else {
		golog.Error(fmt.Sprintf("Error while reading profile, does not exists. %s", profileDir))
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
