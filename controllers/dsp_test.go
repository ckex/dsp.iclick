package controllers

import (
	"testing"
	"io/ioutil"
	"fmt"
	"strings"
)

func Test_Dsp(t *testing.T) {
	t.Log(" start test ... ")
	filePth := "/Users/Ckex/work/go-workspace/src/dsp.iclick/conf/ad-dsp.template"
	data, err := ioutil.ReadFile(filePth)
	if err != nil {
		t.Log(err)
	}

	str := string(data)
	fmt.Println(str)

	str = strings.Replace(str,"#click#","http://wwww.baidu.com",-1)
	fmt.Println(str)

	str = strings.Replace(str,"#imageurl#","http://wwww.i-click.com",-1)
	fmt.Println(str)

}
