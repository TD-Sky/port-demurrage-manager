package models

type User struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}

type UserInfo struct {
	Name  string   `json:"username"`
	Roles []string `json:"roles"`
}
