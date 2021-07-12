package model

import "gorm.io/gorm"

type Pagetype struct {
	PageSize int `form:"page_size"`
	Page     int `form:"page"`
}

var defaultPageInfo = &Pagetype{
	PageSize: 10,
	Page:     1,
}

func Paginate(p *Pagetype) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if p.Page == 0 {
			p.Page = 1
		}

		switch {
		case p.PageSize > 100:
			p.PageSize = 100
		case p.PageSize <= 0:
			p.PageSize = 10
		}

		offset := (p.Page - 1) * p.PageSize
		return db.Offset(offset).Limit(p.PageSize)
	}
}
