package arcs

import (
	"sort"
	"sync"
)

type LArc struct {
	Locator  string
	Order    float64
	Children []*LArc
	lock     sync.RWMutex
}

func (node *LArc) IterQuery(loc string) (*LArc, int) {
	node.lock.RLock()
	defer node.lock.RLock()
	if node.Locator == loc {
		return node, -1
	}
	for i, c := range node.Children {
		ret, _ := c.IterQuery(loc)
		if ret != nil {
			return ret, i
		}
	}
	return nil, -1
}

func NewLArc(arcs []Arc, arcrole string) *LArc {
	var root *LArc
	root.Children = make([]*LArc, 0, len(arcs))
	go func() {
		root.lock.Lock()
		defer root.lock.Unlock()
		sort.SliceStable(arcs, func(i, j int) bool { return arcs[i].Order < arcs[j].Order })
		for _, arc := range arcs {
			if arc.Arcrole == arcrole {
				from, _ := root.IterQuery(arc.From)
				if from != nil {
					to, toIndex := root.IterQuery(arc.To)
					if to != nil {
						root.Children[toIndex] = root.Children[len(root.Children)-1]
						root.Children = root.Children[:len(root.Children)-1]
						from.Children = append(from.Children, to)
					} else {
						order := arc.Order
						from.Children = append(from.Children, &LArc{
							Locator:  arc.To,
							Order:    order,
							Children: make([]*LArc, 0, len(arcs)),
						})
					}
				} else {
					from = &LArc{
						Locator:  arc.From,
						Children: make([]*LArc, 0, len(arcs)),
					}
					root.Children = append(root.Children, from)
					to, toIndex := root.IterQuery(arc.To)
					if to != nil {
						root.Children[toIndex] = root.Children[len(root.Children)-1]
						root.Children = root.Children[:len(root.Children)-1]
						from.Children = append(from.Children, to)
					} else {
						order := arc.Order
						from.Children = append(from.Children, &LArc{
							Locator:  arc.To,
							Order:    order,
							Children: make([]*LArc, 0, len(arcs)),
						})
					}
				}
			}
		}
	}()
	return root
}

func (node *LArc) Paths(prior Path) []Path {
	if node == nil {
		return []Path{}
	}
	newPath := append(prior, node.Locator)
	node.lock.RLock()
	defer node.lock.RLock()
	if len(node.Children) <= 0 {
		return []Path{
			newPath,
		}
	}
	var ret []Path
	for _, child := range node.Children {
		ret = append(ret, child.Paths(newPath)...)
	}
	return ret
}
