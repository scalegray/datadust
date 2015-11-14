package monitor

import (
	"bufio"
	"bytes"
	"regexp"
	"strings"
	"time"
)

const (
	STAT    = "stat"
	CPUINFO = "cpuinfo"
)

//CPU Stats
type CpuStat struct {
	Cpu       *Core
	Cores     []*Core
	Timestamp time.Time
}

type Core struct {
	Id    string
	User  string
	Nice  string
	Sys   string
	Irq   string
	Soft  string
	Steal string
	Guest string
	Gnice string
	Idle  string
}

func (c *CpuStat) SysExec(rec *Collector) {
	time.Sleep(1000 * time.Millisecond)
	//NumCores := getNumCores()
	statData, _ := ProcRead(STAT)

	scanner := bufio.NewScanner(bytes.NewReader([]byte(statData)))
	for scanner.Scan() {
		singleLine := scanner.Text()
		fields := strings.Fields(singleLine)
		cpuKey := fields[0]

		if strings.HasPrefix(cpuKey, "cpu") && CheckCPU(cpuKey) != nil {
			z := parseCore(fields)
			c.Cores = append(c.Cores, z)
		} else if strings.HasPrefix(cpuKey, "cpu") {
			all := parseCore(fields)
			c.Cpu = all
		}
	}
	rec.CpuStat = c
}
func CheckCPU(cpuKey string) []string {
	re := regexp.MustCompile("[0-9]+")
	no := re.FindAllString(cpuKey, -1)
	return no
}

func parseCore(f []string) *Core {

	core := &Core{ //terrible piece of design - get validated schema
		Id:    f[0],
		User:  f[1],
		Nice:  f[2],
		Sys:   f[3],
		Irq:   f[4],
		Soft:  f[5],
		Steal: f[6],
		Guest: f[7],
		Gnice: f[8],
		Idle:  f[9],
	}

	return core
}

//func (c *CpuStat) Send() {}

func init() {
	RegisterPlugin("cpu", &CpuStat{})
}

//TODO: temporary hack, move this out and return a full cpu_info struct
func getNumCores() string {

	mi, _ := ProcRead(CPUINFO)
	newScanner := bufio.NewScanner(bytes.NewReader([]byte(mi)))
	var NumCores string
	for newScanner.Scan() {
		line := newScanner.Text()
		cores := strings.Split(line, ":")

		cpuKey := cores[0]
		if strings.HasPrefix(cpuKey, "cpu cores") {
			NumCores = cores[1]

		}
	}
	return NumCores
}
