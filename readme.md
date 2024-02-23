# gin 应用服务

## 介绍

- 权限分配服务
- 根据角色分配菜单，接口，页面元素
- 业务服务层将角色与该服务同步

- [表结构](./db/permissions.sql)
- [接口逻辑](./docs/接口逻辑.md)

## 效果预览地址

- http://hpyyb.cn/wide/permissions/
- 前端代码：https://github.com/yubo9807/wide/tree/main/src/sub-permissions/views/permissions

## 启动

- 启动： `go run src/server.go`
- 打包： `./scripts/build.sh`
