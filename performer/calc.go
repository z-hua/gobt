package performer

import (
	"fmt"
	"github.com/z-hua/gobt"
)

type _calc struct {
	key string
	op  string
	val int
}

func NewCalc(key, op string, val int) gobt.IExecutor {
	return &_calc{key: key, op: op, val: val}
}

func (c *_calc) Name() string {
	return "calc"
}

func (c *_calc) Enter(rt *gobt.Runtime, node *gobt.Node) {
	val := rt.Variant(c.key)
	switch c.op {
	case "+":
		val += c.val
	case "-":
		val -= c.val
	case "*":
		val *= c.val
	case "/":
		val /= c.val
	case "%":
		val %= c.val
	default:
		panic(fmt.Sprintf("unsupport openand(%s)", c.op))
	}
	rt.SetVariant(c.key, val)
	rt.SetSuccess()
	rt.Leave(node)
}

func (c *_calc) Aback(rt *gobt.Runtime, node *gobt.Node) {
	panic("leaf node")
}
