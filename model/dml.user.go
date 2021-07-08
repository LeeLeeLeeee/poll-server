package model

import "log"

type UserForm struct {
	ID          string `json:"id"`
	UserId      string `json:"user_id" binding:"required"`
	Passwrod    string `json:"password" binding:"required"`
	Email       string `json:"email" binding:"required"`
	FirstName   string `json:"firstname" binding:"required"`
	LastName    string `json:"lastname" binding:"required"`
	IsSuperuser bool   `json:"is_superuser" binding:"required"`
}

type UserQuerySet struct {
}

func (UserQuerySet) getTableName() string {
	return "tb_user"
}

func (u UserQuerySet) SelectOne(id interface{}) {
	user := &UserForm{}

	Gdb.Table(u.getTableName()).First(&user, id)
	log.Println(user)

}

// func (u *UserQuerySet) InsertOne(data interface{}) {
// 	db := u.db.DB
// 	db.Create(data.(UserForm))
// }

// func (u *UserQuerySet) InsertMany(data interface{}) {
// 	db := u.db.DB
// 	data.(UserForm)
// 	db.CreateBatchSize(data, length)
// }
