package utils

import (
	"fmt"
	"github.com/pow1e/psobf-gen/pkg/flags"
	"io/ioutil"
	"log"
	"path/filepath"
)

func OutputScript(obfuscated string) {
	// 判断是否文件输出
	if flags.Output != "" {
		if ext := filepath.Ext(flags.Output); ext != ".ps1" {
			log.Fatal("拓展名无效，请重试")
		}

		err := ioutil.WriteFile(flags.Output, []byte(obfuscated), 0644)
		if err != nil {
			log.Fatal("文件写入失败")
		}
	} else {
		fmt.Println("混淆powershell脚本:", obfuscated)
	}
}
