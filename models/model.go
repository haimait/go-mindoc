package models

import (
	"github.com/go-admin-team/go-admin-core/tools/search"
	"gorm.io/gorm"
)

//type Model struct {
//	ID        uint           `gorm:"primarykey" json:"id"`
//	CreatedAt time.Time      `json:"created_at"`
//	UpdatedAt time.Time      `json:"updated_at"`
//	DeletedAt gort.DeletedAt `gorm:"index" json:"-"`
//}

type PageInfo struct {
	Page   int `json:"page" validate:"min=1"`   // 第几页
	Size   int `json:"size" validate:"min=1"`   // 每页结果数量  值为 -1 时代码要所有数据
	Offset int `json:"offset" validate:"min=1"` // 跳过多少条
}

/*
	SetPagination ....
	page 第几页
	size 每页结果数量
    sizeDefault 默认值
	offset 跳过多少条
*/
func (t *PageInfo) SetPagination(pi PageInfo, sizeDefault int) {
	if pi.Page <= 0 {
		t.Page = 1
	}
	if pi.Size <= 0 {
		t.Size = sizeDefault
	}
	offset := (t.Page - 1) * t.Size
	if offset < 0 {
		offset = 0
	}
	t.Offset = offset
}

// Paginate ...
// page第几页 Size每页多少条面
func Paginate(page PageInfo) func(db *gorm.DB) *gorm.DB {
	var pageInfo PageInfo
	pageInfo.SetPagination(page, 10)
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pageInfo.Offset).Limit(pageInfo.Size)
	}
}

func MakeCondition(q interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		condition := &search.GormCondition{
			GormPublic: search.GormPublic{},
			Join:       make([]*search.GormJoin, 0),
		}
		search.ResolveSearchQuery("mysql", q, condition)
		for _, join := range condition.Join {
			if join == nil {
				continue
			}
			db = db.Joins(join.JoinOn)
			for k, v := range join.Where {
				db = db.Where(k, v...)
			}
			for k, v := range join.Or {
				db = db.Or(k, v...)
			}
			for _, o := range join.Order {
				db = db.Order(o)
			}
		}
		for k, v := range condition.Where {
			db = db.Where(k, v...)
		}
		for k, v := range condition.Or {
			db = db.Or(k, v...)
		}
		for _, o := range condition.Order {
			db = db.Order(o)
		}
		return db
	}
}
