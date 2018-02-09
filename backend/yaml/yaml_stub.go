// +build !yaml

package yaml

import (
	model "github.com/berttejeda/go389/model"
	"errors"
)

func NewYamlBackendHandler(name string) (model.BackendHandler, error) {
	return nil, errors.New("yaml backend not supported. Try build with yaml tag")
}
