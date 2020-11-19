package red

import (
	"github.com/mediocregopher/radix/v3"
)

type Client struct {
	core Connecter
}
func (c Client) Ping() (err error) {
	_, err = c.Do(nil, PING, "hello world")
	return err
}
func (c Client) Close() error {
	return c.core.Close()
}
type PoolConfig struct {
	Network string `eg:"tcp"`
	Addr string `eg:"127.0.0.1:6379"`
	PoolSize int `eg:"10"`
	PoolOpts []radix.PoolOpt
}
func NewPool (conf PoolConfig) (c Client, err error) {
	pool, err := radix.NewPool(conf.Network, conf.Addr, conf.PoolSize, conf.PoolOpts...)
	if err != nil {return}
	c = Client{
		core: pool,
	}
	return
}

type Connecter interface {
	Do(a radix.Action) error
	Close() error
	NumAvailConns() int
}