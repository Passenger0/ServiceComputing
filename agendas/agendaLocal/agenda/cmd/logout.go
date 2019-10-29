package cmd

import (
	"agenda/entity"
	"agenda/service"
	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "user logout the account",
	Long: `use "agenda logout" to logout!No need for flags' value`,

	Run: func(cmd *cobra.Command, args []string) {
	
		isLogin := entity.GetCurUser().GetLoginStatus()
		// 未登录则不可退出
		if !isLogin {
			service.Error.Println(entity.GetCurUser().UserName + "  logout failed: You did not login yet!")
			return
		}
	
		//改变status.json为登录状态并输出信息
		entity.GetCurUser().LogOut()
		service.Info.Println(entity.GetCurUser().UserName + "  logout succeed! ")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
