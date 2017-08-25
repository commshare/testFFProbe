package ffprobe

import (
	"bytes"
	"encoding/json"
	"errors"
	"os/exec"
	"time"
	"fmt"
	"log"
)

var ErrBinNotFound error = errors.New("ffprobe bin not found")
var ErrTimeout error = errors.New("process timeout exceeded")
//var ErrTimeoutKillFail error = errors.New("")

var binPath string = "/usr/bin/ffprobe"

func SetFFProbeBinPath(newBinPath string) {
	binPath = newBinPath
}

func GetProbeData(filePath string, timeout time.Duration) (data *ProbeData, err error) {
	cmd := exec.Command(
		binPath,
		"-v", "quiet",
		"-print_format", "json",
		"-show_format",	
		"-show_streams",
		filePath,
	)
	var outputBuf bytes.Buffer
	cmd.Stdout = &outputBuf

	err = cmd.Start()
	if err != nil {
		if err == exec.ErrNotFound {
			err = ErrBinNotFound
		}
		return
	}

	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-time.After(timeout):
		err = cmd.Process.Kill()
		if err == nil {
			err = ErrTimeout
		}else{
			err=errors.New(fmt.Sprintf("timeout reached, but failed to kill: %s", err.Error()))
		}
		return
	case err = <-done:
		if err != nil {
			err = errors.New(fmt.Sprintf("process done with error = %s", err.Error()))
			return
		}
	}

	data = &ProbeData{}
	err = json.Unmarshal(outputBuf.Bytes(), data)
	if err != nil {
		log.Println("---------------json umarshal probe data fail------------------")
		return
	}
	return
}
