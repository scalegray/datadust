package monit

import (
	//  "fmt"
	log "github.com/Sirupsen/logrus"
)

type Collector struct {
	//this contain all the metrics
	Cpu []*CpuStat
}
type Monit interface {
	SysExec(collector *Collector)
	//Send()
}

var Plugins = make(map[string]Monit)

//returns a monit interface
func RegisterPlugin(name string, plugin Monit) {
	if plugin == nil {
		log.Error("Cannot register an empty plugin")
	} else {
		Plugins[name] = plugin
		log.Info("Registered plugin: ", name)
	}

}

///func (col *Collector) Add() {
//  col.
//}
