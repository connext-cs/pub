package tokens

import (
	"encoding/json"
	"errors"
	"net/http"
	"github.com/connext-cs/pub/etcd"
	"github.com/connext-cs/pub/log"
	"strconv"
)

const (
	SecretKey = "welcome to use connextpaas!"
)

const USERIDHEADER = "userid"
const USERTYPEHEADER = "usertype"

type Enum_UserType uint8

const (
	User_None  Enum_UserType = 0 //无此用户
	User_Super Enum_UserType = 1 //超级用户
	User_Ldap  Enum_UserType = 2 //Ldap用戶
)

//token 在 etcd中存储的数据格式
type TokenValue struct {
	UserID   int64
	UserType Enum_UserType
	UserName string
}

// GetUserInfoFromRequest
// Input
// 		req: *http.Request
// OutPut
//      TokenValue:userid, usertype, username
//      error
func GetUserInfoFromRequest(req *http.Request) (*TokenValue, error) {
	tokenValue, err := ReadToken(req.Header.Get("Authorization"))
	if err != nil {
		log.Error("error:",err)
		return nil, err
	}
	return tokenValue, nil
}

func GetTokenValue(r *http.Request) (*TokenValue, error) {
	userid := r.Header.Get(USERIDHEADER)
	log.Info("userid:", userid)
	usertype := r.Header.Get(USERTYPEHEADER)
	log.Info("usertype:", usertype)

	idvalue, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		idvalue = 0
	}
	var usertypevalue Enum_UserType
	var usertypeInt int
	usertypeInt, err = strconv.Atoi(usertype)
	if err != nil {
		usertypevalue = User_None
	} else {
		if usertypeInt == int(User_Super) {
			usertypevalue = User_Super
		} else if usertypeInt == int(User_Ldap) {
			usertypevalue = User_Ldap
		} else {
			usertypevalue = User_None
		}
	}

	if idvalue <= 0 {
		if idvalue != -10000 {
			err = errors.New("userid is error.")
		}
	}
	if usertypevalue == User_None {
		err = errors.New("usertype is out of range.")
	}

	return &TokenValue{
		UserID:   idvalue,
		UserType: usertypevalue,
	}, err
}

//从etcd中获取tokenValue信息
func ReadToken(tokenkey string) (*TokenValue, error) {
	value, status, err := etcd.EtcdGet(etcd.ETCDROOT + "/" + tokenkey)
	if err != nil {
		return nil, err
	}
	if status == etcd.KeyNotInDb {
		err = errors.New("token not find in etcd.")
		return nil, err
	}
	var userbase TokenValue
	err = json.Unmarshal([]byte(value), &userbase)
	if err != nil {
		return nil, err
	}

	return &userbase, err
}

//往etcd写入token
func WriteToken(tokenkey string, userid int64, userType Enum_UserType, userName string) error {
	var userbase TokenValue
	userbase = TokenValue{
		UserID:   userid,
		UserType: userType,
		UserName: userName,
	}
	data, err := json.Marshal(userbase)
	if err != nil {
		log.Info("WriteToken, err:", err.Error())
		return err
	}
	err = etcd.EtcdPutLease(etcd.ETCDROOT+"/"+tokenkey, string(data), 999993600) //  默认token 为 多小时
	if err != nil {
		log.Info("WriteToken, EtcdPut err:", err.Error())
		return err
	}

	return err
}

// DeleteToken ... 从etcd删除token
func DeleteToken(tokenkey string) error {
	datastatus, err := etcd.EtcdDel(etcd.ETCDROOT + "/" + tokenkey)
	if err != nil {
		log.Info("DeleteToken, err:", err.Error())
		return err
	}
	log.Info("datastatus:", datastatus)
	return err
}
