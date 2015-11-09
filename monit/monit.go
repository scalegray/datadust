package monit


import (
//  "fmt"
  log "github.com/Sirupsen/logrus"
)

type Monit interface {

  SysExec() (*CpuStat)
  Send()
}

var Plugins = make(map[string]Monit)

//returns a monit interface
func RegisterPlugin(name string, plugin Monit) {
 if plugin == nil {
    log.Error("Cannot register an empty plugin")
  } else {
    Plugins[name] = plugin
   log.Info("Registered plugin --> ", name)
  }

}
