/*
** Copyright (C) 2015 Yeshwanth Kumar <morpheyesh@gmail.com>
**
** Licensed under the Apache License, Version 2.0 (the "License");
** you may not use this file except in compliance with the License.
** You may obtain a copy of the License at
**
**     http://www.apache.org/licenses/LICENSE-2.0
**
** Unless required by applicable law or agreed to in writing, software
** distributed under the License is distributed on an "AS IS" BASIS,
** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
** See the License for the specific language governing permissions and
** limitations under the License.
**
 */

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/scalegray/datadust/monitor"
)

const CPU = "cpu"

func main() {
	for {
		Monitor()
		time.Sleep(100 * time.Millisecond)

	}
	http.ListenAndServe("localhost:9999", nil)
}

func Monitor() {
	fmt.Println("b000yah")
	collector := &monitor.Collector{}
	for _, p := range monitor.Plugins {
		p.SysExec(collector)
		//p.Send(monitChannel)
	}
	fmt.Println(collector.CpuStat.Cpu.User)
	fmt.Println(collector.MemStat.MemTotal)

}
