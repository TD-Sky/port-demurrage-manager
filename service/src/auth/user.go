package auth

type User struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}

func (self *User) Correct() bool {
	return self.Name == "admin" && self.Password == "12345678"
}

type UserInfo struct {
	Name  string   `json:"username"`
	Roles []string `json:"roles"`
}
