# cloudgo
## 概述

开发简单 web 服务程序 cloudgo，了解 web 服务器工作原理。

环境：Windows10/curl-7.55.1/ApacheBench 2.3

工具：VS Code



## 框架选择

框架：**martini**

理由：

* 具有中文版使用文档，简单，上手快。
* 可与其他 Go 的包配合工作
* 可轻松添加工具



## 运行程序

运行程序，端口设置为8080时

```go
 go run main.go -p8080
[martini] listening on :8080 (development)
```

打开浏览器，输入（Passenger0为测试用户名）:

```go
http://localhost:8080/Passenger0
```

浏览器输出欢迎语句：

![1](assets/1.png)

cloudgo输出提示信息：

```go
go run main.go -p8080
[martini] listening on :8080 (development)
[martini] Started GET /Passenger0 for [::1]:54035
[martini] Completed 200 OK in 995.9µs
```



## 测试

### curl测试

运行cloudgo ：
```
go run main.go -p 8080
```

打开cmd，输入：
```
curl -v http://localhost:8080/Passenger0
```

-----------

输出：

```go
curl -v http://localhost:8080/Passenger0
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> GET /Passenger0 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.55.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Mon, 11 Nov 2019 13:16:00 GMT
< Content-Length: 17
< Content-Type: text/plain; charset=utf-8
<
Hello Passenger0
* Connection #0 to host localhost left intact
```



### ab测试

运行cloudgo ：

```
go run main.go -p 8080
```

打开cmd，输入：

```
ab -n 1000 -c 100 http://localhost:8080/Passenger0
```

------

输出：

```go
$ab -n 1000 -c 100 http://localhost:8080/Passenger0
This is ApacheBench, Version 2.3 <$Revision: 1843412 $> 
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/ 
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests 
Completed 200 requests 
Completed 300 requests 
Completed 400 requests 
Completed 500 requests 
Completed 600 requests 
Completed 700 requests 
Completed 800 requests 
Completed 900 requests 
Completed 1000 requests 
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /Passenger0
Document Length:        17 bytes

Concurrency Level:      100
Time taken for tests:   1.146 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      134000 bytes
HTML transferred:       17000 bytes
Requests per second:    872.65 [#/sec] (mean)
Time per request:       114.593 [ms] (mean)
Time per request:       1.146 [ms] (mean, across all concurrent requests)
Transfer rate:          114.19 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.4      0       1
Processing:    24  108  20.9    113     145
Waiting:        1  108  21.2    113     144
Total:         25  109  20.9    113     145

Percentage of the requests served within a certain time (ms)
  50%    113
  66%    118
  75%    121
  80%    124
  90%    129
  95%    131
  98%    135
  99%    136
 100%    145 (longest request)

```

#### ab测试参数解读：

Server Software:                        #服务器使用的软件，即响应头的Server字段

Server Hostname:                        # 服务器主机名，即请求头的 Host 字段

Server Port:                                 #服务器请求端口

Document Path:                            # 文档路径，即请求头的请求路径

Document Length:                           #文档长度，即响应头的 Content-Length 字段

Concurrency Level:                          # 并发等级，即每次的并发数

Time taken for tests:                       #测试花费的时间

Complete requests:                       # 完成的请求数

Failed requests:                              #失败的请求数

Total transferred:                             #总传输字节数

HTML transferred:                           # HTML 报文总传输字节数

Requests per second:                      # 平均每秒的请求数

Time per request:                              # 平均每个请求花费的时间

Time per request:                              # 平均每个请求花费的时间（考虑并发）

Transfer rate:                                     #平均每秒传输的千字节数

Connection Times (ms)                      #传输时间统计

Connect:                                            # 连接时间

Processing:                                        #处理时间

Waiting:                                             # 等待时间

Total:                                                 #总时间



Percentage of the requests served within a certain time (ms)            #一定时间内服务了的请求数所占的百分比

  50%    113                                         #113毫秒内服务了50%的请求

··· ···

 100%    145 (longest request)          #145毫秒内服务了100%的请求(最长请求)