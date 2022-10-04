package arcs

import (
	"sort"
	"time"
)

const SARC_DURATION = time.Nanosecond * 500

type SArc struct {
	Locator  string
	Order    float64
	Children []*SArc
	sorting  bool
	lock     bool
	arcs     []Arc
	hashmap  map[string]*SArc
}

func (node *SArc) HashQuery(loc string) *SArc {
	if node.sorting {
		node.sorting = false
		go node.unsafeSort()
	}
	if node.lock {
		<-time.After(SARC_DURATION)
		return node.HashQuery(loc)
	}
	return node.unsafeHashQuery(loc)
}

func (root *SArc) unsafeSort() {
	defer func() {
		root.lock = false
	}()
	root.hashmap = make(map[string]*SArc)
	arcs := root.arcs
	sort.SliceStable(arcs, func(i, j int) bool { return arcs[i].Order < arcs[j].Order })
	for _, arc := range arcs {
		from := root.unsafeHashQuery(arc.From)
		if from != nil {
			to := from.unsafeHashQuery(arc.To)
			if to != nil {
				toIndex := from.getIndex(arc.To)
				from.Children[toIndex] = from.Children[len(from.Children)-1]
				from.Children = from.Children[:len(from.Children)-1]
			} else {
				to = root.unsafeHashQuery(arc.To)
				if to != nil {
					toIndex := root.getIndex(arc.To)
					root.Children[toIndex] = root.Children[len(root.Children)-1]
					root.Children = root.Children[:len(root.Children)-1]
				} else {
					order := arc.Order
					to = &SArc{
						Locator:  arc.To,
						Order:    order,
						Children: make([]*SArc, 0, len(arcs)),
						hashmap:  make(map[string]*SArc),
					}
				}
			}
			from.Children = append(from.Children, to)
			from.hashmap[arc.To] = to
			root.hashmap[arc.To] = to
		} else {
			from = &SArc{
				Locator:  arc.From,
				Children: make([]*SArc, 0, len(arcs)),
				hashmap:  make(map[string]*SArc),
			}
			root.Children = append(root.Children, from)
			to := root.unsafeHashQuery(arc.To)
			if to != nil {
				toIndex := root.getIndex(arc.To)
				root.Children[toIndex] = root.Children[len(root.Children)-1]
				root.Children = root.Children[:len(root.Children)-1]
			} else {
				order := arc.Order
				to = &SArc{
					Locator:  arc.To,
					Order:    order,
					Children: make([]*SArc, 0, len(arcs)),
					hashmap:  make(map[string]*SArc),
				}
			}
			from.Children = append(from.Children, to)
			root.hashmap[arc.From] = from
			from.hashmap[arc.To] = to
			root.hashmap[arc.To] = to
		}
	}
	root.arcs = make([]Arc, 0)
}

func (node *SArc) unsafeHashQuery(loc string) *SArc {
	if node.Locator == loc {
		return node
	}
	return node.hashmap[loc]
}

func (node *SArc) getIndex(loc string) int {
	if node.Locator == loc {
		return -1
	}
	for i, c := range node.Children {
		if c.Locator == loc {
			return i
		}
	}
	return -1
}

func NewSArc(arcs []Arc, arcrole string) *SArc {
	root := &SArc{}
	root.Children = make([]*SArc, 0)
	root.arcs = make([]Arc, 0, len(arcs))
	root.lock = true
	root.sorting = true
	for _, arc := range arcs {
		if arc.Arcrole != arcrole {
			continue
		}
		root.arcs = append(root.arcs, arc)
	}
	return root
}

func (node *SArc) Paths(prior Path) []Path {
	if node == nil {
		return []Path{}
	}
	if node.sorting {
		node.sorting = false
		go node.unsafeSort()
	}
	if node.lock {
		<-time.After(SARC_DURATION)
		return node.Paths(prior)
	}
	newPath := append(prior, node.Locator)
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
