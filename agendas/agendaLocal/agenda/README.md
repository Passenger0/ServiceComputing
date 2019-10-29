本地版agenda，请将agenda文件移至\$GOPATH/src目录下。

# 可运行命令

 - **agenda -h ：列出程序说明**
 - **agenda \<command> -h :列出命令说明**
 - **agenda register -u username –p password –e email -t telphone ：用户注册**
 - **agenda login -u usename -p password：用户登录**
 - **agenda logout：用户登出**
 - **agenda qryuser：查询所有用户信息**
 - **agenda deluser -p password：注销当前账号**

# 运行示例

1. **agenda -h**

运行go install将其安装到$GOPATH/bin，然后运行agenda（同agenda -h），运行结果如下：

```go
F:\Go\gowork\src\agenda>agenda
A simple version of agenda,a meeting-manage system.

Usage:
  agenda [command]

Available Commands:
  deluser     delete the current user from the system
  help        Help about any command
  login       User login an account
  logout      user logout the account
  qryuser     query all the users
  register    Used to register an account

Flags:
      --config string   config file (default is $HOME/.agenda.yaml)
  -h, --help            help for agenda
  -t, --toggle          Help message for toggle

Use "agenda [command] --help" for more information about a command.
subcommand is required
```
2. **agenda \<command> -h** ：以agenda register - h 为例：

```go
F:\Go\gowork\src\agenda>agenda register -h
Use the register command in one of the forms below to register an account:
        1. agenda register -u username -p password -e email -t telephone
        2. agenda register -uusername  -ppassword  -eemail  -ttelephone

        Flags:
                -u username
                -p password
                -e email
                -t telephone

Usage:
  agenda register [flags]

Flags:
  -e, --email string   email address
  -h, --help           help for register
  -p, --pass string    password
  -t, --tel string     telephone number
  -u, --user string    username

Global Flags:
      --config string   config file (default is $HOME/.agenda.yaml)
```

3. **agenda register -u username –p password –e email -t telphone** ：用户注册（必须提供四个flag的值）

下面的命令注册了一个用户名为fsql，密码为123，邮箱为fsq@mail.com，电话为12345678的用户
```go
F:\Go\gowork\src\agenda>agenda register -u fsql -p 123 -e fsq@mail.com -t 12345678
Info: 2019/10/29 16:44:38 fsql  register succeed!
```
假如用户名已存在：

```go
F:\Go\gowork\src\agenda>agenda register -u fsq -p 123 -e fsq@mail.com -t 12345678
Error: 2019/10/29 16:44:03 fsq  register failed: username has been used!
```
注：Info与Error在输出到标准输出的同时也会写入info.log和error.log（可在agenda/data/目录中查看）】

4. **agenda login -u usename -p password**：用户登录

```go
F:\Go\gowork\src\agenda>agenda login -u fsql -p 123
Info: 2019/10/29 16:50:38 fsql  login succeed!
```
当前登录的用户信息会存在与agenda/data/curUser.txt中：

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191029165207534.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0V4Y2Vm,size_16,color_FFFFFF,t_70)

5. **agenda qryuse**r：在登陆之后查询所有用户信息


```go
F:\Go\gowork\src\agenda>agenda qryuser
There are  3  users：
Name--Email--Telephone
fsq   test@qq.com   12345678
fsq2   test@qq.com   12345678
fsql   fsq@mail.com   12345678
Info: 2019/10/29 16:53:01 fsql  qryuser succeed!
```
6. **agenda logout：**用户登出

```go
F:\Go\gowork\src\agenda>agenda logout
Info: 2019/10/29 16:54:26 fsql  logout succeed!
```

7. **agenda deluser -p password**：注销当前账号（在注销前需要保证用户已登录）

```go

F:\Go\gowork\src\agenda>agenda deluser 
Error: 2019/10/29 16:56:05 fsql  deluser failed: password must be provided

F:\Go\gowork\src\agenda>agenda deluser -p 123
Info: 2019/10/29 16:56:13 fsql  deluser succeed!
```