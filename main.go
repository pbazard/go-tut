package main

import (
	"fmt"
)

func main() {
	fmt.Println(sayHi("Marco"))

	t := Triangle{
		Polygon{
			Sides: 3,
		},
	}
	fmt.Println(t.NSides()) // 3

	flow := NetFlowV5{
		NetFlow{
			Version: 5,
		},
	}
	fmt.Println(flow.NVersion()) // 5

	flow9 := NetFlowV9{
		NetFlow{
			Version: 9,
		},
	}
	fmt.Println(flow9.NVersion()) // 5
}

type NetFlowV5 struct {
	NetFlow // anonymous field
}

type NetFlowV9 struct {
	NetFlow // anonymous field
}

type NetFlow struct {
	Version   int
	FlowCount int
}

func (n *NetFlow) NVersion() int {
	return n.Version
}

func (n *NetFlow) NFlowCount() int {
	return n.FlowCount
}

type Record struct {
	Protocol string
}

type Triangle struct {
	Polygon // anonymous field
}

type Polygon struct {
	Sides int
}

func (p *Polygon) NSides() int {
	return p.Sides
}

func sayHi(person string) string {
	return fmt.Sprintf("Hi %s", person)
}
