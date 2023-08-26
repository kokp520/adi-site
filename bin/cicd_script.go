package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
    // 設定 Hugo 專案目錄的路徑
    hugoProjectPath := "/Users/adi.wu/blog/kokp520.github.io"

    // 執行 Hugo 命令
    if err := executeCommand(hugoProjectPath, "hugo"); err != nil {
        fmt.Println("Hugo 執行失敗:", err)
        return
    }

    // 執行 Git 命令
    if err := executeCommands(hugoProjectPath, "git", "add", "."); err != nil {
        fmt.Println("Git add 失敗:", err)
        return
    }

    if err := executeCommands(hugoProjectPath, "git", "commit", "-m", "update"); err != nil {
        fmt.Println("Git commit 失敗:", err)
        return
    }

    if err := executeCommands(hugoProjectPath, "git", "push"); err != nil {
        fmt.Println("Git push 失敗:", err)
        return
    }

    fmt.Println("完成！")
}

func executeCommand(dir, command string, args ...string) error {
    cmd := exec.Command(command, args...)
    cmd.Dir = dir
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

func executeCommands(dir, command string, args ...string) error {
    cmd := exec.Command(command, args...)
    cmd.Dir = dir
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

func getAbsolutePath(relativePath string) (string, error) {
    absPath, err := filepath.Abs(relativePath)
    if err != nil {
        return "", err
    }
    return absPath, nil
}
