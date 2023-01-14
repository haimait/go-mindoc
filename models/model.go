package models

import "gorm.io/gorm"

//type Model struct {
//	ID        uint           `gorm:"primarykey" json:"id"`
//	CreatedAt time.Time      `json:"created_at"`
//	UpdatedAt time.Time      `json:"updated_at"`
//	DeletedAt gort.DeletedAt `gorm:"index" json:"-"`
//}

type PageInfo struct {
	Page int `json:"page" validate:"min=1"` // 第几页
	Size int `json:"size" validate:"min=1"` // 每页结果数量  值为 9999时代码要所有数据
}

// SetPage ... 第几页
func (t *PageInfo) SetPage() {
	if t.Page <= 0 {
		t.Page = 1
	}
}

// SetSize ... 每页结果数量
func (t *PageInfo) SetSize(size int) {
	if t.Size <= 0 {
		t.Size = size
	}

}

// GetOffset 每页结果数量
func (t *PageInfo) GetOffset() int {
	if t.Page <= 0 {
		t.Page = 1
	}
	if t.Size <= 0 {
		t.Size = 10
	}
	offset := (t.Page - 1) * t.Size
	if offset < 0 {
		offset = 0
	}
	return offset
}

// Paginate ...
// page第几页 Size每页多少条面
func Paginate(page, size int) func(db *gorm.DB) *gorm.DB {
	var pageInfo PageInfo
	pageInfo.Page = page
	pageInfo.Size = size
	pageInfo.SetPage()
	pageInfo.SetSize(10)

	return func(db *gorm.DB) *gorm.DB {
		offset := (pageInfo.Page - 1) * pageInfo.Size
		if offset < 0 {
			offset = 0
		}
		return db.Offset(offset).Limit(pageInfo.Size)
	}
}

func FuncPage(page PageInfo) func(db *gorm.DB) *gorm.DB {
	return Paginate(page.Page, page.Size)
}

// GetPage 第几页
//func (t *PageInfo) GetPage() int {
//	if t.Page <= 0 {
//		t.Page = 1
//	}
//	return t.Page
//}

func ManualPage(page, size, current int, all ...bool) bool {
	if len(all) > 0 {
		return true
	}
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 1
	}
	return current > (page-1)*size && current <= page*size
}
