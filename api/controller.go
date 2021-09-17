package api

import (
	"fmt"

	"github.com/anuvu/zot/pkg/log"
)

type Controller struct {
	Config  *Config
	Address string
	Log     log.Logger
}

func NewController(config *Config) *Controller {
	var controller Controller

	logger := log.NewLogger(config.Log.Level, config.Log.Output)

	controller.Config = config
	controller.Address = fmt.Sprintf("http://%s:%s", config.HTTP.Host, config.HTTP.Port)
	fmt.Println(controller.Address)
	
	controller.Log = logger

	return &controller
}

func (c *Controller) Run() {
	RunZotExporter(c)
}
