package pam

import (
	cfg "github.com/kernel164/go389/cfg"
	log "github.com/kernel164/go389/log"
	model "github.com/kernel164/go389/model"
)

type PamAuthHandler struct {
	model.AuthHandler
	settings PamAuthSettings
}

type PamAuthSettings struct {
	model.BaseAuth
	Service string
}

func NewPamAuthHandler(name string) (model.AuthHandler, error) {
	settings := PamAuthSettings{}
	log.Info("Auth Config name", name)
	cfg.GetAuthCfg(name, &settings)
	return PamAuthHandler{settings: settings}, nil
}

func (h PamAuthHandler) Auth(userName string, backendPasswd string, checkPasswd string) error {
	log.Info("Validating system password for", userName, "using service", h.settings.Service)
	return PAMAuth(h.settings.Service, userName, checkPasswd)
}
