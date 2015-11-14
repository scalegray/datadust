package monitor

import "testing"

func Test(t *testing.T) {
	check.TestingT(t)
}

type S struct{}

var _ = check.Suite(&S{})

func (s *S) TestprocRead(c *check.C) {

	_, error := procRead("stat")
	c.Assert(error, check.IsNil)
}
