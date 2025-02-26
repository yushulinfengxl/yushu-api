# Config Golang

## set go proxy

### 查看当前代理

    go env | findstr GOPROXY

### Go 1.13 及以上（推荐）

#### 打开你的终端并执行

    go env -w GO111MODULE=on
    go env -w GOPROXY=https://goproxy.198587.xyz,direct

### macOS 或 Linux

#### 打开你的终端并执行

    export GO111MODULE=on
    export GOPROXY=https://goproxy.198587.xyz

#### 或者

    echo "export GO111MODULE=on" >> ~/.profile
    echo "export GOPROXY=https://goproxy.198587.xyz" >> ~/.profile
    source ~/.profile

### Windows

#### 打开你的 PowerShell 并执行

    C:\> $env:GO111MODULE = "on"
    C:\> $env:GOPROXY = "https://goproxy.198587.xyz"

#### 或者

    1. 打开“开始”并搜索“env”
    2. 选择“编辑系统环境变量”
    3. 点击“环境变量…”按钮
    4. 在“<你的用户名> 的用户变量”章节下（上半部分）
    5. 点击“新建…”按钮
    6. 选择“变量名”输入框并输入“GO111MODULE”
    7. 选择“变量值”输入框并输入“on”
    8. 点击“确定”按钮
    9. 点击“新建…”按钮
    10. 选择“变量名”输入框并输入“GOPROXY”
    11. 选择“变量值”输入框并输入“https://goproxy.198587.xyz”
    12. 点击“确定”按钮