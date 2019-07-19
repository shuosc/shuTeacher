# API Reference

## 模型

### 教师基本信息

```json
{
  "id":   教师证号,
  "name": 姓名
}
```

## web api

- `GET /ping`

  检查服务是否可用，应该直接返回`pong`。

- `GET /teacher?id=[教师证号]`

  获得教师信息
