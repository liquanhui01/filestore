filestore是一个基于Go语言和Gin框架开发的文件传输和存储系统，系统正在积极的更新和维护中。
### 系统具有以下功能：
1. 用户模块：用户注册、登陆校验和授权；
2. 文件上传模块：分为普通传输和秒传两种方式，存储在阿里云OSS；
3. 文件下载模块；
4. 文件夹模块：每个用户都可以设置多个文件夹，在指定的文件夹下存储文件；
5. 垃圾箱模块：文件列表和清空功能。
### 技术栈：
1. 服务端：
    - Go语言
    - Gin框架
    - Gorm框架
    - Redis
    - RabbitMq实现异步上传文件
    - 数据库：MySQL5.7
2. 客户端：
    - Vue
    - elementui-plus框架


![截屏2021-10-30 18.42.17](https://segmentfault.com/img/bVcUATi/view)
