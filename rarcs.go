package arcs

import "sort"

type RArc struct {
	Locator  string
	Order    float64
	Children []*RArc
	hashmap  map[string]*RArc
}

func (node *RArc) HashQuery(loc string) *RArc {
	if node.Locator == loc {
		return node
	}
	return node.hashmap[loc]
}

func (node *RArc) getIndex(loc string) int {
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

func NewRArc(arcs []Arc, arcrole string) *RArc {
	root := &RArc{}
	root.hashmap = make(map[string]*RArc)
	root.Children = make([]*RArc, 0, len(arcs))
	sort.SliceStable(arcs, func(i, j int) bool { return arcs[i].Order < arcs[j].Order })
	for _, arc := range arcs {
		if arc.Arcrole == arcrole {
			from := root.HashQuery(arc.From)
			if from != nil {
				to := from.HashQuery(arc.To)
				toIndex := from.getIndex(arc.To)
				if to != nil && toIndex >= 0 {
					from.Children[toIndex] = from.Children[len(from.Children)-1]
					from.Children = from.Children[:len(from.Children)-1]
				} else {
					to = root.HashQuery(arc.To)
					toIndex := root.getIndex(arc.To)
					if to != nil && toIndex >= 0 {
						root.Children[toIndex] = root.Children[len(root.Children)-1]
						root.Children = root.Children[:len(root.Children)-1]
					} else {
						order := arc.Order
						to = &RArc{
							Locator:  arc.To,
							Order:    order,
							Children: make([]*RArc, 0, len(arcs)),
							hashmap:  make(map[string]*RArc),
						}
					}
				}
				from.Children = append(from.Children, to)
				from.hashmap[arc.To] = to
				root.hashmap[arc.To] = to
			} else {
				from = &RArc{
					Locator:  arc.From,
					Children: make([]*RArc, 0, len(arcs)),
					hashmap:  make(map[string]*RArc),
				}
				root.Children = append(root.Children, from)
				to := root.HashQuery(arc.To)
				toIndex := root.getIndex(arc.To)
				if to != nil && toIndex >= 0 {
					root.Children[toIndex] = root.Children[len(root.Children)-1]
					root.Children = root.Children[:len(root.Children)-1]
				} else {
					order := arc.Order
					to = &RArc{
						Locator:  arc.To,
						Order:    order,
						Children: make([]*RArc, 0, len(arcs)),
						hashmap:  make(map[string]*RArc),
					}
				}
				from.Children = append(from.Children, to)
				root.hashmap[arc.From] = from
				from.hashmap[arc.To] = to
				root.hashmap[arc.To] = to
			}
		}
	}
	return root
}

func (node *RArc) Paths(prior Path) []Path {
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
