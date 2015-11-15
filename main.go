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
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/scalegray/datadust/monitor"
)

const CPU = "cpu"

func main() {
	collector := &monitor.Collector{}

	for {
		getMetrics(collector)
		time.Sleep(1000 * time.Millisecond)

	}
	http.ListenAndServe("localhost:9999", nil)
}

func getMetrics(c *monitor.Collector) {
	for _, p := range monitor.Plugins {
		p.SysExec(c)
		//p.Send(monitChannel)
	}
	postData(c)
}

func postData(c *monitor.Collector) {
	fmt.Println("conveverting data")
	data, _ := json.Marshal(c)
	fmt.Println("==================")
	fmt.Println(string(data))
}
