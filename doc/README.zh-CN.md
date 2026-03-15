# TRON 靓号地址生成器（TRON Vanity Address Generator）

一个轻量级的 **开源** TRON 靓号地址生成工具。**安全**，**全离线** 生成带自定义前缀 / 后缀的 TRON 地址，简单、快速、对开发者友好。

## 功能特点

- **完全离线**：无网络请求、无遥测、无任何自动上传。
- **开源可审计**：小而专注的代码库，方便安全审计。
- **纯 CPU**：单一静态二进制，支持 Linux / macOS / Windows。
- **多线程**：自动利用全部 CPU 核心，加速暴力搜索。
- **安全随机数**：使用密码学安全的随机数生成器。

## 下载

无需安装 Go 或 Git，直接下载已编译好的二进制文件即可使用：

- 最新版本：  
  https://github.com/bigbigha/TRON-Vanity-Address-Generator/releases/latest

在 Release 页面中，根据你的系统下载对应文件：

- **Windows**：`tronvanity-windows-amd64.exe`
- **macOS（Intel）**：`tronvanity-darwin-amd64`
- **macOS（Apple Silicon）**：`tronvanity-darwin-arm64`
- **Linux**：`tronvanity-linux-amd64`

建议同时下载 `SHA256SUMS.txt`，校验下载文件是否被篡改。

## 快速开始（非开发者）

1. 从 Releases 页面下载适合你系统的二进制文件，并放到任意目录。
2. 打开终端 / 命令行，切换到该目录。
3. 运行命令生成靓号地址。

## 辅助脚本（非开发者）

如果你不熟悉命令行，可以使用仓库中的辅助脚本：

- [Windows / macOS / Linux 辅助脚本说明](../scripts/README.zh-CN.md)


### Windows 示例

```cmd
:: 生成以 T888 开头的地址
tronvanity-windows-amd64.exe --prefix T888

:: 生成以 9999 结尾的地址
tronvanity-windows-amd64.exe --suffix 9999
```

### macOS 示例

```bash
# 仅第一次需要：赋予执行权限
chmod +x tronvanity-darwin-amd64      # Intel Mac
chmod +x tronvanity-darwin-arm64      # Apple Silicon Mac

# 生成以 T888 开头的地址
./tronvanity-darwin-amd64 --prefix T888

# 生成以 9999 结尾的地址
./tronvanity-darwin-arm64 --suffix 9999
```

### Linux 示例

```bash
# 仅第一次需要：赋予执行权限
chmod +x tronvanity-linux-amd64

# 生成以 T888 开头的地址
./tronvanity-linux-amd64 --prefix T888

# 生成以 9999 结尾的地址
./tronvanity-linux-amd64 --suffix 9999
```

> 提示：本工具可在完全离线（断网、隔离）的电脑上运行，用于生成大额资金地址时更安全。

## 命令行用法

```bash
# 仅匹配前缀
tronvanity --prefix T888

# 仅匹配后缀
tronvanity --suffix 9999

# 同时匹配前缀和后缀，使用 8 个 worker，找到 3 个后退出
tronvanity --prefix T666 --suffix 888 --workers 8 --count 3

# 生成 5 个以 T123 开头的地址
tronvanity --prefix T123 --count 5
```

常用参数：

- `--prefix`：要匹配的前缀（如 `T888`）
- `--suffix`：要匹配的后缀（如 `9999`）
- `--workers`：并发 worker 数量（默认＝CPU 核心数）
- `--count`：找到多少个匹配后退出（默认 1）

## 安全建议

1. **尽量离线生成**：高价值地址建议在离线 / 物理隔离的环境中生成。
2. **仔细核对地址**：给地址充值前，确认前缀 / 后缀确实符合预期。
3. **妥善保管私钥**：使用硬件钱包、密码管理器或离线加密备份存储，避免在网页或不可信应用中粘贴私钥。
4. **从小额开始测试**：先用小额测试收款和转出流程，再存入大额资产。

## 工作原理（简要）

1. 使用 Go 的密码学安全随机数生成 32 字节私钥。
2. 使用 secp256k1 椭圆曲线（TRON / Ethereum 使用的同一曲线）推导对应公钥。
3. 对去掉前缀字节 `0x04` 的非压缩公钥做 Keccak-256 哈希，取最后 20 字节。
4. 在前面加上 TRON 主网前缀 `0x41`，得到 21 字节地址数据。
5. 使用 Base58Check 编码为标准的 `T...` TRON 地址。
6. 多个 worker 并行重复以上过程，检查是否匹配指定前缀 / 后缀，匹配则输出地址与私钥。

全流程都在本地执行，无任何网络请求，也不会发送或上传私钥。

## 免责声明

本软件按 "现状" 提供，不附带任何形式的保证。使用风险由你自行承担。  
在使用真实资金前，请务必核对生成的地址和私钥，并确保你有能力安全管理相关资产。
