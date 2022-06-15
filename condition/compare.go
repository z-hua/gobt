package condition

import (
	"fmt"
	"github.com/z-hua/gobt"
)

/**
 * 比较节点
 * 比较变量的值
 * key : 变量名
 * cmp : 比较操作(>, >=, <, <=, ==, !=)
 * val : 操作数
 */
type _compare struct {
	key string
	cmp string
	val int
}

func NewCompare(key, cmp string, val int) gobt.IExecutor {
	return &_compare{
		key: key,
		cmp: cmp,
		val: val,
	}
}

func (c *_compare) Name() string {
	return "compare"
}

func (c *_compare) Enter(rt *gobt.Runtime, node *gobt.Node) {
	val := rt.Variant(c.key)

	met := false
	switch c.cmp {
	case ">":
		met = val > c.val
	case ">=":
		met = val >= c.val
	case "<":
		met = val < c.val
	case "<=":
		met = val <= c.val
	case "==":
		met = val == c.val
	case "!=":
		met = val != c.val
	default:
		panic(fmt.Sprintf("operand(%s) invalid", c.cmp))
	}
	if met {
		rt.SetSuccess()
	} else {
		rt.SetFailure()
	}
	rt.Leave(node)
}

func (c *_compare) Aback(rt *gobt.Runtime, node *gobt.Node) {
	panic("leaf node")
}
