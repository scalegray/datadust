package monitor

import (
	//  "fmt"
	log "github.com/Sirupsen/logrus"
)

type Collector struct {
	//this contain all the metrics
	CpuStat *CpuStat
	MemStat *MemStat
}
type Monitor interface {
	SysExec(collector *Collector)
	//Send()
}

var Plugins = make(map[string]Monitor)

//returns a monitor interface
func RegisterPlugin(name string, plugin Monitor) {
	if plugin == nil {
		log.Error("Cannot register an empty plugin")
	} else {
		Plugins[name] = plugin
		log.Info("Registered plugin: ", name)
	}
}
