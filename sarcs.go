package arcs

import (
	"sort"
	"sync"
)

type SArc struct {
	Locator  string
	Order    float64
	Children []*SArc
	lock     sync.RWMutex
	hashmap  map[string]*SArc
}

func (node *SArc) HashQuery(loc string) *SArc {
	node.lock.RLock()
	defer node.lock.RLock()
	if node.Locator == loc {
		return node
	}
	return node.hashmap[loc]
}

func NewSArc(arcs []Arc, arcrole string) *SArc {
	root := &SArc{}
	root.Children = make([]*SArc, 0, len(arcs))
	root.hashmap = make(map[string]*SArc)
	go func() {
		root.lock.Lock()
		defer root.lock.Unlock()
		sort.SliceStable(arcs, func(i, j int) bool { return arcs[i].Order < arcs[j].Order })
		for _, arc := range arcs {
			if arc.Arcrole == arcrole {
				from := root.HashQuery(arc.From)
				if from != nil {
					to := root.HashQuery(arc.To)
					if to != nil {
						continue
					} else {
						order := arc.Order
						to = &SArc{
							Locator:  arc.To,
							Order:    order,
							Children: make([]*SArc, 0, len(arcs)),
							hashmap:  make(map[string]*SArc),
						}
						from.Children = append(from.Children, to)
						from.hashmap[arc.To] = to
						root.hashmap[arc.To] = to
					}
				} else {
					from = &SArc{
						Locator:  arc.From,
						Children: make([]*SArc, 0, len(arcs)),
						hashmap:  make(map[string]*SArc),
					}
					root.Children = append(root.Children, from)
					root.hashmap[arc.From] = from
					to := root.HashQuery(arc.To)
					if to != nil {
						continue
					} else {
						order := arc.Order
						to = &SArc{
							Locator:  arc.To,
							Order:    order,
							Children: make([]*SArc, 0, len(arcs)),
						}
						from.Children = append(from.Children, to)
						from.hashmap[arc.To] = to
						root.hashmap[arc.To] = to
					}
				}
			}
		}
	}()
	return root
}

func (node *SArc) Paths(prior Path) []Path {
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
