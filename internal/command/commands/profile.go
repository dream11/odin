package commands

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v3"

	"github.com/dream11/odin/api/component"
	"github.com/dream11/odin/api/profile"
	"github.com/dream11/odin/api/service"
	"github.com/dream11/odin/internal/commandline"
	"github.com/dream11/odin/odin"
	"github.com/dream11/odin/pkg/dir"
	"github.com/dream11/odin/pkg/shell"
)

type Chart struct {
	Name string `yaml:"name"`
}

type Profile command

func (p *Profile) Run(args []string) int {
	// Define flagset
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError)

	// create flags
	fileName := flagSet.String("file", "profile.yaml", "file name of profile yaml")
	profileName := flagSet.String("profile", "demo", "name of profile to be used")
	profileVersion := flagSet.String("version", "0.0.1", "version of profile to be used")
	envName := flagSet.String("env", "demo", "name of environment to deploy the profile in")

	// positional parse flags from [3:]
	flagSet.Parse(os.Args[3:])

	if p.Create {
		commandline.Interface.Info(fmt.Sprintf("Creating profile %s@%s", *profileName, *profileVersion))
		commandline.Interface.Info(*fileName)
		// TODO: read profile yaml from given file and call profile create api
		return 0
	}

	if p.Delete {
		commandline.Interface.Info(fmt.Sprintf("Deleting profile %s@%s", *profileName, *profileVersion))
		// TODO: take profile name and version and call profile delete api
		return 0
	}

	if p.List {
		commandline.Interface.Info("Listing all profiles")
		// TODO: call profiles api and display all profiles
		return 0
	}

	if p.Describe {
		commandline.Interface.Info(fmt.Sprintf("Describing %s", *envName))
		// TODO: call profile api and display all listed versions
		return 0
	}

	// initiate fetching data only when
	// deploy or destroy of a profile has to be done
	if p.Deploy || p.Destroy {
		//-------------------------------------------------------------------------
		// API IMPLEMENTATION
		//-------------------------------------------------------------------------
		commandline.Interface.Info(fmt.Sprintf("Fetching profile: %s@%s", *profileName, *profileVersion))
		// fetch profile from playground using profile name and version
		// now unmarshal the profile into api/profile.Profile
		profile := profile.Profile{
			Name:    "p1",
			Version: "1.0.0",
			Services: []profile.Service{
				{
					Name:    "fantasy-tour",
					Version: "1.1.1",
				},
			},
		}

		// Throw error if profile not successfuly unmarshaled
		commandline.Interface.Info(fmt.Sprintf("Profile %s@%s fetched successfully!", profile.Name, profile.Version))

		// on parsing the above retreived profile
		// we now have a list of service version
		for _, service := range profile.Services {
			commandline.Interface.Info(fmt.Sprintf("Fetching service: %s@%s", service.Name, service.Version))
			// now fetch service details from playground
			// now unmarshal each service into api/service.Service
			// Throw error if services not successfuly unmarshaled
			commandline.Interface.Info(fmt.Sprintf("Service %s@%s fetched successfully!", service.Name, service.Version))
		}

		services := service.Services{
			{
				Name:    "fantasy-tour",
				Version: "1.1.1",
				Components: []service.Component{
					{
						Name:    "bb-rds",
						Version: "1.1.1",
					},
					{
						Name:    "fantasy-tour",
						Version: "1.1.1",
					},
					{
						Name:    "fantasy-tour-aerospike",
						Version: "1.1.1",
					},
					{
						Name:    "fantasy-tour-admin",
						Version: "1.1.1",
					},
					{
						Name:    "fantasy-tour-admin-rds",
						Version: "1.1.1",
					},
					{
						Name:    "fantasy-tour-admin-redis",
						Version: "1.1.1",
					},
				},
			},
		}

		// on parsing the above retreived services
		// we now have a list of component version
		for _, service := range services {
			for _, component := range service.Components {
				commandline.Interface.Info(fmt.Sprintf("Fetching component: %s@%s", component.Name, component.Version))
				// now fetch component details from playground
				// now unmarshal each component to api/component.Component
				// Throw error if components not successfuly unmarshaled
				commandline.Interface.Info(fmt.Sprintf("Component %s@%s fetched successfully!", component.Name, component.Version))
			}
		}

		components := component.Components{
			{
				Name:    "bb-rds",
				Version: "1.1.1",
				Type:    "datastore",
				Artifact: component.Artifact{
					Url:     "",
					Version: "",
					Type:    "",
				},
			},
			{
				Name:    "fantasy-tour",
				Version: "1.1.1",
				Type:    "application",
				Artifact: component.Artifact{
					Url:     "",
					Version: "",
					Type:    "",
				},
			},
			{
				Name:    "fantasy-tour-aerospike",
				Version: "1.1.1",
				Type:    "datastore",
				Artifact: component.Artifact{
					Url:     "",
					Version: "",
					Type:    "",
				},
			},
			{
				Name:    "fantasy-tour-admin",
				Version: "1.1.1",
				Type:    "application",
				Artifact: component.Artifact{
					Url:     "",
					Version: "",
					Type:    "",
				},
			},
			{
				Name:    "fantasy-tour-admin-rds",
				Version: "1.1.1",
				Type:    "datastore",
				Artifact: component.Artifact{
					Url:     "",
					Version: "",
					Type:    "",
				},
			},
			{
				Name:    "fantasy-tour-admin-redis",
				Version: "1.1.1",
				Type:    "datastore",
				Artifact: component.Artifact{
					Url:     "",
					Version: "",
					Type:    "",
				},
			},
		}

		//-------------------------------------------------------------------------

		//-------------------------------------------------------------------------
		// FILE GENERATION & DEPLOY
		//-------------------------------------------------------------------------
		workDir := odin.WorkDir.Location
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
		commandline.Interface.Warn(fmt.Sprintf("Generating files for %s@%s", profile.Name, profile.Version))
		commandline.Interface.Info(fmt.Sprintf("Location: %s", profileDir))

		profileExists, err := dir.Exists(profileDir)
		if err != nil {
			commandline.Interface.Error(err.Error())
			return 1
		}

		if profileExists {
			// commandline.Interface.Warn(fmt.Sprintf("Running profile %s on %s@%s", action, profile.Name, profile.Version))
			commandline.Interface.Info(fmt.Sprintf("Location: %s", profileDir))

			for _, service := range profile.Services {
				serviceDir := path.Join(profileDir, service.Name, service.Version)
				serviceExists, err := dir.Exists(serviceDir)
				if err != nil {
					commandline.Interface.Error(err.Error())
					return 1
				}

				if serviceExists {
					// commandline.Interface.Warn(fmt.Sprintf("Running profile %s on %s@%s/%s@%s", action, profile.Name, profile.Version, service.Name, service.Version))
					serviceDetails, err := services.GetService(service.Name)
					if err != nil {
						commandline.Interface.Error(err.Error())
						return 1
					}

					for _, component := range serviceDetails.Components {
						componentDir := path.Join(serviceDir, component.Name, component.Version)
						componentExists, err := dir.Exists(componentDir)
						if err != nil {
							commandline.Interface.Error(err.Error())
							return 1
						}

						if componentExists {
							// commandline.Interface.Warn(fmt.Sprintf("Running profile %s on %s@%s/%s@%s/%s@%s", action, profile.Name, profile.Version, service.Name, service.Version, component.Name, component.Version))
							componentDetails, err := components.GetComponent(component.Name)
							if err != nil {
								commandline.Interface.Error(err.Error())
								return 1
							}

							chart, err := parseHelmChart(path.Join(componentDir, "Chart.yaml"))
							if err != nil {
								commandline.Interface.Error(err.Error())
								return 1
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
							if p.Deploy {
								actionCommand = fmt.Sprintf("cd %s && helm upgrade --install %s d11-helm-charts/%s -f %s -f %s -n %s",
									componentDir,
									componentDetails.Name,
									chart.Name,
									path.Join(componentDir, "values.yaml"),
									path.Join(componentDir, "values-stag.yaml"),
									*envName,
								)

								status = shell.Exec(actionCommand)
								if status > 0 {
									return status
								}
							} else if p.Destroy {
								actionCommand = fmt.Sprintf("cd %s && helm uninstall %s -n %s",
									componentDir,
									componentDetails.Name,
									*envName,
								)

								status = shell.Exec(actionCommand)
								if status > 0 {
									return status
								}
							}
						} else {
							commandline.Interface.Error(fmt.Sprintf("Error while reading component, does not exists. %s", componentDir))
							return 1
						}
					}

				} else {
					commandline.Interface.Error(fmt.Sprintf("Error while reading service, does not exists. %s", serviceDir))
					return 1
				}
			}
		} else {
			commandline.Interface.Error(fmt.Sprintf("Error while reading profile, does not exists. %s", profileDir))
			return 1
		}

	}

	commandline.Interface.Error("Not a valid command")
	return 1
}

