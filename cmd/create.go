/*
Copyright © 2020 Utkarsh Srivastava <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"
)

var supportedLanguages = []string{"cpp", "js", "go", "py"}
var template, language, name *string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create command is used to create either a file or directory by using default or a custom template",
	Long:  `create command is used to create either a file or directory by using default or a custom template`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(*template, *name, *language)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createCmd.Flags().BoolP("directory", "d", false, "Default is false, set to true to specify if a directory is to be created")
	template = createCmd.Flags().StringP("template", "t", "", "")
	language = createCmd.Flags().StringP("language", "l", "", "")
	name = createCmd.Flags().StringP("name", "n", "", "")
}
