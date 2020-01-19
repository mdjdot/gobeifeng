/*
	Redis
	redis是nosql数据库
	Remote dictionary server性能非常高，单机能够达到15万 qps
	redis适合做缓存，也可以做持久化

	Redis 命令
	redis-cli // 启动redis客户端
	redis-cli -h ip -p 6379 -a password --raw
	select 0 // redis默认有16个数据库，编号0~15
	keys * // 查看所有key
	set name zhangsan // 添加或修改key
	setex name 10 zhangsan // 设置有过期时间的key
	mset name zhangsan age 14 // 批量设置key value
	hset user name zhangsan
	hmset user gender man jon coder
	expire name 10 // 10秒后过期
	get name // 获取value
	mget name age // 批量获取value
	hget user name
	hmget user name age
	hgetall user // 获取user的字段和值
	del name // 删除key
	exists name user // key是否存在
	hexists user age // key是否存在字段
	hlen user // key有几个字段
	dbsize // 查看key总数
	flushdb // 清空当前数据库
	flushall // 清空所有数据库

	Redis 数据类型
	String // value最大是512M
	set  zhangsan zhangsanvalue
	get zhangsan
	mset  lisi lisivalue wangwu wangwuvalue
	mget lisi wangwu
	Hash // 适合存储对象
	hset user name zhangsan
	hmset user age 16 job coder
	hget user name
	hgetall user
	hmget user nam age
	List // 有序列表，可以从首尾添加数据
	lpush name zhangsan lisi // lisi zhangsan
	rpush name wangwu // lisi zhangsan wangwu
	lrange name 0 3 // lisi zhangsan wangwu
	lpop name // lisi        // zhangsan wangwu // 如果value全部移除，key自动删除
	rpop name // wangwu        //zhangsan
	del name
	Set
	set是string类型的无序集合，元素的值也不能重复
	sadd email 123@123.com 456@456.com // 添加元素
	smembers email // 查看元素
	sismember email 123@123.com 是否存在元素
	srem eamil 123@123.com // 删除元素
	Zset
	zadd names 1 zhangsan 2 lisi 3 wangwu
	zrange names 0 2
	zrange names 0 2 withscores
	zrank names lisi // 返回元素索引
	zrem name lisi // 移除元素
*/

package main

import "github.com/garyburd/redigo/redis"

import "fmt"

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conn.Do("AUTH", "dmtest")
	conn.Do("SET", "a", 123)
	a, err := redis.Int(conn.Do("GET", "a"))
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
}
