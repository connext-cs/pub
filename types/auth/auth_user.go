package auth

import (
	"github.com/connext-cs/pub/jsontime"
)

//用户类型
const (
	//超级管理员
	UserTypeSuperAdmin = iota
	//普通用户
	UserTypeCustom
	//LDAP用户
	UserTypeLdap
)

type AuthUsers struct {
	Id               string            `json:"user_id"`
	Type             int               `json:"user_type" xorm:"-"`
	UserName         string            `json:"user_name"`
	NickName         string            `json:"nick_name"`
	Phone            string            `json:"phone"`
	Position         string            `json:"position" xorm:"-"`
	Status           int               `json:"status"`
	Email            string            `json:"email"`
	OrganizationId   string            `json:"-"`
	OrganizationName string            `json:"-"`
	Ldapdn           string            `json:"-"`
	Deleted          int               `json:"-"`
	CreatedBy        string            `json:"-"`
	CreatedTime      jsontime.JsonTime `json:"create_time" xorm:"created"`
	UpdatedBy        string            `json:"-"`
	UpdatedTime      jsontime.JsonTime `json:"-" xorm:"updated"`
}

