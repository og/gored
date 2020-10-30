package red

import (
	"errors"
	gconv "github.com/og/x/conv"
	"strings"
	"time"
)

type Key struct {
	value string
}
func (k Key) String() string {
	return k.value
}
func NewKey(k ...string) Key {
	return Key{value: strings.Join(k, ":")}
}

func stringsToKeys(sList []string) (keys []Key) {
	for _, s := range sList {
		keys = append(keys, NewKey(s))
	}
	return
}
func keysToStrings(keys []Key) (stringKeys []string) {
	for _, key := range keys {
		stringKeys = append(stringKeys, key.String())
	}
	return
}


func (c Client) DEL(keys ...Key) (length int, err error) {
	sKeys := keysToStrings(keys)
	_, err = c.Do(&length, DEL, sKeys...)
	if err !=nil {return}
	return
}

func (c Client) KEYS(pattern string) (keys []Key, err error) {
	var sList []string
	_, err = c.Do(&sList, KEYS, pattern)
	if err != nil {return}
	return stringsToKeys(sList), nil
}
func (c Client) EXISTS(key Key) (exists bool, err error) {
	_, err = c.DoKey(&exists, EXISTS, key)
	if err != nil {return}
	return
}

func (c Client) PEXPIRE(key Key, duration time.Duration) (success bool, err error) {
	if duration < time.Millisecond {
		return false, errors.New("gored: PEXPIRE(key, duration) duration can not less at millisecond (" + duration.String() + ")")
	}
	_, err = c.DoKey(&success, PEXPIRE, key, gconv.Int64String(duration.Milliseconds()))
	if err != nil {return}
	return
}

func (c Client) PEXPIREAT(key Key, at time.Time) (success bool, err error) {
	_, err = c.DoKey(&success, PEXPIREAT, key, gconv.Int64String(at.UnixNano() / 1e6))
	if err != nil {return}
	return
}
type PTTLFault struct {
	ErrKeyNotExist bool
	ErrKeyNoAssociatedExpire bool
}
func (err PTTLFault) Has() bool {
	return err.ErrKeyNotExist || err.ErrKeyNoAssociatedExpire
}
func (err PTTLFault) Handle(
	ErrKeyNotExist func(_ErrKeyNotExist int),
	ErrKeyNoAssociatedExpire func(_ErrKeyNoAssociatedExpire bool),
	) {
	switch {
	case err.ErrKeyNotExist:
		ErrKeyNotExist(0)
	case err.ErrKeyNoAssociatedExpire:
		ErrKeyNoAssociatedExpire(false)
	default:
		panic(errors.New("PTTError not a error"))
	}
}
func (c Client) PTTL(key Key) (duration time.Duration, fault PTTLFault, err error) {
	var result int64
	_, err = c.DoKey(&result, PTTL, key)
	if err != nil {return}
	switch result {
	case -2:
		fault = PTTLFault{ErrKeyNotExist:true}
		return
	case -1:
		fault = PTTLFault{ErrKeyNoAssociatedExpire:true}
		return
	default:
		duration = time.Duration(int64(time.Millisecond) * result)
	}
	return
}
type PERSISTFault struct {
	ErrKeyNotExistOrNotAssociatedTimeout bool
}
func (c Client) PERSIST(key Key) (fault PERSISTFault, err error) {
	result := 0
	_, err = c.DoKey(&result, PERSIST, key)
	if err != nil {return}
	if result == 0 {
		return PERSISTFault{ErrKeyNotExistOrNotAssociatedTimeout:true}, nil
	}
	return
}
func (c Client) RENAME(key Key, newKey Key) (err error) {
	_, err = c.DoKey(nil, RENAME, key, newKey.String())
	if err != nil {return}
	return nil
}
func (c Client) RENAMENX(key Key, newKey Key) (newKeyExist bool, err error) {
	result := 0
	_, err = c.DoKey(&result, RENAMENX, key, newKey.String())
	if err != nil {return}
	if result == 0 {
		newKeyExist = true
	}
	return
}