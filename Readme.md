---

module: hhyt/database/redisdb
function: 封装redisgo操作redis数据库（第二版）
version: 0.2.0
path: hhyt/database/redis@v0.2.0
---

目录

[TOC]



# 引用

## go.mod

replace 的物理磁盘位置要根据物理目录的实际位置给定。

```go
module lx

go 1.15

require (
    hhyt/database/redisdb v0.1.0
)

replace(
    hhyt/database/redisdb => "../../hhyt/database/redis@v0.2.0"
)
```

正式使用是也可以使用`go get github.com/tinybear1976/redisdb2`进行引用

## main.go

连接和基本操作。当引用程序开始时，可以进行多个连接的初始化，即New()，每个连接初始化后，连接池指针就被记录在模块内部，直到使用Destroy()，手动销毁所有的连接池指针。在进行后续的操作时，首先需要申请连接对象指针，申请成功后，可以通过这个连接指针，进行SET、GET、DEL、KEYS、HMSET、HMGET、HGETALL、EXISTS操作，单次或多次操作后，必须显式性调用Disconnect关闭该连接。如果彻底不需要连接，可以使用Destroy销毁所有的连接池。

```go
package main

import (
    "fmt"
    "hhyt/database/redisdb2"
)

func main() {
    //连接：在连接时第一个参数为连接标识，后面的顺序分别为redis服务的IP:Port，密码，数据库编号。
    redisdb2.New("local", "127.0.0.1:6379", "", 0)
    // 首先申请连接
    local,err:=redisdb2.Connect("local")
    //基本操作：每次操作的第一个参数为连接标识。
    redisdb2.SET(local, "dd", "1234")
    val, err := redisdb2.GET(local, "dd")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(val)
    redisdb2.Disconnect(local)
    //如果有必要，可以将redisdb里面记录的连接记录全部清除。
    redisdb2.Destroy()
}
```



# 函数

## New

建立并保存一个Redis连接，在之后的使用过程中，通过连接标识创建连接并进行数据库操作。

```go
func New(serverTag, server, password string, dbnum int)
```

入口参数：
| 参数名    | 类型   | 描述                                         |
| --------- | ------ | -------------------------------------------- |
| serverTag | string | Redis数据库连接标识                          |
| server    | string | Redis地址，格式  ip:port （端口一般为 6379） |
| password  | string | 密码，默认空                                 |
| dbnum     | int    | 数据库编号                                   |

返回值：无



## Destroy

销毁所有模块内保存的（由New创建的）连接池指针。

```go
func Destroy()
```

入口参数：无

返回值：无



## Connect

每个段落操作开始的第一个动作，申请连接。

```go
func Connect(serverTag string) (*redis.Conn, error)
```

入口参数：

| 参数名    | 类型   | 描述                |
| --------- | ------ | ------------------- |
| serverTag | string | Redis数据库连接标识 |

返回值：

| 返回变量 | 类型        | 描述                                      |
| -------- | ----------- | ----------------------------------------- |
|          | *redis.Conn | 连接对象指针，如果申请不成功，则为nil     |
|          | error       | 返回操作结果的错误信息，如果正确则返回nil |



## Disconnect

每个段落操作的最后一个动作，关闭连接。

```go
func Diconnect(conn *redis.Conn) (err error)
```

入口参数：

| 参数名 | 类型        | 描述         |
| ------ | ----------- | ------------ |
| conn   | *redis.Conn | 连接对象指针 |

返回值：

| 返回变量 | 类型  | 描述                                      |
| -------- | ----- | ----------------------------------------- |
|          | error | 返回操作结果的错误信息，如果正确则返回nil |



## BGREWRITEAOF

手动强制重写AOF。因为服务器是异步操作，所以只是发出指令，不返回任何操作结果。

```go
func BGREWRITEAOF(conn *redis.Conn)
```

入口参数：

| 参数名 | 类型        | 描述         |
| ------ | ----------- | ------------ |
| conn   | *redis.Conn | 连接对象指针 |

返回值：

无



## SET

设置一个键值对。字符串类型。**提示：使用后应该记得使用Close(conn)关闭**

```go
func SET(conn *redis.Conn, key, value string) error
```

入口参数：

