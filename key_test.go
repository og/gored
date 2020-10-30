package red_test

import (
	red "github.com/og/gored"
	gtest "github.com/og/x/test"
	"testing"
	"time"
)

func TestClient_EXISTS(t *testing.T) {
	as := gtest.NewAS(t)
	{
		exists, err := c.EXISTS(red.NewKey(time.Now().String()))
		as.NoError(err)
		as.Equal(exists, false)
	}
	{
		key := red.NewKey("EXISTSKEY")
		c.SET(key, "a", time.Second)
		exists, err := c.EXISTS(key)
		as.NoError(err)
		as.Equal(exists, true)
	}
}
func TestClient_PEXPIRE(t *testing.T) {
	as := gtest.NewAS(t)
	{
		success, err := c.PEXPIRE(red.NewKey(time.Now().String()), time.Second)
		as.NoError(err)
		as.Equal(success, false)
	}
	{
		key := red.NewKey("PEXPIREsucesss")
		err := c.SETNeverExpire(key, "abc")
		as.NoError(err)
		success, err := c.PEXPIRE(key, time.Second)
		as.NoError(err)
		as.Equal(success, true)
		{
			value, has, err := c.GET(key)
			as.NoError(err)
			as.Equal(has, true)
			as.Equal(value, "abc")
		}
		time.Sleep(time.Millisecond*1300)
		{
			value, has, err := c.GET(key)
			as.NoError(err)
			as.Equal(has, false)
			as.Equal(value, "")
		}
	}
}
func TestClient_PEXPIREAT(t *testing.T) {
	as := gtest.NewAS(t)
	{
		success, err := c.PEXPIREAT(red.NewKey(time.Now().String()), time.Now().Add(time.Second))
		as.NoError(err)
		as.Equal(success, false)
	}
	{
		key := red.NewKey("PEXPIREATsucesss")
		err := c.SETNeverExpire(key, "abc")
		as.NoError(err)
		success, err := c.PEXPIREAT(key, time.Now().Add(time.Second))
		as.NoError(err)
		as.Equal(success, true)
		{
			value, has, err := c.GET(key)
			as.NoError(err)
			as.Equal(has, true)
			as.Equal(value, "abc")
		}
		time.Sleep(time.Millisecond*1300)
		{
			value, has, err := c.GET(key)
			as.NoError(err)
			as.Equal(has, false)
			as.Equal(value, "")
		}
	}
}
func TestClient_PTTL(t *testing.T) {
	as := gtest.NewAS(t)
	key := red.NewKey(time.Now().String())
	{
		duration, fault, err := c.PTTL(key)
		as.NoError(err)
		as.Equal(duration, time.Duration(0))
		as.Equal(fault, red.PTTLFault{
			ErrKeyNotExist:           true,
		})
	}
	{
		err := c.SET(key, "a", time.Millisecond*1111)
		as.NoError(err)
		duration, fault, err := c.PTTL(key)
		as.NoError(err)
		as.Range(int(duration.Milliseconds()), 1100, 1111)
		as.Equal(fault, red.PTTLFault{
		})
	}
	{
		err := c.SETNeverExpire(key, "a")
		as.NoError(err)
		duration, fault, err := c.PTTL(key)
		as.NoError(err)
		as.Equal(duration, time.Duration(0))
		as.Equal(fault, red.PTTLFault{
			ErrKeyNoAssociatedExpire: true,
		})
	}
}

func TestClient_KEYS(t *testing.T) {
	as := gtest.NewAS(t)
	key1 := red.NewKey("keys1")
	key2 := red.NewKey("keys2")
	{
		as.NoError(c.SET(key1, "1", time.Second))
		as.NoError(c.SET(key2, "2", time.Second))
		keys, err := c.KEYS("noexistkey")
		as.NoError(err)
		as.Equal(keys, []red.Key(nil))
	}
	{
		keys, err := c.KEYS("keys*")
		as.NoError(err)
		as.Equal(keys, []red.Key{red.NewKey("keys1"), red.NewKey("keys2")})
	}
}
func TestClient_PERSIST(t *testing.T) {
	as := gtest.NewAS(t)
	{
		fault, err := c.PERSIST(red.NewKey(time.Now().String()))
		as.NoError(err)
		as.Equal(fault, red.PERSISTFault{ErrKeyNotExistOrNotAssociatedTimeout: true})
	}
	{
		key := red.NewKey(time.Now().String())
		err := c.SET(key, "PERSISTkey", time.Second)
		as.NoError(err)
		fault, err := c.PERSIST(key)
		as.NoError(err)
		as.Equal(fault, red.PERSISTFault{})
	}
}
func TestClient_RENAME(t *testing.T) {
	as := gtest.NewAS(t)
	oldKey := red.NewKey("RENAMEoldKey")
	newKey := red.NewKey("RENAMEnewKey")
	var err error
	err = c.SET(oldKey, "a", time.Second)
	{
		value, hasValue, err := c.GET(oldKey) ; as.NoError(err)
		as.Equal(value, "a")
		as.Equal(hasValue, true)
	}
	{
		value, hasValue, err := c.GET(newKey) ; as.NoError(err)
		as.Equal(value, "")
		as.Equal(hasValue, false)
	}
	as.NoError(err)
	err = c.RENAME(oldKey, newKey)
	{
		value, hasValue, err := c.GET(oldKey) ; as.NoError(err)
		as.Equal(value, "")
		as.Equal(hasValue, false)
	}
	{
		value, hasValue, err := c.GET(newKey) ; as.NoError(err)
		as.Equal(value, "a")
		as.Equal(hasValue, true)
	}
	{
		err := c.RENAME(red.NewKey("notExistKey"), red.NewKey("newKey"))
		as.ErrorString(err, "ERR no such key")
	}
}
func TestClient_RENAMENX(t *testing.T) {
	as := gtest.NewAS(t)
	oldKey := red.NewKey("RENAMENXoldKey")
	newKey := red.NewKey("RENAMENXnewKey")
	{
		var err error
		newKeyExist ,err := c.RENAMENX(oldKey, newKey)
		as.ErrorString(err ,"ERR no such key")
		as.Equal(newKeyExist, false)
	}
	{
		var err error
		err = c.SET(oldKey, "a", time.Second)
		as.NoError(err)
		newKeyExist ,err := c.RENAMENX(oldKey, newKey)
		as.NoError(err)
		as.Equal(newKeyExist, false)
		value, hasValue, err := c.GET(newKey)
		as.NoError(err)
		as.Equal(hasValue, true)
		as.Equal(value, "a")

		err = c.SET(oldKey, "a", time.Second)
		as.NoError(err)
		newKeyExist ,err = c.RENAMENX(oldKey, newKey)
		as.NoError(err)
		as.Equal(newKeyExist, true)
	}

}