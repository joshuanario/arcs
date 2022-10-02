package arcs

import (
	"sort"
	"sync"
	"time"
)

const SARC_DURATION = time.Microsecond * 50

type SArc struct {
	Locator  string
	Order    float64
	Children []*SArc
	lock     bool
	mutex    sync.Mutex
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
	node.mutex.Lock()
	defer node.mutex.Unlock()
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
						rroot.Children[toIndex] = rroot.Children[len(rroot.Children)-1]
						rroot.Children = rroot.Children[:len(rroot.Children)-1]
						from.Children = append(from.Children, to)
						from.mutex.Lock()
						rroot.mutex.Lock()
						from.hashmap[arc.To] = to
						rroot.hashmap[arc.To] = to
						from.mutex.Unlock()
						rroot.mutex.Unlock()
					} else {
						order := arc.Order
						to = &SArc{
							Locator:  arc.To,
							Order:    order,
							Children: make([]*SArc, 0, len(aarcs)),
							hashmap:  make(map[string]*SArc),
						}
						from.Children = append(from.Children, to)
						from.mutex.Lock()
						root.mutex.Lock()
						from.hashmap[arc.To] = to
						rroot.hashmap[arc.To] = to
						from.mutex.Unlock()
						root.mutex.Unlock()
					}
				} else {
					from = &SArc{
						Locator:  arc.From,
						Children: make([]*SArc, 0, len(aarcs)),
						hashmap:  make(map[string]*SArc),
					}
					rroot.Children = append(rroot.Children, from)
					rroot.mutex.Lock()
					rroot.hashmap[arc.From] = from
					rroot.mutex.Unlock()
					to := rroot.unsafeHashQuery(arc.To)
					if to != nil {
						toIndex := rroot.getIndex(arc.To)
						rroot.Children[toIndex] = rroot.Children[len(rroot.Children)-1]
						rroot.Children = rroot.Children[:len(rroot.Children)-1]
						from.Children = append(from.Children, to)
						from.mutex.Lock()
						rroot.mutex.Lock()
						from.hashmap[arc.To] = to
						rroot.hashmap[arc.To] = to
						from.mutex.Unlock()
						rroot.mutex.Unlock()
					} else {
						order := arc.Order
						to = &SArc{
							Locator:  arc.To,
							Order:    order,
							Children: make([]*SArc, 0, len(aarcs)),
						}
						from.Children = append(from.Children, to)
						from.mutex.Lock()
						rroot.mutex.Lock()
						from.hashmap[arc.To] = to
						rroot.hashmap[arc.To] = to
						from.mutex.Unlock()
						rroot.mutex.Unlock()
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
