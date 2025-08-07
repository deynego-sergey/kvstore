package setup

import "sync"

type IApplicationConfig interface {
}

type applicationConfig struct{}

// @$(go env GOPATH)/bin/mockgen -destination=./mocks/mock_application.go -package=mocks github.com/josephbudd/cwt/src/internal/app Application
func (c *applicationConfig) init() error {

}

var o sync.Once
var config IApplicationConfig

func New() IApplicationConfig {
	o.Do(func() {
		_c := &applicationConfig{}
		if e := _c.init(); e != nil {
			panic(e)
		}
		config = _c
	})
	return config
}
