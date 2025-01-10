# 权限分配服务对接

## 介绍

- 权限分配服务
- 根据角色分配菜单，接口，页面元素
- 业务服务层将角色与该服务同步

- [表结构](./db/permissions.sql)
- [接口逻辑](./docs/接口逻辑.md)

## 应用接入

1. `./server` 后会生成一个 `config.yml` 的配置文件
2. 不管是 Nginx 还是 后端程序代理（推荐），添加请求头 `Open-Id`，`config.yml` 中设置的 openId
3. openId 不可暴露在前端
4. 业务接口权限分配，需要调用该服务的 {baseURL}/v1/api/interface/authority 接口来获取是否具有权限

## 效果预览地址

- 前端代码：https://github.com/yubo9807/admin-template (本地启动查看效果)

## 启动

- 启动： `go run src/server.go`
- 打包(linux)： `./scripts/build.sh`
- 生产： `./server`
