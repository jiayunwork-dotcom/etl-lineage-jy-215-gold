# etl-lineage

多源数据 ETL 血缘追踪与影响分析（CLI）。

读入一份血缘规格（每张表及其上游依赖），构建有向无环图（DAG），
支持：拓扑排序、上游/下游血缘追溯、以及「某张表变更会影响哪些下游表」的影响分析。

## 用法

```
etl-lineage -spec lineage.txt [-node dim_orders] [-dot -]
```

- `-spec`   血缘规格文件，`-` 表示标准输入；每行 `目标 <- 依赖1,依赖2`
- `-node`   计算该节点的下游影响集合（每行一个，排序输出）；不指定则输出 DOT 图
- `-dot`    DOT 图输出，`-` 表示标准输出

## 规格格式

```
dim_orders   <- stg_orders
fact_sales   <- dim_orders, dim_customers, stg_sales
report_kpi   <- fact_sales
```
