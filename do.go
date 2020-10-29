package red

import (
	"github.com/mediocregopher/radix/v3"
)

type Empty struct {
	IsNil        bool
	IsEmptyArray bool
}
func (c Client) Do(ptr interface{}, cmd CMD, args ...string) (empty Empty, err error) {
	rcv := radix.MaybeNil{
		Rcv: ptr,
	}
	err = c.core.Do(radix.Cmd(
		&rcv,
		cmd.String(),
		args...
	))
	if err != nil {return}
	empty.IsNil = rcv.Nil
	empty.IsEmptyArray = rcv.EmptyArray
	return
}
func (c Client) DoKey(ptr interface{}, cmd CMD, key Key, args ...string) (empty Empty, err error) {
	args = append([]string{key.String()}, args...)
	return c.Do(ptr, cmd, args...)
}