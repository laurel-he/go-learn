mysql索引优化
===
# 查询表
## 查询表有哪些字段
`desc table_name;`
## 查询建表语句
`SHOW CREATE TABLE table_name;`
## 查询表有哪些索引
`SHOW INDEX FROM table_name;`
# 删除索引

`ALTER TABLE table_name DROP INDEX index_name;`
# 普通索引
## 介绍
最基本索引，没有任何限制
## 语法
`CREATE INDEX index_name on tabl_name (column);`;
``