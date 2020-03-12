/*
Copyright © 2020 Utkarsh Srivastava <srivastavautkarsh8097@gmail.com>

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
	"log"
	"os"
	"path"

	"github.com/utkarsh-pro/tempgen/helper"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [language] [template]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Help()
			os.Exit(1)
		}

		language, template := args[0], args[1]
		supportedLanguages = viper.GetStringSlice("supportedLanguages")

		// Check if language is already supported
		languageExists := isPresent(supportedLanguages, language)

		// Check if given template is a directory
		templateIsDir, _ := helper.IsDirectory(template)

		currentPath := helper.GetCurrentPath()

		if languageExists == true {
			existsDir, err := helper.CheckDirectory(path.Join(currentPath, "templates", language, "dir"))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			existsFile, err := helper.CheckDirectory(path.Join(currentPath, "templates", language, "file"))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if templateIsDir == true && existsDir == true {
				fmt.Println(language, "directory template already exists, do you want to continue? [Y/N]")
				var choice string
				fmt.Scanln(&choice)
				if choice == "N" || choice == "n" {
					fmt.Println("Aborting...")
					os.Exit(0)
				}
			} else if templateIsDir == false && existsFile == true {
				fmt.Println(language, "file template already exists, do you want to continue? [Y/N]")
				var choice string
				fmt.Scanln(&choice)
				if choice == "N" || choice == "n" {
					fmt.Println("Aborting...")
					os.Exit(0)
				}
			}
		}

		if templateIsDir == true {
			if err := helper.CopyDir(template, path.Join(currentPath, "templates", language, "dir")); err != nil {
				log.Fatal("Error copying directory", err)
			}
		} else {
			if err := os.MkdirAll(path.Join(currentPath, "templates", language, "file"), os.ModePerm); err != nil {
				log.Fatal("Error copying file", err)
			}
			if err := helper.CopyFile(template, path.Join(currentPath, "templates", language, "file", "main"+"."+language)); err != nil {
				log.Fatal("Error copying file", err)
			}
		}

		// Write changes into config file
		if languageExists == false {
			if err := helper.WriteLanguageToConfig(language); err != nil {
				fmt.Println("Reverting...")
				os.RemoveAll(path.Join(helper.GetCurrentPath(), "templates", language))
				fmt.Println("Reverted")
				log.Fatal("Error writing config: ", err)
			}
		}

		fmt.Println("Successfully added", language)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
