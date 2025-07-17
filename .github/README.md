# GitHub Actions 工作流说明

本项目包含以下 GitHub Actions 工作流，专门针对 auto-refresh-server、gateway、proxy 三个服务：

## 工作流文件

### 1. Docker 构建和推送 (docker.yml)
- **触发条件**: 推送到 master 分支或创建 tag
- **功能**:
  - 为三个服务构建 Docker 镜像
  - 推送到 Docker Hub
  - 支持多架构构建 (linux/amd64, linux/arm64)
  - 镜像命名：
    - `nilorg/wechat-auto-refresh-server`
    - `nilorg/wechat-gateway`
    - `nilorg/wechat-proxy`

### 2. 发布 (release.yml)
- **触发条件**: 推送 git tag (v*)
- **功能**:
  - 为三个服务构建多平台二进制文件
  - 生成校验和文件
  - 自动生成变更日志
  - 创建 GitHub Release

### 3. 开发构建 (dev-build.yml)
- **触发条件**: 手动触发
- **功能**:
  - 可选择构建特定服务或全部服务
  - 构建开发版本 Docker 镜像 (tag: dev)
  - 参考 `scripts/build-dev-docker.sh` 的构建方式

## 使用说明

### 首次设置

在使用 Docker 相关工作流之前，需要在 GitHub 仓库设置中添加以下 Secrets：

- `DOCKER_USERNAME`: Docker Hub 用户名
- `DOCKER_PASSWORD`: Docker Hub 密码或访问令牌 (推荐使用 Personal Access Token)

### 运行流程

1. **自动触发**: 工作流会在代码推送、PR 创建、标签发布时自动运行
2. **手动构建**: 在 GitHub Actions 页面可以手动触发开发构建
3. **发布流程**: 创建 git tag 并推送即可自动发布新版本

## 镜像拉取示例

```bash
# 拉取最新版本
docker pull nilorg/wechat-auto-refresh-server:latest
docker pull nilorg/wechat-gateway:latest
docker pull nilorg/wechat-proxy:latest

# 拉取开发版本
docker pull nilorg/wechat-auto-refresh-server:dev
docker pull nilorg/wechat-gateway:dev
docker pull nilorg/wechat-proxy:dev
```
