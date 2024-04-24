# soybean-admin-go



post auth/login

```json
{
  "userName":"Admin",
  "password":"123456"
}
```
```json
{
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjpbeyJ1c2VyTmFtZSI6IkFkbWluIn1dLCJpYXQiOjE2OTg0ODQ5NzIsImV4cCI6MTczMDA0NDc5OSwiYXVkIjoic295YmVhbi1hZG1pbiIsImlzcyI6IlNveWJlYW4iLCJzdWIiOiJBZG1pbiJ9.rLqWqgErEAgX4EVy_Kl_Eb1_bDmXyiZ9bZaupVgyv5M",
        "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjpbeyJ1c2VyTmFtZSI6IkFkbWluIn1dLCJpYXQiOjE2OTg0ODQ5ODQsImV4cCI6MTc2MTU4MDc5OSwiYXVkIjoic295YmVhbi1hZG1pbiIsImlzcyI6IlNveWJlYW4iLCJzdWIiOiJBZG1pbiJ9.QLrSKVlFXAzBb3v0BFzyxBbzrVMg58SF9oLS46Z8bFI"
    },
    "code": "0000",
    "msg": "请求成功"
}
```
以下接口 header 必带

Authorization : Bearer token

get auth/getUserInfo

```json
{
    "data": {
        "userId": "2",
        "userName": "Admin",
        "roles": [
            "R_ADMIN"
        ],
        "buttons": [
            "B_CODE2",
            "B_CODE3"
        ]
    },
    "code": "0000",
    "msg": "请求成功"
}
```

get systemManage/getUserList

```json
{
  "current": 1,
  "size": 10,
  "status": 1,
  "userName": "xxx",
  "userGender": 1,
  "nickName": "nick",
  "userPhone": "13522224444",
  "userEmail": "123@qq.com"
}
```
```json
{
    "data": {
        "records": [
            {
                "id": 1,
                "createBy": "Laura Perez",
                "createTime": "1970-02-13 09:04:15",
                "updateBy": "Susan Lee",
                "updateTime": "2000-01-27 05:31:57",
                "status": "1",
                "userName": "qYdjYFvjN3863c",
                "userGender": "2",
                "nickName": "Jeffrey Clark",
                "userPhone": "13718351528",
                "userEmail": "w.dbxb@wbgo.info",
                "userRoles": [
                    "R_JLQT_LLPK",
                    "R_XVO_LIH",
                    "R_XBU_XKW",
                    "R_HVAV_CBS",
                    "R_HEGQ_VLRH"
                ]
            }
        ],
        "current": 1,
        "size": 10,
        "total": 200
    },
    "code": "0000",
    "msg": "请求成功"
}
```