package datamodels

type User struct {
	ID           int64  `json:"ID" sql:"id" form:"ID"`
	NickName     string `json:"nickName" sql:"nickName" form:"nickName"`
	UserName     string `json:"userName" sql:"userName" form:"userName"`
	HashPassword string `json:"-" sql:"password" form:"password"`
}
