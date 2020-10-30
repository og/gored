package red

import (
	"errors"
	gconv "github.com/og/x/conv"
	"time"
)


func (c Client) SET(key Key, value string, duration time.Duration) error {
	if duration < time.Millisecond {
		return errors.New("gored: SET(key, value, duration) duration can not less at millisecond (" + duration.String() + ")")
	}
	_, err := c.DoKey(nil, SET, key, value, "PX", gconv.Int64String(duration.Milliseconds()))
	return err
}
func (c Client) SETNeverExpire(key Key, value string)  (error){
	_, err := c.DoKey(nil, SET, key, value)
	return err
}
func (c Client) GET(key Key) (value string, has bool, err error) {
	empty, err := c.DoKey(&value, GET, key)
	if err != nil {
		return
	}
	if empty.IsNil == false {
		has = true
	}
	return
}
