简易投票系统接口文档
-------

#### 通用返回 

参数名  | 类型 | 约束  | 描述
-------|------|------|-----------
code   | int  | 必选  | 返回码，非200为失败
message | string | 必选 | 具体消息,成功为success
data   | object |  可选  | 具体接口携带返回

- example
``` json
{
    "code": 200,
    "message": "success",
    "data": ""
}

```
 
##### 通用返回码
```
  Success                 = {200, "success", ""}
  ErrParamWrong           = {300, "Param Type Not Match", ""}
  ErrParamEmailWrong      = {301, "Param Email is  invalid", ""}
  ErrParamIdentityNoWrong = {302, "Param Identity Number is invalid", ""}
  ErrNotFound             = {400, "Not Found", ""}
  ErrIsExist              = {401, "Ident or email used", ""}
  ErrEcIsFinish           = {402, "Not Found", "ec is finished"}
  ErrHaveVoted            = {403, "only can vote 1 times", ""}
  ErrVoteInvalid          = {404, "vote the invalid candidate", ""}
  ErrSystem               = {500, "System Busy", ""}
  ErrCancdidateNotEnough  = {600, "candidate not enough", ""}
```
#### 选举:创建

- 创建一个选举活动
- 请求方式:POST
- URL:/v1/electioncampaign?expire=202212042008&description=justfortest

- 入参

参数名  | 类型 | 约束  | 提交方式 |描述
-------|------|------|---------|-----------
expire | string | 必选 |query | 过期时间:202212042008,精确到秒;为方便后续测试，暂未做优化
description | string | query | 可选 | 选举描述


- 返回值
通用返回码

#### 选举:开始
- 开始一个选举活动
- 请求方式:PUT
- URL:/v1/electioncampaign/{ecId}/finish

- 入参
  无


- 返回值
通用返回码

#### 选举:结束
- 结束一个选举活动
- 请求方式:PUT
- URL:/v1/electioncampaign/{ecId}/start

- 入参
  无


- 返回值
通用返回码

#### 选举:获取
- 获取所有选举活动
- 请求方式:GET
- URL:/v1/electioncampaign

- 入参
  无


- 返回值

- 选举活动信息

参数名  | 类型 | 约束   |描述
-------|------|-------|-----------
ecId   | int  | 必选  | 选举活动Id
finished | bool | 必选 | 是否结束,false 未结束
expire | string | 必选 | 过期时间,为节省时间，精确到秒
startTime | string | 必选 |  开始时间
finishTime | string | 必选 |  结束时间
description | string | 必选 |  选举描述
created | string | 必选 | 创建时间
candidates | object array | 必选 | 选举人信息


- 投票人信息

参数名       | 类型     | 约束  |描述
------------|---------|------|---------
username    | string  | 可选  | 姓名,长度255以内
email       | string  | 必选  | 投票人email地址,用于接收投票结果信息，选举活动内唯一，长度255以内
IdentityNo  | string  | 必选  | 可选 | 香港身份证号A123456(7)，仅做格式校验,未做有效性校验，选举活动内唯一
CandidateId | int     | 必选  | 选举人Id
voteNumber  | int     | 必选  | 获得投票数


- example
```json
{
    "code": 200,
    "message": "successful",
    "data": [
        {
            "ecId": 1,
            "finished": true,
            "expire": "2022-12-05T04:08:18+08:00",
            "startTime": "2022-12-04T15:30:48+08:00",
            "finishTime": "2022-12-05T07:02:41+08:00",
            "description": "test",
            "created": "0001-01-01T00:00:00Z",
            "voteNumber": 15,
            "candidates": [
                {
                    "candidateId": 1,
                    "name": "eric",
                    "sex": 0,
                    "age": 18,
                    "description": "just for test",
                    "voteNumber": 12
                },
                {
                    "candidateId": 2,
                    "name": "eric2",
                    "sex": 0,
                    "age": 19,
                    "description": "just for test",
                    "voteNumber": 3
                },
                {
                    "candidateId": 3,
                    "name": "eric3",
                    "sex": 0,
                    "age": 20,
                    "description": "just for test",
                    "voteNumber": 0
                }
            ]
        },
        {
            "ecId": 2,
            "finished": false,
            "expire": "2022-12-05T05:08:11+08:00",
            "startTime": "2022-12-04T14:41:27+08:00",
            "finishTime": "0001-01-01T00:00:00Z",
            "description": "test",
            "created": "0001-01-01T00:00:00Z",
            "voteNumber": 1,
            "candidates": [
                {
                    "candidateId": 4,
                    "name": "test2",
                    "sex": 1,
                    "age": 18,
                    "description": "just for test2",
                    "voteNumber": 1
                },
                {
                    "candidateId": 5,
                    "name": "test3",
                    "sex": 0,
                    "age": 19,
                    "description": "just for test4",
                    "voteNumber": 0
                }
            ]
        },
        {
            "ecId": 3,
            "finished": false,
            "expire": "2022-12-05T05:08:12+08:00",
            "startTime": "0001-01-01T00:00:00Z",
            "finishTime": "0001-01-01T00:00:00Z",
            "description": "test",
            "created": "0001-01-01T00:00:00Z",
            "voteNumber": 0,
            "candidates": []
        },
        {
            "ecId": 4,
            "finished": false,
            "expire": "2022-12-05T05:08:00+08:00",
            "startTime": "0001-01-01T00:00:00Z",
            "finishTime": "0001-01-01T00:00:00Z",
            "description": "test",
            "created": "0001-01-01T00:00:00Z",
            "voteNumber": 0,
            "candidates": []
        }
    ]
}
```


#### 候选人：添加
- 添加一个候选人
- 请求方式:POST
- URL:/v1/electioncampaign/{ecId}/candidate
- 入参

