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
  "github.com/scalegray/datadust/monit"

)

const CPU = "cpu"

func main() {

//1. get from config
var monitElements [4]string
monitElements[0] = CPU


for _, x := range monitElements {
    //2. execute each plugin concurrently
    //3.store it in influxDB
}


}
