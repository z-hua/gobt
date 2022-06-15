package decorator

import (
	"github.com/z-hua/gobt"
	"log"
)

type _logger struct {
	enterMsg string
	abackMsg string
}

func NewLogger(enterMsg, abackMsg string) gobt.IExecutor {
	return &_logger{enterMsg: enterMsg, abackMsg: abackMsg}
}

func (l *_logger) Name() string {
	return "logger"
}

func (l *_logger) Enter(rt *gobt.Runtime, node *gobt.Node) {
	log.Println(l.enterMsg)
	if len(node.Children) != 0 {
		rt.Ahead(node.Children[0], node)
	}
}

func (l *_logger) Aback(rt *gobt.Runtime, node *gobt.Node) {
	log.Println(l.abackMsg)
	rt.Leave(node)
}
