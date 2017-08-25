package scprobe


/*a interface to reprensent input or output for ffprobe*/
type Ctx interface {
	Name() string
	AddParam(...Param)
	Slice() []string
}

type BaseCtx struct {
	Filename string
	Params *ParamSet
}

func (b BaseCtx) AddParam(param ... Param){
	b.Params.Add(param...)
}

func (b BaseCtx) Name() string {
	return b.Filename
}
/*represents an input for ffprobe/ffmpeg*/
type InputCtx struct {
	BaseCtx
}

/*a new input Ctx to ffprobe , must pass  file path and param set */
func NewInputCtx(filepath string,params *ParamSet) Ctx{
	/*TODO 返回的是一个地址啊*/
	return &InputCtx{
		BaseCtx:BaseCtx{
			Filename:filepath,
			Params:params,
		},
	}
}

/*how to pass InputCcx to ffprobe*/
func (i InputCtx) Slice() (result []string){
	result = append(result,i.Params.Slice()...)
	result = append(result,"-i",i.Filename)
	return
}

type OutputCtx struct {
	BaseCtx
}
func (o OutputCtx)Slice() (result []string) {
	result = append(result,o.Params.Slice()...)
	result = append(result,o.Filename,"-y")
	return
}
func NewOutput(filepath string,params *ParamSet) Ctx {
	return &OutputCtx{
		BaseCtx : BaseCtx{
			Filename:filepath,
			Params:params,
		},
	}
}