# 分布式锁实现原理与最佳实践 
### 分布式锁应用场景

很多应用场景是需要系统保证幂等性的（如api服务或消息消费者），并发情况下或消息重复很容易造成系统重入，那么分布式锁是保障幂等的一个重要手段。

另一方面，很多抢单场景或者叫交易撮合场景，如dd司机抢单或唯一商品抢拍等都需要用一把“全局锁”来解决并发造成的问题。在防止并发情况下造成库存超卖的场景，也常用分布式锁来解决。

### 实现分布式锁方案

这里介绍常见两种：redis锁、zookeeper锁

### 阅读全文链接
[分布式锁实现原理与最佳实践](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247483673&idx=1&sn=233c609a71fe8d0e8e3a0b5db920a7cc&chksm=e8215e09df56d71fa7be052174b9014ee71f480715b4685fcc8f1453937c36523015cb6d96d6&token=1061787983&lang=zh_CN#rd)

扫码关注微信订阅号支持：

![技术岁月](https://raw.githubusercontent.com/skyhackvip/ratelimit/master/techyears.jpg)
