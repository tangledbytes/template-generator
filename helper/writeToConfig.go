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

package helper

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var cfgFile = "config"

// Config struct defines the configuration
type Config struct {
	SupportedLanguages []string `yaml:"supportedLanguages"`
	Defaults           struct {
		Language string `yaml:"language"`
		Mode     string `yaml:"mode"`
	} `yaml:"defaults"`
}

func init() {
	viper.AddConfigPath(GetCurrentPath())
	viper.SetConfigName(cfgFile)

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// WriteLanguageToConfig will write config to the file
func WriteLanguageToConfig(language string) error {
	C := &Config{}

	err := viper.Unmarshal(&C)

	if err != nil {
		return err
	}

	C.SupportedLanguages = append(C.SupportedLanguages, language)

	d, err := yaml.Marshal(&C)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path.Join(GetCurrentPath(), cfgFile+".yaml"), d, 0644)
	if err != nil {
		return err
	}

	return nil
}

// RemoveLanguageFromConfig will remove a language from the config
func RemoveLanguageFromConfig(language string) error {
	C := &Config{}
	cfgFile := "config"
	j := 0

	err := viper.Unmarshal(&C)

	if err != nil {
		return err
	}

	for _, lang := range C.SupportedLanguages {
		if lang != language {
			C.SupportedLanguages[j] = lang
			j++
		}
	}

	C.SupportedLanguages = C.SupportedLanguages[:j]

	d, err := yaml.Marshal(&C)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path.Join(GetCurrentPath(), cfgFile+".yaml"), d, 0644)
	if err != nil {
		return err
	}

	return nil
}
