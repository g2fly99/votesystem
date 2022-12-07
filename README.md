# votesystem
a simple vote system

#### 关于权限
1. 管理员和投票员使用不同的接口
1. 登陆时，需要携带对应的session
1. 接口中暂未涉及

#### 候选人上限
1. 未限定候选人上限，也未做分页处理;默认仅显示前1000条数据
2. 不同选举活动;候选人未做去重
3. 同一个选举活动，候选人以姓名去重

#### 投票数据
1. 当前版本，投票数量，未考虑分表的设计；
2. 有时间的话，后续可以设计按月分表或者按年分表；考虑到查询的效率，MySQL单表建议数据在500w行数据以内

#### 选举参数
1. 选举有效期设置到秒，未做过多的设计
1. 创建时设置过期时间
1. 后续可以增加开始的时候设置时间

#### 关于邮件
1. 选举结束时，需要发送邮件;
1. 选举到期时，自动发送邮件
1. 管理员点结束时，触发邮件发送
1. 发送邮件时，遍历所有的投票人，把相关结果发送给所有投票人
1. 为了避免垃圾邮件，以及测试错误邮箱的配置，仅做了打印日志的处理。

# interface
see interface.md


# how to run ？
there are 3 way to start the program:

## dockerfile 1
1. download Dockerfile
2. get ready mysql and redis
3. redis: 127.0.0.1:6379; password:Aa123456!@#
4. mysql: 127.0.0.1:3306; database:ipcc; root password:  Aa123456!@#
5. after build then run the container

## dockerfile 2
1. download Dockerfile and directory:conf
2. get ready mysql and redis
3. change Dockerfile, add ENV: ./conf/:/go/src/votesystem/conf
4. change config
5. after build then run the container

## docker-compose
1. download docker-compose.tar
2. then run docker-compose up




