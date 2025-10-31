# Kafka Producer Benchmarks (Go)

- 目标：基于不同 batch.size、linger.ms、compression、并发度评测吞吐与延迟
- 输出：p50/p95/p99 延迟，吞吐（records/s, MB/s）
- 可扩展：支持多 Broker、多 Topic、多分区测试
