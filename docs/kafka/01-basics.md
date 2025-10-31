# Kafka 基础

- 核心概念：Topic、Partition、Broker、Producer、Consumer、Consumer Group
- 消息投递语义：at-most-once、at-least-once、exactly-once（EOS）
- 存储机制：日志分段、索引、压缩（log compaction）
- ISR 机制与副本同步

## 快速上手
- 启动本地集群（Kraft 模式）
- 创建 Topic：分区数与副本因子
- 生产与消费示例
