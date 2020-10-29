package red_test

import (
	red "github.com/og/gored"
	gtest "github.com/og/x/test"
	"testing"
	"time"
)



func KeyMobileSignInAuthCode(mobile string) red.Key {
	return red.NewKey("mobile", "authCode", "mobile", mobile)
}
func  TestGetSetDeL(t *testing.T) {
	as := gtest.NewAS(t)
	var err error
	_=err
	key := KeyMobileSignInAuthCode("13000000000")
	{
		_, err = c.DEL(key)
		as.NoError(err)
		{
			value, has, err := c.GET(key)
			as.NoError(err)
			as.Equal(has, false)
			as.Equal("", value)
		}
		as.NoError(c.SETNeverExpire(key, "red"))
		{
			value, has, err := c.GET(key)
			as.NoError(err)
			as.Equal(has, true)
			as.Equal("red", value)
		}
		_, err = c.DEL(key)
		as.NoError(err)
		{
			{
				value, has, err := c.GET(key)
				as.NoError(err)
				as.Equal(has, false)
				as.Equal("", value)
			}
		}
	}
	{
		_, err := c.DEL(key)
		as.NoError(err)
		err = c.SET(key, "abcd", time.Second)
		as.NoError(err)
		{
			code, hasCode, err := c.GET(key)
			as.NoError(err)
			as.True(hasCode)
			as.Equal(code, "abcd")
		}
		time.Sleep(1*time.Second)
		{
			code, hasCode, err := c.GET(key)
			as.NoError(err)
			as.False(hasCode)
			as.Equal(code, "")
		}
	}
}
