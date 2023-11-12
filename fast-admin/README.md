## 项目介绍

[fast-admin]是基于 Vue3 + Vite4+ TypeScript5 + Element-Plus + Pinia 等最新主流技术栈构建的后台管理前端模板（配套后端源码）。

项目有以下特性：

- 基于 fast-admin 升级到 vue3 版本，无自定义封装，易上手，减少学习成本。
- 提供了配套的 Golang 后端接口
- 权限系统功能齐全，包括用户管理、角色管理、菜单管理、字典管理和部门管理等，以满足您对权限管理的需求。
- 项目还提供了基础设施支持，包括动态路由、按钮级别的权限控制、国际化支持、代码规范、Git 提交规范以及常用组件的封装，以便开发人员更高效地开发和维护项目。

## 环境准备

| 环境                 | 名称版本                                                     | 备注                                                         |
| -------------------- | :----------------------------------------------------------- | ------------------------------------------------------------ |
| **开发工具**         | VSCode                                                       | [下载地址](https://code.visualstudio.com/Download)           |
| **运行环境**         | Node 16+                                                     | [下载地址](http://nodejs.cn/download)                        |
| **VSCode插件(必装)** | 1. `Vue Language Features (Volar)` <br/> 2. `TypeScript Vue Plugin (Volar)`  <br/>3. 禁用 Vetur | |

## 项目启动

```bash
# 克隆代码 && 切换目录
cd fast-admin

# 安装 pnpm
npm install pnpm -g

# 安装依赖
pnpm install

# 启动运行
pnpm run dev
```

## 项目部署

```bash
# 项目打包
pnpm run build:prod

# 上传文件至远程服务器
将打包生成在 `dist` 目录下的文件拷贝至 `/usr/share/nginx/html` 目录

# nginx.cofig 配置
server {
 listen     80;
 server_name  localhost;
 location / {
   root /usr/share/nginx/html;
   index index.html index.htm;
 }
 # 反向代理配置
 location /prod-api/ {
   proxy_pass http://api.imguo.com/; # api.imguo.com替换成你的后端API地址
 }
}
```
