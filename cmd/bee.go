// Copyright 2013 radical authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// Package cmd ...
package cmd

import (
	"github.com/W3-Partha/Radical/cmd/commands"
	_ "github.com/W3-Partha/Radical/cmd/commands/api"
	_ "github.com/W3-Partha/Radical/cmd/commands/bale"
	_ "github.com/W3-Partha/Radical/cmd/commands/dev"
	_ "github.com/W3-Partha/Radical/cmd/commands/dlv"
	_ "github.com/W3-Partha/Radical/cmd/commands/dockerize"
	_ "github.com/W3-Partha/Radical/cmd/commands/generate"
	_ "github.com/W3-Partha/Radical/cmd/commands/hprose"
	_ "github.com/W3-Partha/Radical/cmd/commands/migrate"
	_ "github.com/W3-Partha/Radical/cmd/commands/new"
	_ "github.com/W3-Partha/Radical/cmd/commands/pack"
	_ "github.com/W3-Partha/Radical/cmd/commands/radiantpro"
	_ "github.com/W3-Partha/Radical/cmd/commands/radicalfix"
	_ "github.com/W3-Partha/Radical/cmd/commands/rs"
	_ "github.com/W3-Partha/Radical/cmd/commands/run"
	_ "github.com/W3-Partha/Radical/cmd/commands/server"
	_ "github.com/W3-Partha/Radical/cmd/commands/update"
	_ "github.com/W3-Partha/Radical/cmd/commands/version"
	"github.com/W3-Partha/Radical/utils"
)

func IfGenerateDocs(name string, args []string) bool {
	if name != "generate" {
		return false
	}
	for _, a := range args {
		if a == "docs" {
			return true
		}
	}
	return false
}

var usageTemplate = `Radical is a Fast and Flexible tool for managing your Radiant Web Application.

You are using radical for radiant v2.x. If you are working on radiant v1.x, please downgrade version to radical v1.12.0

{{"USAGE" | headline}}
    {{"radical command [arguments]" | bold}}

{{"AVAILABLE COMMANDS" | headline}}
{{range .}}{{if .Runnable}}
    {{.Name | printf "%-11s" | bold}} {{.Short}}{{end}}{{end}}

Use {{"radical help [command]" | bold}} for more information about a command.

{{"ADDITIONAL HELP TOPICS" | headline}}
{{range .}}{{if not .Runnable}}
    {{.Name | printf "%-11s"}} {{.Short}}{{end}}{{end}}

Use {{"radical help [topic]" | bold}} for more information about that topic.
`

var helpTemplate = `{{"USAGE" | headline}}
  {{.UsageLine | printf "radical %s" | bold}}
{{if .Options}}{{endline}}{{"OPTIONS" | headline}}{{range $k,$v := .Options}}
  {{$k | printf "-%s" | bold}}
      {{$v}}
  {{end}}{{end}}
{{"DESCRIPTION" | headline}}
  {{tmpltostr .Long . | trim}}
`

var ErrorTemplate = `radical: %s.
Use {{"radical help" | bold}} for more information.
`

func Usage() {
	utils.Tmpl(usageTemplate, commands.AvailableCommands)
}

func Help(args []string) {
	if len(args) == 0 {
		Usage()
	}
	if len(args) != 1 {
		utils.PrintErrorAndExit("Too many arguments", ErrorTemplate)
	}

	arg := args[0]

	for _, cmd := range commands.AvailableCommands {
		if cmd.Name() == arg {
			utils.Tmpl(helpTemplate, cmd)
			return
		}
	}
	utils.PrintErrorAndExit("Unknown help topic", ErrorTemplate)
}
