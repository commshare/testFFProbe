package scprobe

import "fmt"

type Command struct {
	Path string
	Input Ctx
	Outputs []Ctx
}


func NewCommand(path string,input Ctx,outputs...Ctx)(cmd *Command , err error ){
	if path == "" {
		return nil ,fmt.Errorf("err to create a Command with no path")
	}
	if input == nil {
		return nil,fmt.Errorf("err to create a Command with no input")
	}
	cmd = &Command{
		Path: path,
		Input : input ,
	}

	for _, output := range outputs {
		if output != nil {
			cmd.Outputs = append(cmd.Outputs,output)
		}
	}
	return cmd,nil

}
/*Returns a []string slice of how a ffmpeg/ffprobe call should be represented.*/
func (c * Command)Slice() (results []string){
	results = []string{}
	results = append(results,c.Input.Slice()...)
	for _,output := range c.Outputs {
		results = append(results,output.Slice()...)
	}
	return
}