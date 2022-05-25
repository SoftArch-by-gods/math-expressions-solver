package lab2

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (ms *MySuite) TestCalculatePrefix1(c *C) {
	res, err := CalculatePrefix("+ 3 2")
	c.Assert(res, Equals, "5")
	c.Assert(err, IsNil)
}

func (ms *MySuite) TestCalculatePrefix2(c *C) {
	res, err := CalculatePrefix("* - 13 -5 ^ 2 3")
	c.Assert(res, Equals, "144")
	c.Assert(err, IsNil)
}

func (ms *MySuite) TestCalculatePrefix3(c *C) {
	res, err := CalculatePrefix("/ * + 0.5 1.5 7 8")
	c.Assert(res, Equals, "1.75")
	c.Assert(err, IsNil)
}

func (ms *MySuite) TestCalculatePrefix4(c *C) {
	res, err := CalculatePrefix("+ 2 / 31.35 * 5 + - 5 3 ^ 5 4")
	c.Assert(res, Equals, "2.01")
	c.Assert(err, IsNil)
}

func (ms *MySuite) TestCalculatePrefix5(c *C) {
	res, err := CalculatePrefix("- 8.5 + + 150 4 - * 8 0.75 + ^ 9 2 70.25")
	c.Assert(res, Equals, "-0.25")
	c.Assert(err, IsNil)
}

func (ms *MySuite) TestCalculatePrefix6(c *C) {
	_, err := CalculatePrefix("")
	c.Assert(err, ErrorMatches, "expression syntax error")
}

func (ms *MySuite) TestCalculatePrefix7(c *C) {
	_, err := CalculatePrefix("h 9 6")
	c.Assert(err, ErrorMatches, "undefined symbol")
}

func (ms *MySuite) TestCalculatePrefix8(c *C) {
	_, err := CalculatePrefix("+ 10 9 6")
	c.Assert(err, ErrorMatches, "expression syntax error")
}

func (ms *MySuite) TestCalculatePrefix9(c *C) {
	_, err := CalculatePrefix("/ 5 1 ")
	c.Assert(err, ErrorMatches, "expression syntax error")
}

func (ms *MySuite) TestCalculatePrefix10(c *C) {
	_, err := CalculatePrefix("la la la")
	c.Assert(err, ErrorMatches, "expression syntax error")
}

func ExampleCalculatePrefix() {
	res, _ := CalculatePrefix("* + 1 3 - 9 4")
	fmt.Println(res)

	// Output:
	// 20
}
