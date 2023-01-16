[TOC]
## 用户登录

### 请求地址
>content-type	application/grpc
>pb.UserDetailReq

### 请求参数

**入参描述**

|tag key| 描述    | 有效范围       | 示例     |
|:----    |:------|:-----------|--------|
|authType | string | 来源  system | system |
|authKey | string | 登录名称       | admin  |
|password | string | 登录密码       | 123456 |

**入参示例**
```json
{
  "authType": "system",
  "authKey": "admin",
  "password": "123456"
}
```

### 返回参数

|tag key| 描述    | 有效范围    | 示例     |
|:----    |:------|:--------|--------|
|accessToken | string | token   |  |
|refreshToken | string | 刷新token |   |

```json
{
  "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsImlzcyI6ImdvLW1pbmRvYyIsImV4cCI6MTY3Mzc4NDI0M30.9euIuvsfKqBsX1xCGQxgoXDror3Lng3GQF8g1VMqfSI",
  "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsImlzcyI6ImdvLW1pbmRvYyIsImV4cCI6MTY3Mzc4Nzg0M30.klQ4gnKLCjacEP1mBEc3PQ_bzRDzx5BWQgyw9p6-Nsg"
}
```




## 获取用户信息

### 请求地址
>content-type	application/grpc
>pb.UserDetailReq

### 请求参数

**入参描述**

|tag key|描述| 有效范围    | 示例  |
|:----    |:---|:--------|-----|
|userId |int64 | 用户id    | 1   |

**入参示例**
```json
{
  "userId": "1"
}
```
### 返回参数
```json
{
  "userBasic": {
    "userId": "1",
    "username": "admin",
    "nickName": "admin",
    "phone": "13818888888",
    "roleId": "1",
    "avatar": "",
    "sex": "1",
    "email": "12@qq.com",
    "remark": "",
    "status": "2",
    "authKey": "",
    "authType": ""
  }
}
```

