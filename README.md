# Memorandum 备忘录
此项目采用Gin + Gorm + Mysql,基于RESTful API实现的备忘录

## 项目运行
### 克隆项目
    git clone https://github.com/crazyfrankie/memorandum
    cd memorandum
### 下载依赖
    go mod tidy
### 运行项目
    go run main.go

## 接口文档
项目运行之后访问https://localhost:port/swagger/index.html

## 项目主要功能
- 用户注册/登录
- 创建/删除/更新/查找 备忘录
- 分页查找

## 项目主要依赖
- Golang 1.22.4
- Gin
- Gorm
- Mysql
- jwt-go
- logrus
- ini
- go-swagger

## 项目结构

    Memorandum/
    ├── config
    ├── consts
    ├── controller
    ├── docs
    ├── middleware
    ├── pkg
    │  ├── ctl
    │  └── util
    ├── repository
    │  ├── db
    │     ├── dao
    │     └── model
    ├── routes
    └── service

- config:存储配置文件
- consts:存储常量
- controller:定义接口函数
- docs:接口文档
- middleware:应用中间件
- pkg:工具包
- pkg/ctl:响应体定义
- pkg/util:工具函数
- repository:数据操作
- repository/db:持久数据库操作
- repository/db/dao:数据库crud
- repository/dao/model:数据模型定义
- routes:路由逻辑
- service:接口函数实现

  ## 配置文件
  在config下创建config.ini,内容如下
  
      [service]
      AppMode = debug
      port = :your_port
      
      [mysql]
      user = your_user
      password = your_password
      host = localhost
      port = 3306
      db = your_db

