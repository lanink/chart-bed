# 图片服务

支持RIAPI的图片服务

## 使用

1. 运行服务
   > go run ./src/main.go
2. 上传图片
   > `curl --location --request POST '127.0.0.1:8080/upload' 
   --form 'file=@"/path/to/file.jpg"'` 
3. 查看
   > 原图：http://127.0.0.1:8080/images/1.jpg 
   > 设置大小: http://127.0.0.1:8080/images/1.jpg?w=400
   

## 更多信息

[RIAPI参数信息](https://docs.imageflow.io/querystring/introduction.html)

**目前仅简单支持w和h（即width和height）**