func (p *Profile) Help() string {
	if p.Create {
		return commandHelper("create", "profile", []string{
			"--profile=name of profile to create",
			"--version=version of profile to create",
			"--file=yaml file to read profile properties from",
		})
	}
	if p.Delete {
		return commandHelper("delete", "profile", []string{
			"--profile=name of profile to delete",
			"--version=version of profile to delete",
		})
	}
	if p.List {
		return commandHelper("list", "profile", []string{})
	}
	if p.Describe {
		return commandHelper("describe", "profile", []string{
			"--profile=name of profile to describe",
			"--version=version of profile to describe",
		})
	}
	if p.Deploy {
		return commandHelper("deploy", "profile", []string{
			"--profile=name of profile to deploy",
			"--version=version of profile to deploy",
			"--env=name of env to deploy the profile in",
		})
	}
	if p.Destroy {
		return commandHelper("destroy", "profile", []string{
			"--profile=name of profile to destroy",
			"--version=version of profile to destroy",
			"--env=name of env to destroy the profile in",
		})
	}

	return defaultHelper()
}

func (p *Profile) Synopsis() string {
	if p.Create {
		return "create a profile"
	}
	if p.Delete {
		return "delete a profile"
	}
	if p.List {
		return "list all active profiles"
	}
	if p.Describe {
		return "describe a profile"
	}
	if p.Deploy {
		return "deploy a profile"
	}
	if p.Destroy {
		return "destroy a profile"
	}

	return defaultHelper()
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
