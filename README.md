# Douyin-Demo

## 抖音项目服务端简单示例
### 演示

使用ngrok内网穿透的演示地址: https://flowing-remarkably-worm.ngrok-free.app

### 部署
需要连接MySQL以及Minio以读取用户数据和视频文件

运行

```shell
docker-compose up
```

**运行前，在config.yaml中修改为本机对应的ip地址**

config.yaml 样例如下：

```shell
mysql:
  user: root
  password: mysql123
  ip: 127.0.0.1
  port: 3306
  database: douyin
minio:
  endpoint: 192.168.31.246:9000
  accessKeyID: minioadmin
  secretAccessKey: minioadmin
  videoBucket: douyin-video
  imageBucket: douyin-image
```

其中可能会遇到的错误：
1. wait-for.sh 无权限，解决方法：在项目路径下赋予其权限 `sudo chmod 777 wait-for-it.sh`

### 功能说明

* 用户登录数据保存在mysql数据库中
* 默认头像与背景为随意找的可访问url
* 上传视频与封面保存在minio服务器中

### 测试
2023.08.25 未启用测试功能

test 目录下为不同场景的功能测试case，可用于验证功能实现正确性

其中 common.go 中的 _serverAddr_ 为服务部署的地址，默认为本机地址，可以根据实际情况修改

测试数据写在 demo_data.go 中，用于列表接口的 mock 测试

### 当前问题
1. (已解决)视频发布列表：POST给的 user_id 一直为0 （原因：json标签没有和接口文件的对应，导致解析失败）
2. 发布完新的视频后，返回到 publish/list 页面 workCount 数据不更新，需要退出登录以后再重新登录才可以，但是从feed点击视频发布者头像查看workCount没问题；同样的路由给POSTman发送，得到的是正确的workCount
3. (模糊)feed 页面，[官网描述](https://bytedance.feishu.cn/docx/BhEgdmoI3ozdBJxly71cd30vnRc) latest_time 为最新时间戳，而 next_time 为返回视频中发布时间最早的，next_time 作为下次请求的 latest_time，描述矛盾
4. (半解决)消息页面一直在刷新，请求的GET中pre_msg_time作用令人疑惑;**由于页面一直在刷新，而且读取的消息记录列表中的最后一条记录的创建时间被设定为新的请求的pre_msg_time，因此，我们在查找sql中的聊天记录时只返回最早发出的消息，随着客户端的循环请求，逐步按照时间显示出所有的聊天记录，这样得到的聊天记录页面比较符合我们对于聊天功能的认知**

### 后续改进
1. (已完成)docker部署
2. 进行测试
3. 上redis

### 青训营汇总文档
[[小试牛刀]青训营大项目答辩汇报文档](https://kvalmttjdc5.feishu.cn/docx/QyF7du0RdoID9Jx3DMic2Cqfn9g)