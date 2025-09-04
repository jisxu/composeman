# composeman

一个用 Go 编写的轻量级 Docker Compose 管理工具，可以方便地在多个项目目录下批量执行 `docker compose` 命令。

## 功能特性

- **update**：相当于 `docker compose pull && docker compose up -d`
- **start**：启动容器 (`docker compose start`)
- **stop**：停止容器 (`docker compose stop`)
- **restart**：重启容器 (`docker compose restart`)
- **version**：查看程序版本（支持通过 `-ldflags` 注入版本号）

## 安装

### 本地构建

```bash
cd composeman

# 构建二进制
go build -ldflags "-X main.Version=$(git describe --tags --always)" -o composeman
```

### 使用 build.sh 构建（推荐）

```bash
./build.sh
```

构建完成后会生成可执行文件 `composeman`。

## 使用方法

```bash
composeman [命令] <目录1> <目录2> ...
```

### 参数说明

- `命令`（可选，默认 `update`）
  - `update`：更新并重启服务
  - `start`：启动服务
  - `stop`：停止服务
  - `restart`：重启服务
  - `version`：查看版本
- `<目录>`：包含 `compose.yaml` 的项目目录，可传多个

### 示例

```bash
# 在指定目录更新服务（pull + up -d）
composeman update ./docker/project1 ./docker/project2

# 启动服务
composeman start ./docker/project1

# 停止服务
composeman stop ./docker/project2

# 重启多个目录下的服务
composeman restart ./docker/project1 ./docker/project2

# 查看版本
composeman version
```

## 注意事项

- 本工具默认要求项目目录下存在 `compose.yaml` 或 `docker-compose.yaml` 文件。  
- 执行时需要本地环境已正确安装 `docker` 与 `docker compose`。  

---
