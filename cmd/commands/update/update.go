package update

import (
	"flag"
	"os"
	"os/exec"

	"github.com/W3-Partha/Radical/cmd/commands"
	"github.com/W3-Partha/Radical/config"
	radicalLogger "github.com/W3-Partha/Radical/logger"
)

var CmdUpdate = &commands.Command{
	UsageLine: "update",
	Short:     "Update Radical",
	Long: `
Automatic run command "go get -u github.com/W3-Partha/Radical" for selfupdate
`,
	Run: updateRadical,
}

func init() {
	fs := flag.NewFlagSet("update", flag.ContinueOnError)
	CmdUpdate.Flag = *fs
	commands.AvailableCommands = append(commands.AvailableCommands, CmdUpdate)
}

func updateRadical(cmd *commands.Command, args []string) int {
	radicalLogger.Log.Info("Updating")
	radicalPath := config.GitRemotePath
	cmdUp := exec.Command("go", "get", "-u", radicalPath)
	cmdUp.Stdout = os.Stdout
	cmdUp.Stderr = os.Stderr
	if err := cmdUp.Run(); err != nil {
		radicalLogger.Log.Warnf("Run cmd err:%s", err)
	}
	return 0
}
