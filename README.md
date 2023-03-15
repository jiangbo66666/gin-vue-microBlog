# gin-vue-microBlog
基于gin创建的microBlog项目，项目需要数据库的支持，自行创建数据库信息即可运行
## 目录结构

controller层 - routers文件夹，主要执行与前端的api的映射，调用service层的方法

service层 - service文件夹，主要对获取到的数据执行逻辑上的处理，并且返回到controller层

domain+dao层 - models文件夹，定义实体类，查询数据库。

