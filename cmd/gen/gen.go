package gen

import (
	"github.com/pow1e/psobf-gen/cmd"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(genCmd)
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "生成powershell脚本",
	Long:  `生成powershell脚本`,
}
