package cmd

import (
	"fmt"
	"agenda/entity"
	"agenda/service"
	"github.com/spf13/cobra"
)

// qryuserCmd represents the qryuser command
var qryuserCmd = &cobra.Command{
	Use:   "qryuser",
	Short: "query all the users",
	Long: `use "agenda qryuser" to query all the users!No need for flags' value`,

	Run: func(cmd *cobra.Command, args []string) {
		
		isLogin := entity.GetCurUser().GetLoginStatus()
		// 仅登录用户可查询其他用户信息
		if !isLogin {
			service.Error.Println(entity.GetCurUser().UserName + "  qryuser failed: You did not login yet!")
			return
		}

		//service.Info.Println(entity.GetCurUser().UserName + "  qryuser succeed!")
		user_list := service.AllUser.GetAllUsers()
		fmt.Println("There are ", len(user_list), " users：")
		fmt.Println("Name--Email--Telephone")
		for _, user := range user_list {
			fmt.Println(user.Name, " ", user.Email, " ", user.Tel)
		}
		// 输出信息
		service.Info.Println(entity.GetCurUser().UserName + "  qryuser succeed!")
	},
}

func init() {
	rootCmd.AddCommand(qryuserCmd)
}