| 参数名 | 类型        | 描述                    |
| ------ | ----------- | ----------------------- |
| conn   | *redis.Conn | Redis数据库连接对象指针 |
| key    | string      | key                     |
| value  | string      | value                   |

返回值：

| 返回变量 | 类型  | 描述                                      |
| -------- | ----- | ----------------------------------------- |
|          | error | 返回操作结果的错误信息，如果正确则返回nil |



## GET

根据一个键获得一个值。字符串类型。**提示：使用后应该记得使用Close(conn)关闭**

```go
func GET(conn *redis.Conn, key string) (string, error) 
```

入口参数：

| 参数名 | 类型        | 描述                    |
| ------ | ----------- | ----------------------- |
| conn   | *redis.Conn | Redis数据库连接对象指针 |
| key    | string      | key                     |

返回值：

| 返回变量 | 类型   | 描述                                      |
| -------- | ------ | ----------------------------------------- |
|          | string | 根据键返回对应的值                        |
|          | error  | 返回操作结果的错误信息，如果正确则返回nil |



## DEL

根据一个键获得一个值。字符串类型

```go
func DEL(conn *redis.Conn, keys ...interface{}) error  
```

入口参数：

| 参数名 | 类型          | 描述                         |
| ------ | ------------- | ---------------------------- |
| conn   | *redis.Conn   | Redis数据库连接对象指针      |
| key    | []interface{} | 可以一个key，也可以是多个key |

返回值：

| 返回变量 | 类型  | 描述                                      |
| -------- | ----- | ----------------------------------------- |
|          | error | 返回操作结果的错误信息，如果正确则返回nil |

示例1：

```go
func Test() {
  local,err:=redisdb2.Connect("local")
	redisdb2.SET(local, "dd", "1234")
	redisdb2.SET(local, "ee", "1234")
  //可以多个key进行同时删除
	redisdb2.DEL(local, "dd", "ee")
	redisdb2.Disconnect(local)
}
```

示例2：

```go
func Test() {
  local,err:=redisdb2.Connect("local")
	redisdb2.SET(local, "dd", "1234")
	redisdb2.SET(local, "ee", "1234")
	keys := []interface{}{"dd"}
	keys = append(keys, "ee")
	redisdb2.DEL(local, keys...)
	redisdb2.Disconnect(local)
}
```



## KEYS

查询某些特征的key是否存在。

```go
func KEYS(conn *redis.Conn, query string) ([]string, error)  
```

入口参数：

| 参数名 | 类型        | 描述                                                      |
| ------ | ----------- | --------------------------------------------------------- |
| conn   | *redis.Conn | Redis数据库连接对象指针                                   |
| query  | string      | 可以一个key，也可以是多个key。一般需要匹配的部分用 * 表示 |

返回值：

| 返回变量 | 类型     | 描述                                  |
| -------- | -------- | ------------------------------------- |
|          | []string | 返回所有符合特征的键名称              |
|          | error    | 操作错误信息，如果操作正确，则返回nil |
|          |          |                                       |

示例：

```go
func Test() {
  local,err:=redisdb2.Connect("local")
	redisdb2.SET(local, "id::1000", "1234")
	redisdb2.SET(local, "id::1001", "1234")

	keys, _ := redisdb2.KEYS(local, "id::*")
	fmt.Println(len(keys), keys)  //返回结果：    2 [id::1001 id::1000]
	keys, _ = redisdb2.KEYS(local, "ui::")
	fmt.Println(len(keys), keys)  //返回结果：    0 []
  redisdb2.Disconnect(local)
}
```



## HMSET

散列（map）设置。可以一次设置一个字段值，也可以设置多个字段值。

```go
func HMSET(conn *redis.Conn, params ...interface{}) error   
```

入口参数：

| 参数名 | 类型          | 描述                                                  |
| ------ | ------------- | ----------------------------------------------------- |
| conn   | *redis.Conn   | Redis数据库连接对象指针                               |
| params | []interface{} | 此处数据，第一个值应该是Key，后面字段名与值交替出现。 |

返回值：

| 返回变量 | 类型  | 描述                                      |
| -------- | ----- | ----------------------------------------- |
|          | error | 返回操作结果的错误信息，如果正确则返回nil |

示例：

