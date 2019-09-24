package types

import (
	"github.com/connext-cs/pub/jsontime"
)

const (
	UserTypeSuperAdmin = iota
	UserTypeCustom
	UserTypeLdap
)

type AuthUsers struct {
	Id          string            `json:"user_id"`
	Type        int               `json:"user_type" xorm:"-"`
	UserName    string            `json:"user_name"`
	NickName    string            `json:"nick_name"`
	Phone       string            `json:"phone"`
	Position    string            `json:"position" xorm:"-"`
	Status      int               `json:"status"`
	Deleted     int               `json:"-"`
	CreatedBy   string            `json:"-"`
	CreatedTime jsontime.JsonTime `json:"create_time" xorm:"created"`
	UpdatedBy   string            `json:"-"`
	UpdatedTime jsontime.JsonTime `json:"-" xorm:"updated"`
}
