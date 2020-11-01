package datamodels

type User struct {
	ID           int64  `json:"ID" sql:"id" form:"ID"`
	Nickname     string `json:"nickName" sql:"nickName" form:"nickName"`
	Username     string `json:"userName" sql:"userName" form:"userName"`
	HashPassword string `json:"-" sql:"password" form:"password"`
}
