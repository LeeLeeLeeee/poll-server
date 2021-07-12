package model

import (
	"errors"
	"log"
	"strconv"

	"github.com/fatih/structs"
)

type UserForm struct {
	ID          string `json:"id" form:"id"`
	UserId      string `json:"user_id" form:"user_id"`
	Email       string `json:"email" form:"email"`
	FirstName   string `json:"firstname" form:"firstname"`
	LastName    string `json:"lastname" form:"lastname"`
	CreatedDate string `json:"created_date" form:"created_date"`
	UpdatedDate string `json:"update_date" form:"update_date"`
	LastLogin   string `json:"last_login" form:"last_login"`
	LoginCount  string `json:"login_count" form:"login_count"`
	IsSuperuser string `json:"is_superuser" form:"is_superuser"`
}

type UserQuerySet struct {
}

func (UserQuerySet) getTableName() string {
	return "tb_user"
}

func (u UserQuerySet) CheckLogin(userId, password string) (id uint64, err error) {
	user := &UserForm{}

	result := Gdb.Table(u.getTableName()).Where("user_id = ? and password = ?", userId, password).First(&user)

	if result.Error != nil {
		return 0, result.Error
	}

	uid, err := strconv.ParseUint(user.ID, 10, 64)

	if err != nil {
		return 0, err
	}

	return uid, nil
}

func (u UserQuerySet) SelectOne(id interface{}) (res *UserForm, err error) {
	user := &UserForm{}

	result := Gdb.Table(u.getTableName()).First(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u UserQuerySet) Select(param interface{}) (res *[]UserForm, err error) {

	var p *Pagetype
	var pok, ucok bool
	var uc *UserForm

	user := &[]UserForm{}
	s := structs.New(param)

	PageInfo, pfok := s.FieldOk("PageInfo")
	UserFilter, ufok := s.FieldOk("UserFileter")

	if !pfok {
		p = defaultPageInfo
	} else {
		p, pok = PageInfo.Value().(*Pagetype)
		if !pok {
			p = defaultPageInfo
		}
	}

	if !ufok {
		uc = &UserForm{}
	} else {
		uc, ucok = UserFilter.Value().(*UserForm)
		if !ucok {
			uc = &UserForm{}
		}
	}

	result := Gdb.Scopes(Paginate(p)).Table(u.getTableName()).Where(uc).Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (u UserQuerySet) InsertOne(data *User) error {

	result := Gdb.Create(data)

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (u UserQuerySet) InsertMany(data *[]User) error {
	result := Gdb.CreateInBatches(data, len(*data))
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u UserQuerySet) DeleteOne(id interface{}) error {
	dbresult := Gdb.Delete(&User{}, id.(string))
	log.Println(dbresult.Statement.SQL.String())

	if dbresult.Error != nil {
		return dbresult.Error
	}

	if dbresult.RowsAffected == 0 {
		return errors.New("can't find the user")
	}

	return nil
}

func (u UserQuerySet) UpdateOne(id string, param *UserForm) error {

	dbresult := Gdb.Table(u.getTableName()).Where("id = ?", id).Updates(*param)

	if dbresult.Error != nil {
		return dbresult.Error
	}

	if dbresult.RowsAffected == 0 {
		return errors.New("can't find the user")
	}

	return nil
}
