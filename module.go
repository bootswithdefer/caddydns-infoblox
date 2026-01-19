package infoblox

import (
	infoblox "github.com/bootswithdefer/libdns-infoblox"
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
)

type Provider struct{ *infoblox.Provider }

func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.infoblox",
		New: func() caddy.Module { return &Provider{new(infoblox.Provider)} },
	}
}

func init() {
	caddy.RegisterModule(Provider{})
}

func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.Host = caddy.NewReplacer().ReplaceAll(p.Provider.Host, "")
	p.Provider.Version = caddy.NewReplacer().ReplaceAll(p.Provider.Version, "")
	p.Provider.Username = caddy.NewReplacer().ReplaceAll(p.Provider.Username, "")
	p.Provider.Password = caddy.NewReplacer().ReplaceAll(p.Provider.Password, "")
	return nil
}

func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}

		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "host":
				if d.NextArg() {
					p.Provider.Host = d.Val()
				} else {
					return d.ArgErr()
				}
			case "version":
				if d.NextArg() {
					p.Provider.Version = d.Val()
				} else {
					return d.ArgErr()
				}
			case "username":
				if d.NextArg() {
					p.Provider.Username = d.Val()
				} else {
					return d.ArgErr()
				}
			case "password":
				if d.NextArg() {
					p.Provider.Password = d.Val()
				} else {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}

	if p.Provider.Host == "" || p.Provider.Version == "" || p.Provider.Username == "" || p.Provider.Password == "" {
		return d.Err("missing config!")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
