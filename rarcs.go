package arcs

import "sort"

type RArcs struct {
	Locator  string
	Order    float64
	Children []*RArcs
	hashmap  map[string]*RArcs
}

func (node *RArcs) HashQuery(loc string) *RArcs {
	if node.Locator == loc {
		return node
	}
	return node.hashmap[loc]
}

func NewRArcs(arcs []Arc, arcrole string) *RArcs {
	root := &RArcs{}
	root.Children = make([]*RArcs, 0, len(arcs))
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
					to = &RArcs{
						Locator:  arc.To,
						Order:    order,
						Children: make([]*RArcs, 0, len(arcs)),
						hashmap:  make(map[string]*RArcs),
					}
					from.Children = append(from.Children, to)
					from.hashmap[arc.To] = to
					root.hashmap[arc.To] = to
				}
			} else {
				from = &RArcs{
					Locator:  arc.From,
					Children: make([]*RArcs, 0, len(arcs)),
					hashmap:  make(map[string]*RArcs),
				}
				root.Children = append(root.Children, from)
				root.hashmap[arc.From] = from
				to := root.HashQuery(arc.To)
				if to != nil {
					continue
				} else {
					order := arc.Order
					to = &RArcs{
						Locator:  arc.To,
						Order:    order,
						Children: make([]*RArcs, 0, len(arcs)),
					}
					from.Children = append(from.Children, to)
					from.hashmap[arc.To] = to
					root.hashmap[arc.To] = to
				}
			}
		}
	}
	return root
}

func (node *RArcs) Paths(prior Path) []Path {
	if node == nil {
		return []Path{}
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
