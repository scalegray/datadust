package monitor

/*
import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
)

const (
	MEMORY_INFO = "meminfo"
)

type MemStat struct {
	Memory    []*Memory
	Timestamp time.Time
}

type Memory struct {
	MemTotal          int64 `json:"mem_total,omitempty"`
	MemFree           int64 `json:"mem_free,omitempty"`
	Buffers           int64 `json:"buffers,omitempty"`
	Cached            int64 `json:"cached,omitempty"`
	SwapCached        int64 `json:"swap_cached,omitempty"`
	Active            int64 `json:"active,omitempty"`
	Inactive          int64 `json:"inactive,omitempty"`
	ActiveAnon        int64 `json:"active_anon,omitempty"`
	InactiveAnon      int64 `json:"inactive_anon,omitempty"`
	ActiveFile        int64 `json:"active_file,omitempty"`
	InactiveFile      int64 `json:"inactive_file,omitempty"`
	Unevictable       int64 `json:"unevictable,omitempty"`
	Mlocked           int64 `json:"mlocked,omitempty"`
	SwapTotal         int64 `json:"swap_total,omitempty"`
	SwapFree          int64 `json:"swap_free,omitempty"`
	Dirty             int64 `json:"dirty,omitempty"`
	Writeback         int64 `json:"writeback,omitempty"`
	AnonPages         int64 `json:"anon_pages,omitempty"`
	Mapped            int64 `json:"mapped,omitempty"`
	Shmem             int64 `json:"shmem,omitempty"`
	Slab              int64 `json:"slab,omitempty"`
	SReclaimable      int64 `json:"s_reclaimable,omitempty"`
	SUnreclaim        int64 `json:"s_unreclaim,omitempty"`
	KernelStack       int64 `json:"kernel_stack,omitempty"`
	PageTables        int64 `json:"page_tables,omitempty"`
	NFS_Unstable      int64 `json:"nfs_unstable,omitempty"`
	Bounce            int64 `json:"bounce,omitempty"`
	WritebackTmp      int64 `json:"writeback_tmp,omitempty"`
	CommitLimit       int64 `json:"commit_limit,omitempty"`
	Committed_AS      int64 `json:"committed_as,omitempty"`
	VmallocTotal      int64 `json:"vmalloc_total,omitempty"`
	VmallocUsed       int64 `json:"vmalloc_used,omitempty"`
	VmallocChunk      int64 `json:"vmalloc_chunk,omitempty"`
	HardwareCorrupted int64 `json:"hardware_corrupted,omitempty"`
	AnonHugePages     int64 `json:"anon_huge_pages,omitempty"`
	HugePages_Total   int64 `json:"huge_pages_total,omitempty"`
	HugePages_Free    int64 `json:"huge_pages_free,omitempty"`
	HugePages_Rsvd    int64 `json:"huge_pages_rsvd,omitempty"`
	HugePages_Surp    int64 `json:"huge_pages_surp,omitempty"`
	Hugepagesize      int64 `json:"hugepagesize,omitempty"`
	DirectMap4k       int64 `json:"direct_map4k,omitempty"`
	DirectMap2M       int64 `json:"direct_map2_m,omitempty"`
	DirectMap1G       int64 `json:"direct_map1_g,omitempty"`
}

func (c *MemStat) SysExec(rec *Collector) {
	//time.Sleep(1000 * time.Millisecond)
	//NumCores := getNumCores()
	statData, _ := procRead(MEMORY_INFO)

	newScanner := bufio.NewScanner(bytes.NewReader([]byte(statData)))
	for newScanner.Scan() {
		fmt.Println(newScanner.Text())
				singleLine := newScanner.Text()
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
	rec.CpuStat = append(rec.CpuStat, c)
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

//TODO: Move this out to  common
func procRead(t string) (string, error) {
	p, e := ioutil.ReadFile("/proc/" + t)
	if e != nil {
		log.Error("Error in reading proc file ", t)
		return "", e
	}
	d := string(p)
	return d, nil
}

//func (c *CpuStat) Send() {}

func init() {
	RegisterPlugin("cpu", &CpuStat{})
}

//TODO: temporary hack, move this out and return a full cpu_info struct
func getNumCores() string {

	mi, _ := procRead(CPUINFO)
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
*/
