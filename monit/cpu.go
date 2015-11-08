package monit

import (

//  "fmt"
)

type CPU struct {}

func (c *CPU) Collect() {}

func (c *CPU) Send() {}

func init() {
 RegisterPlugin("cpu", &CPU{})

}
