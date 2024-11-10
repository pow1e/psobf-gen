package cmd

import (
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/pow1e/psobf-gent/pkg/flags"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "psobf",
	Short: "Powershell Obfuscator",
	Long: `
	██████╗ ███████╗ ██████╗ ██████╗ ███████╗
	██╔══██╗██╔════╝██╔═══██╗██╔══██╗██╔════╝
	██████╔╝███████╗██║   ██║██████╔╝█████╗  
	██╔═══╝ ╚════██║██║   ██║██╔══██╗██╔══╝  
	██║     ███████║╚██████╔╝██████╔╝██║     
	╚═╝     ╚══════╝ ╚═════╝ ╚═════╝ ╚═╝     
	@pow1e 
	v.1.0
`,
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&flags.Input, "input", "i", "", "设置输入位置")
	RootCmd.PersistentFlags().StringVarP(&flags.Output, "output", "o", "", "设置输出位置,不填则默认终端打印")
	RootCmd.PersistentFlags().StringVarP(&flags.Level, "level", "l", "", "设置混淆强度 1|2|3|4|5|all")
	RootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Execute() {
	cc.Init(&cc.Config{
		RootCmd:  RootCmd,
		Headings: cc.HiGreen + cc.Underline,
		Commands: cc.Cyan + cc.Bold,
		Example:  cc.Italic,
		ExecName: cc.Bold,
		Flags:    cc.Cyan + cc.Bold,
	})
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