参数名  | 类型 | 约束  | 提交方式 |描述
-------|------|------|---------|-----------
candidate | object array | 必选 | body | 选举人信息

- 选举人信息

参数名  | 类型 | 约束  | 提交方式 |描述
-------|------|------|---------|-----------
name | string | 必选 |body | 过期时间:202212042008,精确到秒;为方便后续测试，暂未做优化
sex | int | body | 必选 | 性别:0,女；1，男
age | int | body | 必选 | 可选 | 年龄
description | string | body | 可选 | 选举人描述

- example
```json
{
    "candidate": [
        {

            "name": "eric",
            "sex": 0,
            "age": 18,
            "description": "just for test"
        },
        {

            "name": "eric2",
            "sex": 0,
            "age": 19,
            "description": "just for test"
        },
        {

            "name": "eric3",
            "sex": 0,
            "age": 20,
            "description": "just for test"
        }
    ]
}
```

- 返回值
通用返回码

#### 投票
- 给一个候选人投票
- 请求方式:POST
- URL:/v1/electioncampaign/{ecId}/vote

- 投票人信息信息

参数名  | 类型 | 约束  | 提交方式 |描述
-------|------|------|---------|-----------
username | string | 可选 |body | 姓名,255以内
email | string | body | 必选 | 投票人email地址,用于接收投票结果信息
IdentityNo | string | body | 必选 | 可选 | 香港身份证号，仅做格式校验,未做有效性校验
CandidateId | int | body | 必选 | 选举人Id

- example
```json
{
    "username": "Eric",
    "email": "eric@test.com",
    "IdentityNo": "F123456(8)",
    "CandidateId": 4
}
```

- 返回值
- 选举活动信息

参数名  | 类型 | 约束   |描述
-------|------|-------|-----------
ecId   | int  | 必选  | 选举活动Id
finished | bool | 必选 | 是否结束,false 未结束
expire | string | 必选 | 过期时间,为节省时间，精确到秒
startTime | string | 必选 |  开始时间
finishTime | string | 必选 |  结束时间
description | string | 必选 |  选举描述
created | string | 必选 | 创建时间
candidates | object array | 必选 | 选举人信息

- 候选人信息

参数名      | 类型     | 约束  |  描述
-----------|----------|------|---------
candidateId| int      | 必选 | 所属选举活动Id
name       | string   | 必选 | 投票人姓名
sex        | int      | 必选 | 性别:0,女；1，男
age        | int      | 必选 | 年龄
description | string  | 可选 | 选举人描述
voteNumber  | int     | 必选 | 获得投票数

- example
```json
{
    "code": 200,
    "message": "successful",
    "data": {
        "ecId": 2,
        "finished": false,
        "expire": "2022-12-05T05:08:11+08:00",
        "startTime": "2022-12-04T14:41:27+08:00",
        "finishTime": "0001-01-01T00:00:00Z",
        "description": "test",
        "created": "0001-01-01T00:00:00Z",
        "voteNumber": ,
        "candidates": [
            {
                "candidateId": 4,
                "name": "test2",
                "sex": 1,
                "age": 18,
                "description": "just for test2",
                "voteNumber": 1
            },
            {
                "candidateId": 5,
                "name": "test3",
                "sex": 0,
                "age": 19,
                "description": "just for test4",
                "voteNumber": 1
            }
        ]
    }
}
```

#### 投票:查询选票
- 查询候选人选票，默认每页显示10票
- 请求方式:GET
- URL:/v1/electioncampaign/{ecId}/condidate/{condidateId}/votes?offset=1

- 入参

参数名  | 类型   | 约束  | 提交方式 |描述
-------|--------|------|---------|-----------
offset | string | 可选  |query   | 翻页,第几页

- 返回值

- 投票人信息

参数名       | 类型     | 约束  |描述
------------|---------|------|---------
username    | string  | 可选  | 姓名,长度255以内
email       | string  | 必选  | 投票人email地址,用于接收投票结果信息，选举活动内唯一，长度255以内
IdentityNo  | string  | 必选  | 可选 | 香港身份证号A123456(7)，仅做格式校验,未做有效性校验，选举活动内唯一
CandidateId | int     | 必选  | 投票的候选人Id


```json
{
    "code": 200,
    "message": "successful",
    "data": [
        {
            "username": "Eric",
            "email": "test3@qq.com",
            "identityNo": "C123456(7)",
            "candidateId": 1
        },
        {
            "username": "Eric",
            "email": "a_test3@1.cc",
            "identityNo": "C124456(8)",
            "candidateId": 1
        },
        {
            "username": "BeJson",
            "email": "test2@qq.com",
            "identityNo": "D987654(7)",
            "candidateId": 1
        },
        {
            "username": "BeJson",
            "email": "test@qq.com",
            "identityNo": "A123456(7)",
            "candidateId": 1
        },
        {
            "username": "Eric",
            "email": "test3@qq.com",
            "identityNo": "C123456(7)",
            "candidateId": 1
        },
        {
            "username": "Eric",
            "email": "test3@qq.com.cn",
            "identityNo": "F123456(7)",
            "candidateId": 1
        },
        {
            "username": "Eric",
            "email": "a_test3@qq.com",
            "identityNo": "C123456(8)",
            "candidateId": 1
        },
        {
            "username": "Eric",
            "email": "a_test3@1.com",
            "identityNo": "C123456(9)",
            "candidateId": 1
        },
        {
            "username": "Eric",
            "email": "a_test3@1.cc",
            "identityNo": "C124456(8)",
            "candidateId": 1
        },
        {
            "username": "Eric",
            "email": "a_test3@1.c",
            "identityNo": "C124456(9)",
            "candidateId": 1
        }
    ]
}
```









