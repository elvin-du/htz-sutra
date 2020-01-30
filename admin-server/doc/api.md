黄庭禅经典听读系统管理后端API设计文档
===
[TOC]

# 1. 介绍
因为用户量不大，为了加快开发速度，并易于以后的维护，所以没有采用微服务架构，而是采用把所有服务集成在一个程序中。

* 数据库： MongoDB
* 整个程序从逻辑上分为以下三个主要服务：
	+ 专辑文件管理： 专辑文件上传，删除，专辑信息的编辑，更新等。
	+ 系统管理员账户管理：系统后台管理员权限的管理等。
	+ 搜索服务：负责提供系统的搜索服务。
* 接口约定
	+ 所有的API都是通过POST方法进行调用。
	+ 除了文件服务之外，其余所有的入参和出参都是采用JSON格式。
	+ 通过在URI的第一个参数来表示动作。
		- get：获取；例如：/get/user（获取用户信息）
		- delete：删除；例如：/delete/user（删除用户）
		- post：添加；例如：/post/user（添加用户）
		- put：更新；例如：/put/user（更新用户信息）
	+ 字段命名方式: 单词使用`_`作分隔，而非驼峰方式。

# 2. 公共数据结构


* **Response** 统一的返回样式。**以下接口返回都只写了成功时的data字段。**
```
{
	int code - 执行结果。200 - 成功，其余为失败
	string err - 执行错误结果描述
	string data - json数据结构的执行结果
}
```
* **SutraInfo** 经典信息
```
{
	string id - 唯一标识符
	string name - 经典名称
	string cover - 封面文件ID
	string description - 经典简介
	uint64 played_count - 总共播放次数
	uint64 item_total - 一共有多少集
	string created_at - 创建时间
}
```
* **SutraItem** 每集经典的详情
```
{
	string id - 唯一标识符
	string sutra_id - 经典专辑唯一标识符
	string title - 标题
	string description - 介绍
	string original - 经典原文，存储在数据库，而不是在文件中
	string audio_id - 音频文件id
	string lyric_id - 歌词文件id，包含张讲师所有的话
	uint64 lesson - 第几集
	uint64 played_count - 播放次数
	uint64 duration - 时长
	string created_at
	string hash - 文件的sha256哈希值
}
```

# 4. 文件服务

## 4.1. 上传
* URI：/post/file/upload
	
* Inputs 按照提交表单的方式进行提交
    + mime: 文件类型。请按照http协议规定的文件类型名填写。把其中/转换成-。例如：image/jpeg变成 image-jpeg
	+ file_hash: 文件sha1值，160位，20字节。十六进制表示，用于检查文件是否完整。如果有已经有相同hash值的文件，不会重复存储，会引用同一个文件。
	+ sutra_name: 专辑名称
	+ item_number：专辑项目编号
	+ item_name： 专辑项目名称
	+ item_suffix: 专辑项目文件后缀名
	+ file：专辑项目文件。


* Outputs
```
{
	string file_id - 文件ID。id为空的话，表示上传失败
}
```

* 例子
```
curl -X POST http://localhost:9001/post/file/upload/abc123/application-json \
  -F "file=@./aa.txt" \
  -F "sutra_name=论语" \
  -F "item_number=3" \
  -F "item_name=不食嗟来之食" \
  -F "item_suffix=mp3" \
  -F "file_hash=a8bc8" \
  -F "mime=image-jpeg" \
  -H "Content-Type: multipart/form-data"
```

## 4.2. 下载
* URI：/get/file/download
	+ 下载图片时可以指定高度和宽度。TODO：第一期先不做，也就说只会下载原图
	+ TODO： 对常用大小的图片会进行缓存。
* Inputs
```
{
	string file_id - 文件ID
	int64 height - 下载图片时可以指定高度,单位像素
	int64 width - 下载图片时可以指定宽度,单位像素
}
```

* Outputs
```
文件字节流
```

* 例子
```
curl -H "Content-Type:application/json" -H "Data_Type:msg" -X POST \
--data '{"file_id":"0a8bc","heigth":800,"width":400}' \
http://localhost:9001/get/file/download > tmp.txt
```
