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
package radiantpro

import (
	"strings"

	"github.com/W3-Partha/Radical/cmd/commands"
	"github.com/W3-Partha/Radical/internal/app/module/radiantpro"
	"github.com/W3-Partha/Radical/logger"
)

var CmdRadiantPro = &commands.Command{
	UsageLine: "pro [command]",
	Short:     "Source code generator",
	Long:      ``,
	Run:       RadiantPro,
}

func init() {
	CmdRadiantPro.Flag.Var(&radiantpro.SQL, "sql", "sql file path")
	CmdRadiantPro.Flag.Var(&radiantpro.SQLMode, "sqlmode", "sql mode")
	CmdRadiantPro.Flag.Var(&radiantpro.SQLModePath, "sqlpath", "sql mode path")
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRadiantPro)
}

func RadiantPro(cmd *commands.Command, args []string) int {
	if len(args) < 1 {
		radicalLogger.Log.Fatal("Command is missing")
	}

	if len(args) >= 2 {
		cmd.Flag.Parse(args[1:])
	}

	gcmd := args[0]
	switch gcmd {
	case "gen":
		radiantpro.DefaultRadiantPro.Run()
	case "config":
		radiantpro.DefaultRadiantPro.GenConfig()
	case "migration":
		radiantpro.DefaultRadiantPro.Migration(args)
	default:
		radicalLogger.Log.Fatal("Command is missing")
	}
	radicalLogger.Log.Successf("%s successfully generated!", strings.Title(gcmd))
	return 0
}
