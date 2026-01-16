# AIFileSum

一个基于 AI 的文件内容总结工具，使用 Go + Vue3 构建，通过调用大模型实现智能文本摘要功能。

## 项目概述

AIFileSum 允许用户上传文本文件，系统会自动读取文件内容并通过 AI 生成简洁的摘要。

### 技术栈

**后端:**
- Go 1.25
- Gin Web Framework
- Viper (配置管理)

**前端:**
- Vue 3.5
- Vite 7.3
- Axios

**AI 服务:**
- 通义千问 (Qwen) API

## 迭代过程

### 第一阶段：项目初始化
**AI提示词**
实现一个聊天对话页面，要求支持文件上传，并对文件内容做出总结。
比如以文件形式上传一篇文章， 输出经过大模型总结后的文章核心内容 
后端就使用go+gin,前端使用vue 如果要实现这个功能,我该学习哪些知识? 

**初始化 Go 后端**
- 创建 Go Module (go 1.25)
- 引入 Gin 框架作为 Web 服务器

**初始化 Vue 前端**
- 使用 Vue项目脚手架 创建 Vue3 项目
- 配置开发环境
- 引入 Axios 处理 HTTP 请求

**总结**

 使用 假mock 方法实现file上传与读取

### 第二阶段：核心功能开发

**后端架构设计**

1. 配置管理 (config/config.go)
2. 数据模型 (models/)
3. 业务逻辑 (service/summary.go)

**前端功能实现**

1. API 封装 (src/api/summary.js)
2. UI 组件 (src/components/SummaryPage.vue)
3. 样式设计

**总结**
 成功接入qwen LLM并设置特定的提示词 成功返回 总结性信息
