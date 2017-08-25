package scprobe

/*a set of Params*/
type ParamSet []string

func NewParamSet(params ...Param) *ParamSet {
	items := &ParamSet{}
	items.Add(params ...)
	return items
}
/*return the lenght of this slice*/
func (p *ParamSet)Length() int {
	return len([]string(*p))
}

func (p *ParamSet)Add(params ...Param){
	for _,param := range params {
		slice := param.Slice()
		if len(slice)>0 {
			*p= append(*p,slice...)
		}
	}
}

/*return a []string slice,event if it is empty*/
func (p *ParamSet)Slice() []string {
	return []string(*p)
}