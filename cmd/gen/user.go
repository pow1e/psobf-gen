package gen

import (
	"github.com/pow1e/psobf-gen/pkg/gen/user"
	"github.com/spf13/cobra"
)

func init() {
	genCmd.AddCommand(userCmd)
	userCmd.AddCommand(createUserCmd)
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "用户模块",
	Long:  "用户模块",
}

var createUserCmd = &cobra.Command{
	Use:   "create",
	Short: "创建用户",
	Long:  "创建用户",
	Run: func(cmd *cobra.Command, args []string) {
		user.GenerateCreateUserCommand()
	},
}

var addAdminCmd = &cobra.Command{
	Use:   "addAdmin",
	Short: "添加管理员组",
	Long:  "添加管理员组",
	Run: func(cmd *cobra.Command, args []string) {
		user.GenerateAddUserCommand()
	},
}
