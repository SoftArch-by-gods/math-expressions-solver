package lab2

import (
	"bytes"
	. "gopkg.in/check.v1"
	"strings"
)

func (ms *MySuite) TestComputeInput(c *C) {
	var buf bytes.Buffer
	ch := &ComputeHandler{
		Input:  strings.NewReader("* - 13 -5 ^ 2 3"),
		Output: &buf,
	}

	compErr := ch.Compute()
	res, err := CalculatePrefix("* - 13 -5 ^ 2 3")

	c.Assert(buf.String(), Equals, res)
	c.Assert(compErr, IsNil)
	c.Assert(err, IsNil)
}

func (ms *MySuite) TestComputeExp1(c *C) {
	var buf bytes.Buffer
	ch := &ComputeHandler{
		Input:  strings.NewReader("* 8 5 2 3"),
		Output: &buf,
	}

	compErr := ch.Compute()
	c.Assert(compErr, ErrorMatches, "expression syntax error")
}

func (ms *MySuite) TestComputeExp2(c *C) {
	var buf bytes.Buffer
	ch := &ComputeHandler{
		Input:  strings.NewReader("+ 6 ^ 2 3 "),
		Output: &buf,
	}

	compErr := ch.Compute()
	c.Assert(compErr, ErrorMatches, "expression syntax error")
}

func (ms *MySuite) TestComputeExp3(c *C) {
	var buf bytes.Buffer
	ch := &ComputeHandler{
		Input:  strings.NewReader(""),
		Output: &buf,
	}

	compErr := ch.Compute()
	c.Assert(compErr, ErrorMatches, "expression syntax error")
}

func (ms *MySuite) TestComputeExp4(c *C) {
	var buf bytes.Buffer
	ch := &ComputeHandler{
		Input:  strings.NewReader("d / 2 3"),
		Output: &buf,
	}

	compErr := ch.Compute()
	c.Assert(compErr, ErrorMatches, "undefined symbol")
}
