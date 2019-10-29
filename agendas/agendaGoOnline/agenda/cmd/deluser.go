package cmd

import (
	"git.go-online.org.cn/passenger0/agenda/entity"
	"git.go-online.org.cn/passenger0/agenda/service"
	"github.com/spf13/cobra"
)

// deluserCmd represents the deluser command
var deluserCmd = &cobra.Command{
	Use:   "deluser",
	Short: "delete the current user from the system",
	Long: `use the command in the form below to delete the current user from the system 
	agenda deluser -p password
	
	Flags: 
	-p password
`,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取参数值
		password, _ := cmd.Flags().GetString("pass")
		
		isLogin := entity.GetCurUser().GetLoginStatus()
		//未登录不可注销
		if !isLogin {
			service.Error.Println(entity.GetCurUser().UserName + "  deluser failed: You did not login yet!")
			return
		}

		// 密码不能为空
		if	password == "" {
			service.Error.Println(entity.GetCurUser().UserName + "  deluser failed: password must be provided")
			return
		}

		// 当前用户名与密码参数是否匹配
		username := entity.GetCurUser().UserName
		if !service.AllUser.MatchPass(username, password) {
			service.Error.Println(entity.GetCurUser().UserName + "  deluser failed: wrong password!")
			return
		}

		// 删除当前用户
		service.AllUser.DeleteUser(username)
		
		// 登出系统
		entity.GetCurUser().LogOut()
		
		// 输出信息
		service.Info.Println(entity.GetCurUser().UserName + "  deluser succeed!")
	},
}

func init() {
	rootCmd.AddCommand(deluserCmd)

	deluserCmd.Flags().StringP("pass", "p", "", "password")
}
