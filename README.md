# 黄庭禅经典听读系统


## API服务器（api-server）
### HTTP框架

采用的是gin开源框架。



## 后台管理服务器（admin-server）
### HTTP框架

采用的是gin开源框架。

## 数据库（mongodb）

## 1. 初始化数据库&表

```
//create database and user
use admin
db.createUser(
     {
         user:"htzsutraadmin",
         pwd:"htzsutra123",
         roles:[{role:"root",db:"admin"}]
    }
)

use htz_sutra 
db.createCollection("blockchain_status",{"capped": true, "size": 9999, "max": 1});
db.createCollection("blocks");

db.blocks.createIndex({"height": -1},{"unique": true});
db.transactions.createIndex({"tx_hash": -1}, {"unique": true});
db.transactions.createIndex({"block_height": -1});
db.transactions.createIndex({"from": 1});
db.validators.createIndex({"operator_address": 1}, {"unique": true});
```
