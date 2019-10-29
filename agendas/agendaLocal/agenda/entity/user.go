package entity

//用户信息数据结构
type UserData struct{
	Name string      `json:"name"`
	Password string  `json:"password"`
	Email string     `json:"email"`
	Tel string 		 `json:"telephone"`
}
func (u *UserData) GetName() string {
	return u.Name
}
func (u *UserData) GetPass() string {
	return u.Password
}
func (u *UserData) GetEmail() string {
	return u.Email
}
func (u *UserData) GetTel() string {
	return u.Tel
}

func (u *UserData) SetName(name string) {
	u.Name = name
}
func (u *UserData) SetPass(pass string) {
	u.Password = pass
}
func (u *UserData) SetEmail(email string) {
	u.Email = email
}
func (u *UserData) SetTel(tel string) {
	u.Tel = tel
}
