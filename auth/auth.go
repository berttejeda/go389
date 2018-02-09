package auth

import (
	cfg "github.com/berttejeda/go389/cfg"
	log "github.com/berttejeda/go389/log"
	model "github.com/berttejeda/go389/model"
	pam "github.com/berttejeda/go389/auth/pam"
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

func GetAuthHandler(name string) (model.AuthHandler, error) {
	switch cfg.GetAuthType(name) {
	case "pam":
		log.Info("Authenticating using PAM")
		return pam.NewPamAuthHandler(name)
	}
	return NewDefaultAuthHandler(name)
}

type DefaultAuthHandler struct {
	model.AuthHandler
	settings DefaultAuthSettings
}

type DefaultAuthSettings struct {
	model.BaseAuth
	Hash string
}

func NewDefaultAuthHandler(name string) (model.AuthHandler, error) {
	settings := DefaultAuthSettings{}
	cfg.GetAuthCfg(name, &settings)
	return DefaultAuthHandler{settings: settings}, nil
}

func (h DefaultAuthHandler) Auth(userName string, backendPasswd string, checkPasswd string) error {
	switch h.settings.Hash {
	default:
		hash := sha256.New()
		hash.Write([]byte(checkPasswd))
		if backendPasswd != hex.EncodeToString(hash.Sum(nil)) {
			return errors.New("invalid password.")
		}
	}
	return nil
}
