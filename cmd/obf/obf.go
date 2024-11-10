package obf

import (
	"bufio"
	"fmt"
	"github.com/pow1e/psobf-gent/cmd"
	"github.com/pow1e/psobf-gent/pkg/flags"
	"github.com/pow1e/psobf-gent/pkg/obfuscate"
	"github.com/pow1e/psobf-gent/pkg/utils"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	cmd.RootCmd.AddCommand(obfCmd)
}

var obfCmd = &cobra.Command{
	Use:   "obf",
	Short: "混淆",
	Long:  `输入/读取powershell脚本，对其进行混淆`,
	Run: func(cmd *cobra.Command, args []string) {
		if flags.Level == "0" {
			log.Fatal("请使用-l参数指定模糊等级")
		}
		var (
			psScript string
		)

		// 判断是否从文件中读取
		if flags.Input == "" {
			fmt.Printf("请输入需要混淆的powershell脚本: ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				psScript = scanner.Text()
			}
		} else {
			// 从文件中读取
			bytes, err := ioutil.ReadFile(flags.Input)
			if err != nil {
				log.Fatal("文件读取失败")
			}
			psScript = string(bytes)
		}

		obfuscated := obfuscate.GenerateObfuscatedScript(psScript, flags.Level)

		utils.OutputScript(obfuscated)
	},
}
