# TRON 靓号地址生成器辅助脚本

这些脚本为非技术用户提供了一个简单的界面，可以在**不直接使用命令行**的情况下生成 TRON 靓号地址。

## Windows 脚本

### 使用步骤

1. 从 [GitHub Releases](https://github.com/bigbigha/TRON-Vanity-Address-Generator/releases) 页面下载 `tronvanity-windows-amd64.exe`。
2. 从本目录下载 `windows-run.bat`。
3. 将这两个文件放在**同一个文件夹**中。
4. 双击运行 `windows-run.bat`。
5. 根据屏幕提示输入你想要的前缀（prefix）、后缀（suffix）和要生成的地址数量。

### 脚本特点

- 简单的文本交互界面。
- 不需要任何命令行基础。
- 会自动循环运行，直到你按下 `Ctrl + C` 退出。

## macOS / Linux 脚本

### 使用步骤

1. 从 [GitHub Releases](https://github.com/bigbigha/TRON-Vanity-Address-Generator/releases) 页面下载适合你系统的二进制文件：
    - `tronvanity-darwin-amd64`（Intel Mac）
    - `tronvanity-darwin-arm64`（Apple Silicon Mac）
    - `tronvanity-linux-amd64`（Linux）
2. 从本目录下载 `unix-run.sh`。
3. 将二进制文件和 `unix-run.sh` 放在**同一个文件夹**中。
4. 打开终端（Terminal），切换到该文件夹。
5. 运行：`chmod +x unix-run.sh`
6. 运行：`./unix-run.sh`
7. 根据屏幕提示进行操作。

### 脚本特点

- 自动检测你的操作系统和架构。
- 简单的文本交互界面。
- 不需要高级命令行知识。
- 自动为二进制文件添加可执行权限。
- 会自动循环运行，直到你按下 `Ctrl + C` 退出。

## 安全说明

这些脚本只是对本地二进制程序的简单封装，它们不会：

- 访问互联网
- 上传任何数据
- 安装软件
- 修改系统文件

脚本和主程序都在你的本地电脑上**完全离线**运行。
