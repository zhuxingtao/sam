package sam

import (
	"testing"
)

func TestSubstringNum(t *testing.T) {
	f := make(map[*SamNode]bool)
	var res int
	s := "dafsdasafwreerdfsfdgh"
	sam := NewSam()
	for i := 0; i < len(s); i++ {
		sam.Extend(int(s[i] - 'a'))
	}
	type dfsFunc func(p *SamNode)
	var dfs dfsFunc
	dfs = func(p *SamNode) {
		if f[p] {
			return
		}
		f[p] = true
		if p != sam.Root {
			res += p.Length - p.Link.Length
		}
		for _, v := range p.Trans {
			dfs(v)
		}
	}
	dfs(sam.Root)
	if res != 216 {
		t.Errorf("fail")
	}
}
