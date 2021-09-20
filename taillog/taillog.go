package taillog

import (
	"fmt"

	"github.com/hpcloud/tail"
)

var (
	tailObj *tail.Tail
)

// 收集日志模块
func Init(fileName string) (err error) {
	config := tail.Config{
		ReOpen:    true,
		MustExist: false,
		Poll:      true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
	}
	tailObj, err = tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed, err: ", err)
		return
	}
	return
}

func ReadLogChan() <-chan *tail.Line {
	return tailObj.Lines
}
