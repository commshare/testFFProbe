package scprobe

import (
	"time"
	"bytes"
	"os/exec"
	"fmt"
	"encoding/json"
	"log"
)
var binPath string = "/usr/bin/ffprobe"
var filepath string
func SetFFProbeBinPath(newBinPath string) {
	binPath = newBinPath
}
func GetFFProbeData(filepath string ,timeout time.Duration)(data *FFProbeInfo ,err error){
	fmt.Printf("binpath :%s\n",binPath)

	input,err :=genFFProbeInput(filepath)
	cmd ,err:= genCommand(binPath,input)
	if cmd == nil {
		log.Println("-------gen cmd fail -----")
		return
	}
	var outputBuf bytes.Buffer
	cmd.Stdout =&outputBuf

	err = cmd.Start()
	if err != nil {
		if err == exec.ErrNotFound {
			err = fmt.Errorf("ffprobe bin not found")
		}
		return
	}
	done := make(chan error,1)
	go func(){
		done <- cmd.Wait()
	}()
	select {
	case <-time.After(timeout):
		err = cmd.Process.Kill()
		if err == nil {
			err = fmt.Errorf("ffprobe process timeout")
		}else{
			err = fmt.Errorf("ffprobe cmd process kill fail")
		}
		return
		case err= <-done:
		if err!=nil {
			err = fmt.Errorf("process done with err %s",err.Error())
			return
		}
	}
	fmt.Printf("##out###%v \n",string(outputBuf.Bytes()))

	/*ret:exec: Stdout already set */
/*	out, ret:=cmd.Output()
	fmt.Printf("##out###%v ret:%v \n",string(out),ret)*/
	data = &FFProbeInfo{}
	err = json.Unmarshal(outputBuf.Bytes(),data)
	if err!= nil {
		err = fmt.Errorf("json unmarshal to probe data fail")
		return
	}
	fmt.Printf("$$$$$$$$$probedata:%v\n",data)
	return
}