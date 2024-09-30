中文
基于go(1.23)语言框架gin写的堡垒机api
分为    1.实现层
        2.服务层
        3.中间件层
        4.api层
编写思路
        1.数据接收 ——> 路由函数 ——>josn反序列化——>api层
        2.数据校验——>validator
        3.数据增删改查——>grom
        4.数据发送——>josn序列化——>前端接收
鉴权    JWT


EN
Fortress Machine Built on Go Language Go 1.23

The bastion machine API is written based on the Go language framework Gin, structured into:

Implementation Layer
Service Layer
Middleware Layer
API Layer
Development Approach:

Data Reception → Routing Function → JSON Deserialization → API Layer
Data Validation → Validator
Data CRUD Operations → GORM
Data Sending → JSON Serialization → Frontend Reception
Authentication: JWT
