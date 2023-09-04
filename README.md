# Douyin-Demo

## 抖音项目服务端简单示例

具体功能内容参考飞书说明文档

需要连接MySQL以及Minio以读取用户数据和视频文件

运行

```shell
go build && ./Douyin-Demo
```

### 功能说明

可用fresh 直接运行

* 用户登录数据保存在mysql数据库中
* 默认头像与背景保存到本地 public 目录中，访问时用 127.0.0.1:8080/static/image_name 即可
* 上传视频与封面保存在minio服务器中

### 测试
2023.08.25 未启用测试功能

test 目录下为不同场景的功能测试case，可用于验证功能实现正确性

其中 common.go 中的 _serverAddr_ 为服务部署的地址，默认为本机地址，可以根据实际情况修改

测试数据写在 demo_data.go 中，用于列表接口的 mock 测试

### 当前问题
1. (已解决)视频发布列表：POST给的 user_id 一直为0 （原因：json标签没有和接口文件的对应，导致解析失败）
2. 发布完新的视频后，返回到 publish/list 页面 workCount 数据不更新，需要退出登录以后再重新登录才可以，但是从feed点击视频发布者头像查看workCount没问题；同样的路由给POSTman发送，得到的是正确的workCount
3. (模糊)feed 页面，[官网描述](https://bytedance.feishu.cn/docx/BhEgdmoI3ozdBJxly71cd30vnRc) latest_time 为最新时间戳，而 next_time 为返回视频中发布时间最早的，next_time 作为下次请求的 latest_time，是不是矛盾？
4. 消息页面不显示历史聊天记录，使用Postman测试可以正常返回聊天记录列表；【可能原因：1. 目前返回的是string类型，可能应该返回int类型？2. msgType还是msg_type】

### 后续改进
1. 进行测试
2. 上redis
3. docker部署