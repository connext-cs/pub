package ldap

import (
	"errors"

	"github.com/connext-cs/pub/config"
	"github.com/connext-cs/pub/logs"
	ldap "gopkg.in/ldap.v3"
	"fmt"
)

type User struct {
	Cn  string `json:"username"`
	Dn  string `json:"dn"`
	Uid string `json:"userkey"`
	Gid string `json:"gid"`
}

type Enum_SearchUserStatus uint8

const (
	UserOk        Enum_SearchUserStatus = 0
	UserUnknow    Enum_SearchUserStatus = 1
	UserPassError Enum_SearchUserStatus = 2
)

func LdapInit() (*ldap.Conn, error) {
	// s.Host + ":" + fmt.Sprintf("%d", s.Port)
	host := config.CLdapHost()
	port := config.CLdapPort()
	return ldap.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
}

func GetUserList() ([]User, error) {
	searchdn := config.CLdapSearchdn()
	return userList(searchdn)
}
func userList(searchdn string) ([]User, error) {
	user := new(User)
	users := make([]User, 0)
	l, err := LdapInit()
	defer l.Close()
	if err != nil {
		logs.Error("", err)
		return users, err
	}
	err = l.Bind(config.CLdapUser(), config.CLdapPassword())
	if err != nil {
		logs.Error("", err)
	}

	searchRequest := ldap.NewSearchRequest(
		searchdn, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(objectClass=organizationalPerson))",  // The filter to apply
		[]string{"uidNumber", "gidNumber", "cn"}, // A list attributes to retrieve
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		logs.Error("", err)
		return users, err
	}

	for _, v := range sr.Entries {
		user.Cn = v.GetAttributeValue("cn")
		user.Uid = v.GetAttributeValue("uidNumber")
		user.Gid = v.GetAttributeValue("gidNumber")
		user.Dn = v.DN
		users = append(users, *user)
	}

	return users, nil

}
func UserAuthentication(searchdn, user, password string) (Enum_SearchUserStatus, error) {
	logs.Info("", searchdn, user, password)
	l, err := LdapInit()
	defer l.Close()
	if err != nil {
		logs.Error("", err)
		return UserUnknow, err
	}
	searchRequest := ldap.NewSearchRequest(
		searchdn, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(objectClass=organizationalPerson)(uid="+user+"))", // The filter to apply
		[]string{"uidNumber", "gidNumber", "cn"},              // A list attributes to retrieve
		nil,
	)
	sr, err := l.Search(searchRequest)
	if err != nil {
		logs.Error("", err)
		return UserUnknow, err
	}

	if len(sr.Entries) != 1 {
		err = errors.New("User does not exist or too many entries returned")
		logs.Error("User does not exist or too many entries returned")
		return UserUnknow, err
	}
	userdn := sr.Entries[0].DN
	err = l.Bind(userdn, password)
	if err != nil {
		logs.Error("User Passwork  error", err)
		return UserPassError, err
	}

	return UserOk, nil
}
