package logic

import (
	"fmt"
)

type BasicLogic struct {
	Errno int
	Error string
}

type BaseLogic struct {
	BasicLogic
}

//private
func (this *BasicLogic) throw_error(errno int, err string) error {
	this.Errno = errno
	this.Error = err
	return fmt.Errorf("errno:[%d] , error:[%s]", this.Errno, this.Error)
}
