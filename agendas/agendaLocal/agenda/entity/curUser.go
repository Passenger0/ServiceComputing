package entity

import (
	//"fmt"
	"io"
	"os"
	"log"
	"bufio"
	//"encoding/json"
	//"agenda/service"
)

type curUserInfo struct{
	UserName string	
	Password string
	Email string
	Tel string 
	Islogin bool 	
}

var curUser curUserInfo
var curUserFile string = os.Getenv("GOPATH") + "/src/agenda/data/curUser.txt"
//var curUserFile string = os.Getenv("GOPATH") + "/src/agenda/data/curUser.txt"

func init() {
	curUser.readFromFile()
}

func GetCurUser() *curUserInfo{//获取当前用户信息
	return &curUser
}
func (c *curUserInfo) GetLoginStatus() bool {
	return c.Islogin
}
func checkError(err error) {//处理错误
	if err != nil {
		log.Fatal(err)
	}
}
func (c *curUserInfo)LogIn(username,password,email,tel string) {//用户登录时对curUser和curUser.txt更新
	//email := service.AllUser.GetUserByName(username).GetEmail()
	//tel := service.AllUser.GetUserByName(username).GetTel()
	c = &curUserInfo{
		UserName : username,
		Password : password,
		Email : email,
		Tel : tel,
		Islogin  : true,
	}
	c.writeToFile()
}

func (c *curUserInfo)LogOut() {//用户登出时对curUser和curUser.txt更新
	c = &curUserInfo{
		UserName : " ",
		Password : " ",
		Email : " ",
		Tel : " ",
		Islogin  : false,
	}
	//c.Islogin = false
	c.writeToFile()
}

func (c *curUserInfo)writeToFile() {//将curUser写入文件curUser.txt
	fwrite, err := os.OpenFile(curUserFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	checkError(err)
	
	//写入curUser.txt文件
	var logStatus string
	if c.Islogin {
		logStatus = "True"
	} else {
		logStatus = "False"
	}
	//data := c.UserName + "\n" + c.Password + "\n" + c.Email+"\n" + c.Tel+ "\n"+ logStatus + "\n"
	data := c.UserName + "\n" + c.Password + "\n" + c.Email+"\n" + c.Tel+ "\n"+ logStatus+"\n"
	_, err = fwrite.WriteString(data)

	checkError(err)//检测是否成功写入
	fwrite.Close()
}

func (c *curUserInfo) read(fread io.Reader){

	buffer := bufio.NewReader(fread)//存储已读取数据的缓冲区

	user , err := buffer.ReadString('\n')
	checkError(err)
	c.UserName = user[:len(user)-1]//末尾有一个换行符，所以需要减一
	//fmt.Println("username: ",c.UserName)

	pass , err := buffer.ReadString('\n')
	checkError(err)
	c.Password = pass[:len(pass)-1]
	//fmt.Println("password: ",c.Password)

	email , err := buffer.ReadString('\n')
	checkError(err)
	c.Email = email[:len(email)-1]
	//fmt.Println("email: ",c.Email)

	tel , err := buffer.ReadString('\n')
	checkError(err)
	c.Tel =  tel[:len(tel)-1]
	//fmt.Println("tel: ",c.Tel)

	logStatus,err := buffer.ReadString('\n')
	checkError(err)
	if (logStatus == "True\n"){
		c.Islogin = true
	}else {
		c.Islogin = false
	}
	//fmt.Println("isLogin: ",c.Islogin)
}

func (c *curUserInfo)readFromFile() {
	//判断文件是否存在
	_, err := os.Stat(curUserFile)
	if os.IsNotExist(err) {
		return 
	}
	fread, err := os.OpenFile(curUserFile, os.O_RDONLY, 0755)
	checkError(err)
	c.read(fread)//读入数据
	fread.Close()
}