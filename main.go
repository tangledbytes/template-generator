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

/*
USAGE:
tempgen create [language] [flags]

Flags:
	-d, --dir		If the generated file is a directory
	-n, --name		Name of the file/directory
	-t, --template  location/url of the template

tempgen set-default [language] [flags]

Flags:
	-d, --dir		If the generated file is a directory
	-t, --template  location/url of the template
*/
package main

import "github.com/utkarsh-pro/tempgen/cmd"

func main() {
	cmd.Execute()
}
