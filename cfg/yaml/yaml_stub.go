// +build !yaml

package yaml

import (
	model "github.com/berttejeda/go389/model"
	"errors"
)

func NewYamlCfgHandler(name string, file string) (model.CfgHandler, error) {
	return nil, errors.New("yaml cfg not supported. Try build with yaml tag")
}
