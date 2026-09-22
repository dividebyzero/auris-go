package auris

// RadioGroup keeps a set of Auris radios mutually exclusive without coupling
// the individual Radio widget to application state.
type RadioGroup struct {
	Radios []*Radio
	Selected int
	OnChanged func(int)
}

func NewRadioGroup(labels []string, selected int, changed func(int)) *RadioGroup {
	g:=&RadioGroup{Selected:selected,OnChanged:changed}
	g.Radios=make([]*Radio,len(labels))
	for i,label:=range labels {
		index:=i
		g.Radios[i]=NewRadio(label,i==selected,func(){ g.Select(index) })
	}
	return g
}
func (g *RadioGroup) Select(index int) {
	if index<0 || index>=len(g.Radios) { return }
	g.Selected=index
	for i,r:=range g.Radios { r.SetSelected(i==index) }
	if g.OnChanged!=nil { g.OnChanged(index) }
}
