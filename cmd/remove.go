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
	"os"
	"path"

	"github.com/utkarsh-pro/tempgen/helper"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove [language]",
	Short: "Removes the template of the provided language",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			os.Exit(1)
		}

		supportedLanguages = viper.GetStringSlice("supportedLanguages")
		language := args[0]
		languageExists := isPresent(supportedLanguages, language)

		if languageExists == false {
			fmt.Println("Language doesn't exists!")
			os.Exit(0)
		}

		os.RemoveAll(path.Join(helper.GetCurrentPath(), "templates", language))
		helper.RemoveLanguageFromConfig(language)
		fmt.Println("Successfully removed", language)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
