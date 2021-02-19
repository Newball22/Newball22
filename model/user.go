package model

type UserInfo struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func NewUserInfo(id int, name, addr string) *UserInfo {
	return &UserInfo{
		ID:      id,
		Name:    name,
		Address: addr,
	}
}
