package server

import (
	cfg "github.com/berttejeda/go389/cfg"
	model "github.com/berttejeda/go389/model"
	ldap "github.com/berttejeda/go389/server/ldap"
	"errors"
)

func GetServerHandler(name string, backendHandler model.BackendHandler) (model.ServerHandler, error) {
	switch cfg.GetServerType(name) {
	case "ldap":
		return ldap.NewLdapServerHandler(name, backendHandler)
	}
	return nil, errors.New("server not supported")
}
