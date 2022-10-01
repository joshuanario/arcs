package arcs

import (
	"sort"
	"time"
)

const LARC_DURATION = time.Microsecond * 50

type LArc struct {
	Locator  string
	Order    float64
	Children []*LArc
	lock     bool
}

func (node *LArc) IterQuery(loc string) (*LArc, int) {
	if node.lock {
		<-time.After(LARC_DURATION)
		return node.IterQuery(loc)
	}
	return node.unsafeIterQuery(loc)
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
	root.Children = make([]*LArc, 0, len(arcs))
	root.lock = true
	go func(rroot *LArc, aarcs []Arc, aarcrole string) {
		defer func() { root.lock = false }()
		sort.SliceStable(aarcs, func(i, j int) bool { return aarcs[i].Order < aarcs[j].Order })
		for _, arc := range aarcs {
			if arc.Arcrole == aarcrole {
				from, _ := rroot.unsafeIterQuery(arc.From)
				if from != nil {
					to, toIndex := rroot.unsafeIterQuery(arc.To)
					if to != nil {
						rroot.Children[toIndex] = rroot.Children[len(rroot.Children)-1]
						rroot.Children = rroot.Children[:len(rroot.Children)-1]
						from.Children = append(from.Children, to)
					} else {
						order := arc.Order
						from.Children = append(from.Children, &LArc{
							Locator:  arc.To,
							Order:    order,
							Children: make([]*LArc, 0, len(aarcs)),
						})
					}
				} else {
					from = &LArc{
						Locator:  arc.From,
						Children: make([]*LArc, 0, len(aarcs)),
					}
					rroot.Children = append(rroot.Children, from)
					to, toIndex := rroot.unsafeIterQuery(arc.To)
					if to != nil {
						rroot.Children[toIndex] = root.Children[len(rroot.Children)-1]
						rroot.Children = rroot.Children[:len(rroot.Children)-1]
						from.Children = append(from.Children, to)
					} else {
						order := arc.Order
						from.Children = append(from.Children, &LArc{
							Locator:  arc.To,
							Order:    order,
							Children: make([]*LArc, 0, len(aarcs)),
						})
					}
				}
			}
		}
	}(root, arcs, arcrole)
	return root
}

func (node *LArc) Paths(prior Path) []Path {
	if node == nil {
		return []Path{}
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
