package demo

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

func init() {
	plugin.RegisterPlugin("demo", setup)
}

func setup(c *caddy.Controller) error {
	c.Next() // 'demo'
	if c.NextArg() {
		return plugin.Error("demo", c.ArgErr())
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Demo{}
	})

	return nil
}
