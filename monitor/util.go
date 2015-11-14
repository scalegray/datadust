package monitor

import (
	"io/ioutil"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func ProcRead(t string) (string, error) {
	p, e := ioutil.ReadFile("/proc/" + t)
	if e != nil {
		log.Error("Error in reading proc file ", t)
		return "", e
	}
	d := string(p)
	return d, nil
}

func parseInt64(s string) (i int64) {
	i, _ = strconv.ParseInt(s, 10, 64)
	return
}
