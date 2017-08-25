package scprobe

import (
	"os/exec"
	"fmt"
	"log"
)

/*a interface*/
type CommonInf interface{
	Output() ([]byte,error)
}
/*a func*/
type CommandFunc func(string,...string) CommonInf

/*a var of func*/
var DefaultCommandFunc CommandFunc = func(name string,arg...string)CommonInf {
	return exec.Command(name,arg...)
}



func genFFProbeInput(filepath string)( Ctx, error){
	fmt.Println("######4444################")
	if len(filepath) == 0 {
		return nil,fmt.Errorf("err probe path cannot be empty")
	}
	var params *ParamSet
	params = NewParamSet(
		NewParam("v", "quiet"),
		NewParam("print_format", "json"),
		NewParam("show_format", nil),
		NewParam("show_streams", nil),
	)
	fmt.Printf("#### create input with filepath %s\n",filepath)
	input := NewInputCtx(filepath,params)
	if input == nil {
		fmt.Println("------1------input is nil ")
	}else{
		fmt.Println("------2-----input is not nil")
	}
	fmt.Printf("----3--the input :%v",input)
	return input, nil
}
func genCommand(binPath string ,input Ctx)(cmd *exec.Cmd , err error){
	if input == nil {
		log.Println("-------input is null-----err to gen cmd")
	}

	cmdline, err := NewCommand(binPath, input)
	fmt.Println("Calling", cmdline.Path, cmdline.Slice())

	//cmd1:=DefaultCommandFunc(cmdline.Path, cmdline.Slice()...)
	//cmd1.Output()
	if cmdline == nil {
		log.Println("ERROR To gen cmdline\n")
		return
	}
	cmd=exec.Command(cmdline.Path,cmdline.Slice()...)

	return
}
