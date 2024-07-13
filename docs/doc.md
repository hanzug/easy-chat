### 1. 项目业务

在项目中我们主要业务为一个im聊天服务，根据该项目业务可分为三个核心业务

- 用户业务
  - 用户登入、注册、详情、查找
- 社交【好友、群】
  - 好友：好友添加、列表等
  - 群：进群，退群，列表等
- 聊天：
  - 私聊、群聊、聊天记录等



### 2. 服务架构

在项目中我们将根据如上的三个核心业务，构建出相应的微服务来实现功能，具体架构如下图：

![图片描述](https://img.mukewang.com/wiki/65eb231009010dbf06320376.jpg)
此图为我们将构建的系统架构图，请求会通过api网关，访问到BFF层即对外提供的各个api接口服务，通过BFF请求多个rpc服务将请求数据进行聚合再返回给调用的客户端。

另外值得注意的是在服务的架构中实际上还有一个服务及`im的websocket聊天服务`，当然也有异步任务的`消息队列服务`，分别处理聊天和任务功能。

**关于BFF的理解：理解为在传单体项目架构设计中的**controller，在传统单体架构的设计中往往会一般会分为controller、service、model。而controller在对外提供接口服务的时候往往会调用多个service来实现需求。我们所构建的微服务服务可以理解就是将其中的service抽出来然后形成服务化，而BFF层则可以理解为是controller抽出来形成的服务。

![图片描述](https://img.mukewang.com/wiki/65eb231d097b1ac701870168.jpg)
BFF层在微服务的架构中是帮助客户端在进行请求服务的时候，进行数据的聚合处理，当然也会承担权限鉴权验证等工作。

然后在项目的构建开发中会基于go-zero分别开发好api与rpc服务、再基于docker部署，对整个系统也会设置相关的日志及系统指标监听。



### 3. 项目过程

项目的开发，我们将先从用户服务开始开发，然后再完成社交服务最后构建好im服务的功能。