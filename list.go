package red

func (c Client) LPUSH(key Key, value ...string) (length int, err error) {
	_, err = c.DoKey(&length, LPUSH, key, value...)
	if err != nil {return}
	return
}
func (c Client) RPUSH(key Key, value ...string) (length int, err error) {
	_, err = c.DoKey(&length, RPUSH, key, value...)
	if err != nil {return}
	return
}

func (c Client) LPOP(key Key) (value string, hasValue bool, err error) {
	empty, err := c.DoKey(&value, LPOP, key)
	if err != nil {return}
	hasValue = empty.IsNil == false
	return
}
func (c Client) RPOP(key Key) (value string, hasValue bool, err error) {
	empty, err := c.DoKey(&value, RPOP, key)
	if err != nil {return}
	hasValue = empty.IsNil == false
	return
}
func (c Client) LLEN(key Key) (length int, err error) {
	_, err = c.DoKey(&length, LLEN, key)
	if err != nil {return}
	return
}