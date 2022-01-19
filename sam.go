package sam

type SamNode struct {
	Trans  map[int]*SamNode
	Length int
	Link   *SamNode
	Num    int
}

type Sam struct {
	Root *SamNode
	Last *SamNode
	Num  int
}

func NewSam() *Sam {
	root := &SamNode{
		Trans:  make(map[int]*SamNode),
		Length: 0,
		Link:   nil,
		Num:    0,
	}
	return &Sam{
		Root: root,
		Last: root,
	}
}

func (s *Sam) Extend(c int) {
	s.Num++
	newNode := &SamNode{
		Trans:  make(map[int]*SamNode),
		Length: s.Last.Length + 1,
		Num:    s.Num,
	}
	p := s.Last
	s.Last = newNode
	for ; p != nil && p.Trans[c] == nil; p = p.Link {
		p.Trans[c] = newNode
	}
	if p == nil {
		newNode.Link = s.Root
		return
	}
	q := p.Trans[c]
	if q.Length == p.Length+1 {
		newNode.Link = q
	} else {
		s.Num++
		clone := &SamNode{
			Trans:  make(map[int]*SamNode),
			Length: p.Length + 1,
			Num:    s.Num,
			Link:   q.Link,
		}
		for k, v := range q.Trans {
			clone.Trans[k] = v
		}
		p.Trans[c] = clone
		q.Link = clone
		newNode.Link = clone
		for ; p != nil; p = p.Link {
			if p.Trans[c] == q {
				p.Trans[c] = clone
			}
		}
	}
}
