package arcs

import (
	"sort"
	"time"
)

const LARC_DURATION = time.Nanosecond * 500

type LArc struct {
	Locator  string
	Order    float64
	Children []*LArc
	sorting  bool
	lock     bool
	arcs     []Arc
}

func (node *LArc) IterQuery(loc string) (*LArc, int) {
	if node.sorting {
		node.unsafeSort()
	}
	if node.lock {
		<-time.After(LARC_DURATION)
		return node.IterQuery(loc)
	}
	return node.unsafeIterQuery(loc)
}

func (root *LArc) unsafeSort() {
	defer func() {
		root.lock = false
		root.sorting = false
	}()
	arcs := root.arcs
	sort.SliceStable(arcs, func(i, j int) bool { return arcs[i].Order < arcs[j].Order })
	for _, arc := range arcs {
		from, _ := root.unsafeIterQuery(arc.From)
		if from != nil {
			to, toIndex := root.unsafeIterQuery(arc.To)
			if to != nil && toIndex >= 0 {
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
			to, toIndex := root.unsafeIterQuery(arc.To)
			if to != nil && toIndex >= 0 {
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
	root.arcs = arcs
}

func (node *LArc) unsafeIterQuery(loc string) (*LArc, int) {
	if node.Locator == loc {
		return node, -1
	}
	for i, c := range node.Children {
		ret, _ := c.unsafeIterQuery(loc)
		if ret != nil {
			return ret, i
		}
	}
	return nil, -1
}

func NewLArc(arcs []Arc, arcrole string) *LArc {
	root := &LArc{}
	root.Children = make([]*LArc, 0)
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

func (node *LArc) Paths(prior Path) []Path {
	if node == nil {
		return []Path{}
	}
	if node.sorting {
		node.unsafeSort()
	}
	if node.lock {
		<-time.After(LARC_DURATION)
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
