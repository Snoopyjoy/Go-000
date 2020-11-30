
一个简单的http服务demo 用于验证error传递  
### 问题
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？  

### 方案
三层架构：
1. dao 从数据库获取数据
2. service 业务逻辑层
3. api 接口层  

DAO 报sql.ErrNoRows时，return 一个预定义在dao层里的sentinel err fmt.Errorf("%w") 携带具体信息并wrap后抛出。  
DAO 报其他错误时，Wrap错误后抛给上层  
service 层调用dao报错时直接返回错误给上层
service 层如果需要对未找到记录的err做特殊处理则通过errors.Is判断后处理，并吞掉这个错误
api 层里打印并处理错误

### 启动
```shell
    go run main.go
```

### 测试case
正常： http://127.0.0.1:8080/user/1001  
sql.ErrNoRows: http://127.0.0.1:8080/user/1002  
service 内部错误: http://127.0.0.1:8080/user/1003     