package red_test

import (
	red "github.com/og/gored"
	gtest "github.com/og/x/test"
	"testing"
)

func TestClient_PUSH_POP(t *testing.T) {
	as := gtest.NewAS(t)
	key := red.NewKey("PUSH_POP")
	{
		_, err := c.DEL(key)
		as.NoError(err)
	}
	{
		length , err := c.LPUSH(key, "a")
		as.NoError(err)
		as.Equal(length, 1)
	}
	{
		length , err := c.LPUSH(key, "b","c")
		as.NoError(err)
		as.Equal(length, 3)
	}
	{
		value, hasValue, err := c.LPOP(key)
		as.NoError(err)
		as.Equal(value, "c")
		as.Equal(hasValue, true)
	}
	{
		value, hasValue, err := c.LPOP(key)
		as.NoError(err)
		as.Equal(value, "b")
		as.Equal(hasValue, true)
	}
	{
		length , err := c.RPUSH(key, "z")
		as.NoError(err)
		as.Equal(length, 2)
	}
	{
		value, hasValue, err := c.RPOP(key)
		as.NoError(err)
		as.Equal(value, "z")
		as.Equal(hasValue, true)
	}
	{
		value, hasValue, err := c.RPOP(key)
		as.NoError(err)
		as.Equal(value, "a")
		as.Equal(hasValue, true)
	}
	{
		value, hasValue, err := c.RPOP(key)
		as.NoError(err)
		as.Equal(value, "")
		as.Equal(hasValue, false)
	}
	{
		length, err := c.LLEN(key)
		as.NoError(err)
		as.Equal(length, 0)
	}
	{
		_, err := c.LPUSH(key, "x")
		as.NoError(err)
		length, err := c.LLEN(key)
		as.NoError(err)
		as.Equal(length, 1)
	}
}