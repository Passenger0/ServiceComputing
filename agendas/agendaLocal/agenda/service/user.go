package service

import (
	"bufio"
	"os"
	"agenda/entity"
	"log"
	"encoding/json"
	//"fmt"
)

//用户信息链表
type UserList struct{
	Users []entity.UserData `json:"users"`
}


var AllUser UserList
var userfile string = os.Getenv("GOPATH") + "/src/agenda/data/users.json"

func init() {
	AllUser.readFromFile()
}

func checkError(err error) {//处理错误
	if err != nil {
		log.Fatal(err)
	}
}
//通过名字获得用户
func (u *UserList)GetUserByName(username string) *entity.UserData{
	for _, user := range u.Users {
		if user.Name == username {
			return &user
		}
	}
	return &entity.UserData{}
}

//获取所有用户
func (u *UserList)GetAllUsers() []entity.UserData{
	return u.Users
}

//用户是否存在
func (u *UserList)IsExist(username string) bool{
	for _, user := range u.Users {
		if user.Name == username {
			return true
		}
	}
	return false
}

//验证密码
func (u *UserList)MatchPass(username, password string) bool{
	for _, user := range u.Users {
		if user.Name == username {
			return user.Password == password
		}
	}
	return false
}

//添加用户
func (u *UserList)AddUser(userinfo entity.UserData) {
	u.Users = append(u.Users, userinfo)
	u.writeToFile()
}

//删除用户
func (u *UserList)DeleteUser(username string) bool{
	for i, user := range u.Users {
		if user.Name == username {
			u.Users = append(u.Users[:i], u.Users[i+1:]...)
			u.writeToFile()
			return true
		}
	}
	return false
}
//将用户信息写入文件
func (u *UserList)writeToFile() {
	//UserList转json格式数据
	data, err := json.Marshal(*u)
	checkError(err)
	fwrite, err := os.OpenFile(userfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	checkError(err)

	//写入文件
	_, err = fwrite.Write(data)
	defer fwrite.Close()
	checkError(err)
}
//从json文件获取用户数据
func (u *UserList)readFromFile() {
	//判断文件是否存在
	_, err := os.Stat(userfile)
	if os.IsNotExist(err) {
		os.Mkdir(os.Getenv("GOPATH") + "/src/agenda/data", 0777)
		return 
	}

	fread, err := os.OpenFile(userfile, os.O_RDONLY, 0755)
	checkError(err)

	buffer := bufio.NewReader(fread)//存储已读取数据的缓冲区
	//读取文件
	//buf := make([]byte,
	bytes,err := buffer.ReadString('\n')
	//fmt.Println("userList: ",string(bytes))
	//os.Exit(1)
	//checkError(err)  //一直读到尾，所以肯定是eof
	if len(bytes) > 0{
		//解析json数据到UserList
		err = json.Unmarshal([]byte(bytes), u)
		checkError(err)
		//fmt.Println("userList: ",string(buf))
	}
	

	fread.Close()
}