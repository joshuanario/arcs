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
		return node.unsafeHashQuery(loc)
	}
	return node.unsafeHashQuery(loc)
}

func (node *SArc) unsafeHashQuery(loc string) *SArc {
	if node.Locator == loc {
		return node
	}
	return node.hashmap[loc]
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
			if arc.Arcrole == arcrole {
				from := rroot.unsafeHashQuery(arc.From)
				if from != nil {
					to := root.unsafeHashQuery(arc.To)
					if to != nil {
						continue
					} else {
						order := arc.Order
						to = &SArc{
							Locator:  arc.To,
							Order:    order,
							Children: make([]*SArc, 0, len(aarcs)),
							hashmap:  make(map[string]*SArc),
						}
						from.Children = append(from.Children, to)
						from.hashmap[arc.To] = to
						rroot.hashmap[arc.To] = to
					}
				} else {
					from = &SArc{
						Locator:  arc.From,
						Children: make([]*SArc, 0, len(aarcs)),
						hashmap:  make(map[string]*SArc),
					}
					rroot.Children = append(rroot.Children, from)
					rroot.hashmap[arc.From] = from
					to := rroot.unsafeHashQuery(arc.To)
					if to != nil {
						continue
					} else {
						order := arc.Order
						to = &SArc{
							Locator:  arc.To,
							Order:    order,
							Children: make([]*SArc, 0, len(aarcs)),
						}
						from.Children = append(from.Children, to)
						from.hashmap[arc.To] = to
						rroot.hashmap[arc.To] = to
					}
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
	newPath := append(prior, node.Locator)
	if node.lock {
		<-time.After(SARC_DURATION)
		return node.Paths(prior)
	}
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
