/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/netapp/capv-bootstrap/pkg/config/types"
	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		spec := &types.ConfigSpec{}
		deployManagementCluster(spec)
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
	deployCmd.Flags().StringVarP(&(cliSettings.config), "config", "c", "", "The config.yaml file to use for deployment (bypass prompts)")
}

func deployManagementCluster(spec *types.ConfigSpec) {
	if cliSettings.config == "" {
		fmt.Println("need to generate a config here")
		runEasyConfig()
	}

	if cliSettings.config == "" {
		fmt.Println("no config file provided and genconfig seems to have failed")
		os.Exit(1)
	}

	collectNetworkInformation(spec)
	collectAdditionalConfiguration(spec)
	configureProviderSpecific(spec)
	if runningFromConfig {
		if err := registration.ValidateSpec(spec); err != nil {
			log.Fatalln(err)
		}
	}
}
