# v1-service-readme.MD

## Overview

本项目基于Go进行开发，主要用于对InkFinance Governance Module进行数据支持。

其主要实现：

1. Web3.0账号与Social账户绑定
2. InkFinance不同的合约（包括不限于派生出来的DAO|Committee|Proposal)事件监听及数据关联
3. 链上数据中心化统一存储查询功能
4. 特定业务场景下的接口工具（比如限制某些操作等判断）

## 重点目录结构说明

1. chain 扫链业务代码，通过与
2. conf 数据库配置及初始化
3. constant 系统中常量的定义
4. httprouter 接口路由定义
5. request 请求数据定义
6. scripts 脚本（数据库）
7. service 接口定义及实现
8. cmd 基于GORM的数据库对象生成脚本
9. dal 自动生成的数据库操作对象及数据模型
10. transport 服务接口与Handler绑定的定义

## 如何部署

1. 安装go （建议安装：1.18)
2. 克隆本项目git仓库
3. 安装MySQL数据库
4. 执行初始化脚本
    1. 进入到项目根目录，找到./scripts/db_inkfinance_init.sql 在第3步安装到MySQL执行脚步初始化数据库及相关表结构
5. 执行监听链任务的脚本
    1. 进入到项目根目录，找到./scripts/init_scan_task.sql 执行脚步（确保在第4步执行完成后执行）
6. 配置环境文件
    1. 基于项目根目录第config-[env].yml， 把[env]替换为相关的执行环境，并且把文件内的MySQL 以及区块链RPC访问配置为需要的链接
7. 指定配置文件启动项目
    1. 全部配置完成后，在项目根目录执行：go run main [env]，即开始扫链，同时根据接口返回对应的链上事件抛出的加工后的数据来满足查询需求。