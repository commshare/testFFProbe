package scprobe

import "fmt"

/*contains the key and/or value of a command line parameter*/
type Param struct {
	Key interface{}
	Value interface{}
}

/*create a parameter
if `key` is not a string ,it is nil and ignored
if `value` is nil , param.value is nil
it will return a param, even if all fields are nil
*/
func NewParam(key,value interface{}) (param Param){
	switch key.(type) {
	case string :
		param.Key =key
	default: /*not a string*/
		param.Key = nil
	
	}
	param.Value = value
	return
}


/* convert a param to a []string
if `key` is not a string ,it won't be added to the slice
if `value ` is nil ,it wont't be added to the slice
always returns a []string slice , even if empty
*/
func (p Param)Slice() (pair []string){
	pair =[]string{}
	if p.Key != nil {
		pair = append(pair,fmt.Sprintf("-%v",p.Key))
	}
	if p.Value != nil {
		pair =append(pair,fmt.Sprintf("%v",p.Value))
	}
	return
}

