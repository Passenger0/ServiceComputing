package cmd

import (
	"agenda/entity"
	"agenda/service"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "User login an account",
	Long: `Use the command in one of the forms below to login an account:
	1. agenda login -u username -p password
	2. agenda login -uusername  -ppassword

	Flags: 
	-u username
	-p password
`,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取参数值
		username, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("pass")
	
		isLogin := entity.GetCurUser().GetLoginStatus()
		// 一台电脑不能同时登录多个用户
		if isLogin {
			service.Error.Println(username + "  login failed: You had already logined!")
			return
		}
		
		// 参数是否提供
		if username == "" || password == ""{
			service.Error.Println(username + "  login failed: values for two flags must be provided!")
			return
		}
		
		// 登陆用户名必须存在
		if !service.AllUser.IsExist(username) {
			service.Error.Println(username + "  login failed: username does not exist!")
			return
		}
		// 用户名和密码必须匹配
		if !service.AllUser.MatchPass(username, password) {
			service.Error.Println(username + "  login failed: wrong password!")
			return
		}
		
		//更新curUser信息并输出操作成功信息
		email := service.AllUser.GetUserByName(username).GetEmail()
		tel := service.AllUser.GetUserByName(username).GetTel()

		entity.GetCurUser().LogIn(username, password,email,tel)
		service.Info.Println(username + "  login succeed!")

	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("user", "u", "", "username")
	loginCmd.Flags().StringP("pass", "p", "", "password")
}
