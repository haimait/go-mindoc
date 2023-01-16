package models

import "errors"

const (
	UserAuthTypeSystem  string = "system" //平台内部
	UserAuthTypeSmallWX string = "wxMini" //微信小程序
)

type SysUser struct {
	UserId   int    `gorm:"primaryKey;autoIncrement;comment:编码"  json:"userId"`
	Username string `json:"username" gorm:"size:64;comment:用户名"`
	Password string `json:"-" gorm:"size:128;comment:密码"`
	NickName string `json:"nickName" gorm:"size:128;comment:昵称"`
	Phone    string `json:"phone" gorm:"size:11;comment:手机号"`
	RoleId   int    `json:"roleId" gorm:"size:20;comment:角色ID"`
	Salt     string `json:"-" gorm:"size:255;comment:加盐"`
	Avatar   string `json:"avatar" gorm:"size:255;comment:头像"`
	Sex      string `json:"sex" gorm:"size:255;comment:性别"`
	Email    string `json:"email" gorm:"size:128;comment:邮箱"`
	DeptId   int    `json:"deptId" gorm:"size:20;comment:部门"`
	PostId   int    `json:"postId" gorm:"size:20;comment:岗位"`
	Remark   string `json:"remark" gorm:"size:255;comment:备注"`
	Status   string `json:"status" gorm:"size:4;comment:状态"`
	DeptIds  []int  `json:"deptIds" gorm:"-"`
	PostIds  []int  `json:"postIds" gorm:"-"`
	RoleIds  []int  `json:"roleIds" gorm:"-"`
	//Dept     *SysDept `json:"dept"`
	ControlBy
	ModelTime
}

func (SysUser) TableName() string {
	return "sys_user"
}

func (t *SysUser) Generate() ActiveRecord {
	o := *t
	return &o
}

func (t *SysUser) GetId() interface{} {
	return t.UserId
}

func (t *SysUser) CheckUser(name string) (b bool) {
	var user SysUser
	DB.Table(t.TableName()).Select("user_id").Where("user_name = ?", name).First(&user)
	if user.UserId > 0 {
		return true
	}
	return false
}

// GetUserDetailWithLogin 用户登陆时查询用户信息
func (t *SysUser) GetUserDetailWithLogin(loginName string) (user SysUser, err error) {
	if len(loginName) == 0 {
		err = errors.New("authKey is not empty")
		return
	}
	table := DB.Table(t.TableName())
	table.Where("username LIKE ?", loginName+"%")
	table.Or("phone = ?", loginName)
	if err := table.First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// GetUserDetail 查询用户
func (t *SysUser) GetUserDetail() (user SysUser, err error) {
	table := DB.Debug().Table(t.TableName())
	if t.UserId != 0 {
		table.Where("user_id = ?", t.UserId)
	}
	if t.Username != "" {
		table.Where("username LIKE ?", t.Username+"%")
	}
	if t.Phone != "" {
		table.Where("phone = ?", t.Phone)
	}
	if err := table.First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// GetUserList 查询用户列表
func (t *SysUser) GetUserList(page PageInfo) (users []SysUser, total int64, err error) {
	table := DB.Table(t.TableName()).
		//Select("user_id,username,nick_name,photo,description,created_at").
		Scopes(
			//MakeCondition(),
			Paginate(page),
		).
		Order("user_id desc")

	// if Size < 0  return all data
	if page.Size < 0 {
		table.Limit(-1).Offset(-1)
	}
	if t.UserId > 0 {
		table = table.Where("user_id = ?", t.UserId)
	}
	if t.Username != "" {
		table = table.Where("username LIKE ?", t.Username+"%")
	}
	err = table.Find(&users).
		Limit(-1).Offset(-1).Count(&total).Error

	if err != nil {
		return users, total, err
	}
	return users, total, nil
}

// DeleteUser 删除用户
func (t *SysUser) DeleteUser(uid int) (err error) {
	if uid == 0 {
		err = errors.New("authKey is not empty")
		return
	}
	var user SysUser
	err = DB.Table(t.TableName()).Where("user_id = ? ", uid).Delete(&user).Error
	return err
}

//func NewUserModel(conn sqlx.SqlConn) SysUser {
//	return &SysUser{
//		CachedConn: sqlc.NewConn(conn, c),
//		table:      "`user`",
//	}
//}

//type (
//	UserModel interface {
//		GetUserDetail() (user SysUser, err error)
//	}
//defaultUserModel struct {
//	sqlc.CachedConn
//	table string
//}
//)
