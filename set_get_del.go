package red

import (
	gconv "github.com/og/x/conv"
	"time"
)


func (c Client) SET(key Key, value string, duration time.Duration) error {
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
func (c Client) DEL(keys ...Key) (length int, err error) {
	sKeys := keysToStrings(keys)
	_, err = c.Do(&length, DEL, sKeys...)
	if err !=nil {return}
	return
}