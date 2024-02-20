package matlab

import "testing"

func TestMat_Open(t *testing.T) {
	m := &Mat{}
	err := m.Open(`target_VIS_intensity.mat`, r.String())
	if err != nil {
		t.Log(err)
	}
	dirs := m.GetDir()
	for i, dir := range dirs {
		t.Log(i, dir)
	}
	first := m.GetVariable(dirs[0])
	second := m.GetVariable(dirs[1])
	t.Log(first.IsEmpty())
	t.Log(second.GetElementSize())
	t.Log(second.GetNumberOfElements())
	t.Log(second.GetM())
	t.Log(second.GetN())
}
