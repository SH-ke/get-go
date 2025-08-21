package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var fileType string
	flag.StringVar(&fileType, "type", "go", "文件类型后缀")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "用法: %s [选项] <目录名> [文件名]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "目录名和文件名可用逗号分隔多个值\n")
		fmt.Fprintf(os.Stderr, "若只提供一个参数，则创建多个同名目录和文件\n")
		fmt.Fprintf(os.Stderr, "若提供两个参数且文件名包含逗号，必须与目录数量匹配\n")
		fmt.Fprintf(os.Stderr, "选项:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 || len(args) > 2 {
		flag.Usage()
		os.Exit(1)
	}

	// 分割目录名
	dirs := strings.Split(args[0], ",")
	if len(dirs) == 0 {
		fmt.Fprintln(os.Stderr, "错误: 至少需要一个目录名")
		os.Exit(1)
	}

	// 处理文件名
	filenames := make([]string, len(dirs))
	if len(args) == 2 {
		// 如果有第二个参数
		names := strings.Split(args[1], ",")
		if len(names) == 1 {
			// 单个文件名 - 所有目录使用相同文件名
			for i := range filenames {
				filenames[i] = names[0]
			}
		} else if len(names) == len(dirs) {
			// 多个文件名 - 与目录一一对应
			copy(filenames, names)
		} else {
			fmt.Fprintf(os.Stderr, "错误: 文件名数量(%d)与目录数量(%d)不匹配\n", len(names), len(dirs))
			os.Exit(1)
		}
	} else {
		// 无文件名参数 - 使用目录名作为文件名
		for i, dir := range dirs {
			// 获取目录最后部分作为文件名
			base := filepath.Base(dir)
			// 去除可能存在的路径分隔符
			if base == "." || base == string(filepath.Separator) {
				base = "default"
			}
			filenames[i] = base
		}
	}

	var createdDirs []string // 记录成功创建的目录

	// 创建所有目录
	for _, dir := range dirs {
		// 检查目录是否存在
		if _, err := os.Stat(dir); err == nil {
			fmt.Fprintf(os.Stderr, "错误: 目录已存在 '%s'\n", dir)
			cleanup(createdDirs)
			os.Exit(1)
		}

		// 创建目录（支持多级）
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "创建目录失败 '%s': %v\n", dir, err)
			cleanup(createdDirs)
			os.Exit(1)
		}
		createdDirs = append(createdDirs, dir)
	}

	// 创建所有文件
	for i, dir := range dirs {
		filePath := filepath.Join(dir, filenames[i]+"."+fileType)
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "创建文件失败 '%s': %v\n", filePath, err)
			cleanup(createdDirs)
			os.Exit(1)
		}
		file.Close()
	}

	// 打印创建结果
	fmt.Println("成功创建:")
	for i, dir := range dirs {
		fmt.Printf("  目录: %-20s → 文件: %s.%s\n",
			dir,
			filepath.Join(dir, filenames[i]),
			fileType)
	}
}

// 清理已创建的目录
func cleanup(dirs []string) {
	for _, dir := range dirs {
		os.RemoveAll(dir)
	}
}
