# 目录与文件创建工具

这个命令行工具用于快速创建目录和对应的文件，特别适合需要同时创建多个相关目录和文件的场景。

## 功能特点

- 🚀 一键创建目录和文件
- 🔢 支持批量创建多个目录和文件
- 📁 自动处理多级目录结构
- ⚠️ 内置安全检查和错误处理
- 📝 可自定义文件扩展名

## 安装方法

1. 克隆仓库或下载源码：
   ```bash
   git clone https://github.com/SH-ke/dir-file-creator.git
   cd dir-file-creator
   ```

2. 编译程序：
   ```bash
   go build -o dfc
   ```

3. (可选) 添加到系统路径：
   ```bash
   sudo mv dfc /usr/local/bin/
   ```

## 使用方法

### 基本语法
```bash
dfc [选项] <目录名> [文件名]
```

### 选项
- `-type`: 指定文件扩展名 (默认: "go")
- `-h`/`--help`: 显示帮助信息

### 使用示例

1. **创建单个目录和文件**
   ```bash
   dfc mymodule
   ```
   - 创建目录: `mymodule`
   - 创建文件: `mymodule/mymodule.go`

2. **批量创建多个目录和文件**
   ```bash
   dfc controller,service,repository
   ```
   - 创建目录: `controller`, `service`, `repository`
   - 创建文件:
     - `controller/controller.go`
     - `service/service.go`
     - `repository/repository.go`

3. **指定文件扩展名**
   ```bash
   dfc -type js utils,helpers
   ```
   - 创建目录: `utils`, `helpers`
   - 创建文件:
     - `utils/utils.js`
     - `helpers/helpers.js`

4. **多个目录共享文件名**
   ```bash
   dfc pkg/internal,pkg/external main
   ```
   - 创建目录: `pkg/internal`, `pkg/external`
   - 创建文件:
     - `pkg/internal/main.go`
     - `pkg/external/main.go`

5. **为每个目录指定不同文件名**
   ```bash
   dfc ui,api index,router
   ```
   - 创建目录: `ui`, `api`
   - 创建文件:
     - `ui/index.go`
     - `api/router.go`

6. **创建多级目录结构**
   ```bash
   dfc cmd/server,cmd/client
   ```
   - 创建目录: `cmd/server`, `cmd/client`
   - 创建文件:
     - `cmd/server/server.go`
     - `cmd/client/client.go`

## 使用场景

1. **项目初始化**
   ```bash
   dfc cmd,internal/pkg,api/handlers,config
   ```

2. **模块化开发**
   ```bash
   dfc user,product,order -type ts
   ```

3. **多环境配置**
   ```bash
   dfc env/dev,env/staging,env/prod config
   ```

## 注意事项

1. 目录名和文件名支持逗号分隔的多个值
2. 当只提供一个参数时，会自动创建多个同名目录和文件
3. 当提供两个参数时：
   - 如果第二个参数是单个值，所有目录使用相同文件名
   - 如果第二个参数是多个值，必须与目录数量匹配
4. 如果目录已存在，会显示错误并退出
5. 创建过程中发生错误会自动清理已创建的目录

## 帮助信息

```bash
$ dfc -h
用法: dfc [选项] <目录名> [文件名]
目录名和文件名可用逗号分隔多个值
若只提供一个参数，则创建多个同名目录和文件
若提供两个参数且文件名包含逗号，必须与目录数量匹配
选项:
  -type string
        文件类型后缀 (默认 "go")
```

## 贡献指南

欢迎提交Issue和Pull Request！请确保：
1. 代码符合Go编码规范
2. 添加必要的测试用例
3. 更新README文档

## 许可证

本项目采用 [MIT 许可证](LICENSE)