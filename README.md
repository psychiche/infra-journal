# infra-journal

面向基础设施工程的知识与实验仓库，包含：
- docs：Kafka/Redis 笔记与最佳实践
- labs：高并发与分布式实验（Go / Python）
- wheels：从零实现的迷你组件与封装
- scripts：开发环境脚本

## 目录结构

- docs/
  - kafka/
    - 01-basics.md
    - 02-production-best-practices.md
  - redis/
    - 01-data-structures.md
    - 02-cluster-mode.md
- labs/
  - high-concurrency/
    - kafka-producer-benchmarks/ (Go)
    - redis-pipeline-optimization/ (Python)
  - distributed-systems/
    - kafka-raft-simulation/ (Go)
    - redis-sentinel-monitor/ (Python)
- wheels/
  - mini-kafka/
    - go/ (核心代码)
    - design.md (架构设计)
  - vector-index/
    - go/ (Go 实现)
    - python/ (Python 封装)
- scripts/
  - setup-dev-env.sh

## 快速开始

1. 准备环境

```bash
scripts/setup-dev-env.sh
```

2. 运行示例

- Kafka 生产者基准：
```bash
cd labs/high-concurrency/kafka-producer-benchmarks && go run .
```

- Redis pipeline 测试：
```bash
cd labs/high-concurrency/redis-pipeline-optimization && python3 main.py
```

## 约定
- Go 版本：1.22+
- Python 版本：3.9+
- 本地需具备 Kafka/Redis（或使用容器）
