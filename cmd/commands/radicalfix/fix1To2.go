// Copyright 2020
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package radicalfix

import (
	"os"
	"os/exec"

	radicalLogger "github.com/W3-Partha/Radical/logger"
)

func fix1To2() int {
	radicalLogger.Log.Info("Upgrading the application...")

	cmdStr := `go get -u github.com/W3-Engineers-Ltd/Radiant@develop`
	err := runShell(cmdStr)
	if err != nil {
		radicalLogger.Log.Error(err.Error())
		radicalLogger.Log.Error(`fetch v2.0.1 failed. Please try to run: export GO111MODULE=on
and if your network is not stable, please try to use proxy, for example: export GOPROXY=https://goproxy.cn;'
`)
		return 1
	}

	cmdStr = `find ./ -name '*.go' -type f -exec sed -i '' -e 's/github.com\/astaxie\/radiant/github.com\/radiant\/radiant\/v2\/adapter/g' {} \;`
	err = runShell(cmdStr)
	if err != nil {
		radicalLogger.Log.Error(err.Error())
		return 1
	}
	cmdStr = `find ./ -name '*.go' -type f -exec sed -i '' -e 's/"github.com\/radiant\/radiant\/v2\/adapter"/radiant "github.com\/radiant\/radiant\/v2\/adapter"/g' {} \;`
	err = runShell(cmdStr)
	if err != nil {
		radicalLogger.Log.Error(err.Error())
		return 1
	}
	return 0
}

func runShell(cmdStr string) error {
	c := exec.Command("sh", "-c", cmdStr)
	c.Stdout = os.Stdout
	err := c.Run()
	if err != nil {
		radicalLogger.Log.Errorf("execute command [%s] failed: %s", cmdStr, err.Error())
		return err
	}
	return nil
}