```go
func Test() {
  local,err:=redisdb2.Connect("local")
  //添加一个字段
	redisdb2.HMSET(local, "id::1000", "name", "joe")
  //添加多个字段
	redisdb2.HMSET(local, "id::1001", "name", "nathan", "age", "24", "gender", "male")
  //添加多个字段
	fields := []interface{}{
		"id::1002",
		"name", "jean",
		"age", "18",
		"gender", "female",
	}
	redisdb2.HMSET(local, fields...)
  redisdb2.Disconnect(local)
}
```



## HMGET

散列（map）获取字段值。可以一次获取一个字段值，也可以获取多个字段值。

```go
func HMGET(conn *redis.Conn, params ...interface{}) ([]string, error) 
```

入口参数：

| 参数名 | 类型          | 描述                                          |
| ------ | ------------- | --------------------------------------------- |
| conn   | *redis.Conn   | Redis数据库连接对象指针                       |
| params | []interface{} | 此处数据，第一个值应为Key，后面全部是字段名。 |

返回值：

| 返回变量 | 类型     | 描述                                                         |
| -------- | -------- | ------------------------------------------------------------ |
|          | []string | 按照规定字段顺序，返回对应值，如果有空字符串，也有可能表示该字段根本就不存在。 |
|          | error    | 返回操作结果的错误信息，如果正确则返回nil                    |

示例：

```go
func Test() {
  local,err:=redisdb2.Connect("local")
  //返回一个字段的值
	v1, _ := redisdb2.HMGET(local, "id::1000", "name")
  //返回多个字段的值，其中height字段不存在，结果对应值为空字符串
	v2, _ := redisdb2.HMGET(local, "id::1001", "name", "age", "height")
  //返回多个字段的值
	fields := []interface{}{
		"id::1002",
		"name", "age",
	}
	v3, _ := redisdb2.HMGET(local, fields...)

	fmt.Print(v1, len(v1), "\n", v2, len(v2), "\n", v3, len(v3), "\n")
	redisdb2.Disconnect(local)
}
```



## HGETALL

按照Key返回map[string]string。这个函数主要用于按照指定Key返回某个散列类型的全部键值对。

```go
func HGETALL(conn *redis.Conn, key string) (map[string]string, error)
```

入口参数：

| 参数名 | 类型        | 描述                    |
| ------ | ----------- | ----------------------- |
| conn   | *redis.Conn | Redis数据库连接对象指针 |
| key    | string      | 指定查询的Key。         |

返回值：

| 返回变量 | 类型              | 描述                                      |
| -------- | ----------------- | ----------------------------------------- |
|          | map[string]string | 返回键值对。                              |
|          | error             | 返回操作结果的错误信息，如果正确则返回nil |

示例：

```go
func Test() {
  local,err:=redisdb2.Connect("local")
  //返回一个不存在的Key
  v1, _ := redisdb2.HGETALL(local, "id::2000")  //结果:  map[]
  //返回一个正确的Key
  v2, _ := redisdb2.HGETALL(local, "id::1001")  //结果:  map[age:24 gender:male name:nathan]

	fmt.Print(v1, "\n", v2, "\n")
	redisdb.Disconnect()
}
```



## EXISTS

检查某个Key是否存在。

```go
func EXISTS(conn *redis.Conn, key string) (bool, error)
```

入口参数：

| 参数名 | 类型        | 描述                    |
| ------ | ----------- | ----------------------- |
| conn   | *redis.Conn | Redis数据库连接对象指针 |
| key    | string      | 指定查询的Key。         |

返回值：

| 返回变量 | 类型  | 描述                                      |
| -------- | ----- | ----------------------------------------- |
|          | bool  | 存在返回true，否则返回false。             |
|          | error | 返回操作结果的错误信息，如果正确则返回nil |

示例：

```go
func Test() {
  local,err:=redisdb2.Connect("local")
  //测试返回一个不存在的Key，返回值为false
	v1, _ := redisdb2.EXISTS(local, "id::2000")
  //测试返回一个正确的Key，返回值为true
	v2, _ := redisdb2.EXISTS(local, "id::1001")

	fmt.Print(v1, "\n", v2, "\n")
  redisdb2.Disconnect(local)
}
```



