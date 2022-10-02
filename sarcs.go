package arcs

import (
	"sort"
	"time"
)

const SARC_DURATION = time.Microsecond * 50

type SArc struct {
	Locator  string
	Order    float64
	Children []*SArc
	lock     bool
	hashmap  map[string]*SArc
}

func (node *SArc) HashQuery(loc string) *SArc {
	if node.lock {
		<-time.After(SARC_DURATION)
		return node.HashQuery(loc)
	}
	return node.unsafeHashQuery(loc)
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
	root.Children = make([]*SArc, 0, len(arcs))
	root.hashmap = make(map[string]*SArc)
	root.lock = true
	go func(rroot *SArc, aarcs []Arc, aarcrole string) {
		defer func() { rroot.lock = false }()
		sort.SliceStable(aarcs, func(i, j int) bool { return aarcs[i].Order < aarcs[j].Order })
		for _, arc := range aarcs {
			if arc.Arcrole == aarcrole {
				from := rroot.unsafeHashQuery(arc.From)
				if from != nil {
					to := from.unsafeHashQuery(arc.To)
					if to != nil {
						toIndex := from.getIndex(arc.To)
						from.Children[toIndex] = from.Children[len(from.Children)-1]
						from.Children = from.Children[:len(from.Children)-1]
					} else {
						to = rroot.unsafeHashQuery(arc.To)
						if to != nil {
							toIndex := rroot.getIndex(arc.To)
							rroot.Children[toIndex] = rroot.Children[len(rroot.Children)-1]
							rroot.Children = rroot.Children[:len(rroot.Children)-1]
						} else {
							order := arc.Order
							to = &SArc{
								Locator:  arc.To,
								Order:    order,
								Children: make([]*SArc, 0, len(aarcs)),
								hashmap:  make(map[string]*SArc),
							}
						}
					}
					from.Children = append(from.Children, to)
					from.hashmap[arc.To] = to
					rroot.hashmap[arc.To] = to
				} else {
					from = &SArc{
						Locator:  arc.From,
						Children: make([]*SArc, 0, len(aarcs)),
						hashmap:  make(map[string]*SArc),
					}
					rroot.Children = append(rroot.Children, from)
					to := rroot.unsafeHashQuery(arc.To)
					if to != nil {
						toIndex := rroot.getIndex(arc.To)
						rroot.Children[toIndex] = rroot.Children[len(rroot.Children)-1]
						rroot.Children = rroot.Children[:len(rroot.Children)-1]
					} else {
						order := arc.Order
						to = &SArc{
							Locator:  arc.To,
							Order:    order,
							Children: make([]*SArc, 0, len(aarcs)),
						}
					}
					from.Children = append(from.Children, to)
					rroot.hashmap[arc.From] = from
					from.hashmap[arc.To] = to
					rroot.hashmap[arc.To] = to
				}
			}
		}
	}(root, arcs, arcrole)
	return root
}

func (node *SArc) Paths(prior Path) []Path {
	if node == nil {
		return []Path{}
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
