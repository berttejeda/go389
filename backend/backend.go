package backend

import (
	cfg "github.com/berttejeda/go389/cfg"
	model "github.com/berttejeda/go389/model"
	proxy "github.com/berttejeda/go389/backend/proxy"
	yaml "github.com/berttejeda/go389/backend/yaml"
	"errors"
)

func GetBackendHandler(name string) (model.BackendHandler, error) {
	switch cfg.GetBackendType(name) {
	case "yaml":
		return yaml.NewYamlBackendHandler(name)
	case "proxy":
		return proxy.NewProxyBackendHandler(name)
	}
	return nil, errors.New("backend not supported")
}
