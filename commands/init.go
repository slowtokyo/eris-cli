package commands

import (
	"github.com/eris-ltd/eris-cli/Godeps/_workspace/src/github.com/spf13/cobra"
	ini "github.com/eris-ltd/eris-cli/initialize"
)

// flags to add: --no-clone
var Init = &cobra.Command{
	Use:   "init",
	Short: "Initialize the ~/.eris directory with some default services and actions",
	Long: `Create the ~/.eris directory with actions and services subfolders
and clone eris-ltd/eris-actions eris-ltd/eris-services into them, respectively.
`,
	Run: func(cmd *cobra.Command, args []string) {
		ini.Initialize(do.Pull, do.Verbose)
	},
}
