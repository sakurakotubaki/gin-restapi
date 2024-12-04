# 認証のリクエストボディ

### signup

Request Body:

```json
{
    "email": "fuga@co.jp",
    "password": "1234qw"
}
```

Response Body:

```json
{
    "message": "User created successfully",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZ1Z2FAY28uanAiLCJleHAiOjE3MzMzMDc1MjUsInVzZXJJZCI6MX0.PxzToGE7OKGWN4HuLLcuUvnBxmNugJJPuoeMAJzKJPE",
    "user": {
        "email": "fuga@co.jp",
        "id": 1
    }
}
```

### login

Request Body:

```json
{
    "email": "fuga@co.jp",
    "password": "1234qw"
}
```

Response Body:

```json
{
    "message": "User logged in successfully",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZ1Z2FAY28uanAiLCJleHAiOjE3MzMzMDc3MzcsInVzZXJJZCI6MX0.EFi47prnkazxoqsimFhGpUty0pdLlDRzMUpNsbPHKjg"
}
```