// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cmd

import (
	"agenda/entity"
	"agenda/service"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Used to register an account",
	Long: `Use the register command in one of the forms below to register an account:
	1. agenda register -u username -p password -e email -t telephone 
	2. agenda register -uusername  -ppassword  -eemail  -ttelephone 

	Flags: 
		-u username
		-p password
		-e email
		-t telephone
`,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取参数值
		username, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("pass")
		email, _ := cmd.Flags().GetString("email")
		telephone, _ := cmd.Flags().GetString("tel")
	

		isLogin := entity.GetCurUser().GetLoginStatus()
		// 已经登陆不可注册
		if isLogin {
			service.Error.Println(username + "  register failed: You had already logined!")
			return
		}
		// 参数是否提供
		if username == "" || password == "" || email == "" || telephone == "" {
			service.Error.Println(username + "  register failed: values for four flags must be provided!")
			return
		}
		
		// 注册用户名不允许重复
		if service.AllUser.IsExist(username) {
			service.Error.Println(username + "  register failed: username has been used!")
			return
		}
		//添加用户
		userinfo := entity.UserData{
			Name : username,
			Password : password,
			Email : email,
			Tel : telephone,
		}
		service.AllUser.AddUser(userinfo)
		//输出信息
		service.Info.Println(username + "  register succeed!")

	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	registerCmd.Flags().StringP("user", "u", "", "username")
	registerCmd.Flags().StringP("pass", "p", "", "password")
	registerCmd.Flags().StringP("email", "e", "", "email address")
	registerCmd.Flags().StringP("tel", "t", "", "telephone number")
}
