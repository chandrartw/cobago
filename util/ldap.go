package util

import (
	"os"

	"github.com/jtblin/go-ldap-client"
)

func NewLDAPClient() *ldap.LDAPClient {
	client := &ldap.LDAPClient{
		Base:         os.Getenv("LDAP_BASE"),
		Host:         os.Getenv("LDAP_HOST"),
		Port:         389,
		UseSSL:       false,
		BindDN:       os.Getenv("LDAP_BINDDN"),
		BindPassword: os.Getenv("LDAP_BINDPASSWORD"),
		UserFilter:   "(uid=%s)",
		GroupFilter:  "(memberUid=%s)",
		Attributes:   []string{"givenName", "sn", "mail", "uid"},
	}
	return client
}
