package red_test

import (
	"github.com/og/gored"
	"github.com/og/x/test"
	"testing"
)

var c red.Client
func init () {
	var err error
	c, err = red.NewPool(red.PoolConfig{
		Network:  "tcp",
		Addr:     "127.0.0.1:6379",
		PoolSize:     10,
		PoolOpts: nil,
	})
	if err != nil {panic(err)}
}
func TestNewClient(t *testing.T) {
	as := gtest.NewAS(t)
	{
		c, err := red.NewPool(red.PoolConfig{
			Network:  "tcp",
			Addr:     "127.0.0.1:1111",
			PoolSize:     10,
			PoolOpts: nil,
		})
		as.ErrorString( err, "dial tcp 127.0.0.1:1111: connect: connection refused")
		_=c
	}
	{
		c, err := red.NewPool(red.PoolConfig{
			Network:  "tcp",
			Addr:     "127.0.0.1:6379",
			PoolSize:     10,
			PoolOpts: nil,
		})
		as.NoError(err)
		_ = c
	}
}

func TestClient_Ping(t *testing.T) {
	as := gtest.NewAS(t)
	{
		c, err := red.NewPool(red.PoolConfig{
			Network:  "tcp",
			Addr:     "127.0.0.1:6379",
			PoolSize:     10,
			PoolOpts: nil,
		})
		as.NoError(err)
		as.NoError(c.Close())// 关闭后再ping
		as.ErrorString(c.Ping(), "client is closed")
	}
	c, err := red.NewPool(red.PoolConfig{
		Network:  "tcp",
		Addr:     "127.0.0.1:6379",
		PoolSize:     10,
		PoolOpts: nil,
	})
	as.NoError(err)
	as.NoError(c.Ping())
}